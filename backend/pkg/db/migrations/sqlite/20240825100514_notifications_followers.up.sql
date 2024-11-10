CREATE TABLE IF NOT EXISTS "notifications" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "user_id" INTEGER NOT NULL,
    "message" TEXT NOT NULL,
    "type" TEXT CHECK(type IN ('f_req', 'f_resp', 'g_req', 'g_inv', 'g_eve', 'g_resp', 'other')) NOT NULL,
    "id_ref" INTEGER,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "followers" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "follower_id" INTEGER NOT NULL,
    "followee_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "status" TEXT CHECK(status IN ('pending', 'accepted', 'rejected')) DEFAULT 'pending',
    FOREIGN KEY ("follower_id") REFERENCES "users"("id") ON DELETE CASCADE,
    FOREIGN KEY ("followee_id") REFERENCES "users"("id") ON DELETE CASCADE,
    UNIQUE ("follower_id", "followee_id")
);

CREATE TRIGGER IF NOT EXISTS delete_notifications_on_view
AFTER UPDATE ON notifications
FOR EACH ROW
WHEN NEW.is_read = TRUE AND OLD.is_read = FALSE
BEGIN
    DELETE FROM notifications
    WHERE user_id = NEW.user_id
      AND type IN ('f_resp', 'g_resp', 'other')
      AND id != NEW.id
      AND id NOT IN (
          SELECT id FROM notifications
          WHERE user_id = NEW.user_id
          ORDER BY id DESC
          LIMIT 20
      );
END;

CREATE TRIGGER IF NOT EXISTS before_insert_notification
BEFORE INSERT ON "notifications"
FOR EACH ROW
WHEN EXISTS (
    SELECT 1 FROM "notifications"
    WHERE user_id = NEW.user_id
      AND message = NEW.message
      AND type = NEW.type
      AND is_read = FALSE
      AND id_ref = NEW.id_ref
)
BEGIN
    UPDATE "notifications"
    SET created_at = CURRENT_TIMESTAMP
    WHERE user_id = NEW.user_id
      AND message = NEW.message
      AND type = NEW.type
      AND is_read = FALSE
      AND id_ref = NEW.id_ref;

    SELECT RAISE(IGNORE);
END;

CREATE TRIGGER IF NOT EXISTS after_follow_request_private
AFTER INSERT ON followers
FOR EACH ROW
WHEN (SELECT profile_visibility FROM users WHERE id = NEW.followee_id) = 'private'
BEGIN
    INSERT INTO notifications (user_id, message, type, is_read, created_at, id_ref)
    VALUES (
        NEW.followee_id,
        (SELECT nickname FROM users WHERE id = NEW.follower_id) || ' wants to follow you',
        'f_req',
        FALSE,
        CURRENT_TIMESTAMP,
        NEW.follower_id
    );
END;


CREATE TRIGGER IF NOT EXISTS after_follow_request_public
AFTER INSERT ON followers
FOR EACH ROW
WHEN (SELECT profile_visibility FROM users WHERE id = NEW.followee_id) = 'public'
BEGIN
    UPDATE followers
    SET status = 'accepted'
    WHERE id = NEW.id;

    INSERT INTO notifications (user_id, message, type, is_read, created_at, id_ref)
    VALUES (
        NEW.followee_id,
        'You have a new follower: ' || (SELECT nickname FROM users WHERE id = NEW.follower_id),
        'f_resp',
        FALSE,
        CURRENT_TIMESTAMP,
        NEW.follower_id
    );

    UPDATE users
    SET following_count = (SELECT COUNT(*) FROM followers WHERE follower_id = NEW.follower_id AND status = 'accepted')
    WHERE id = NEW.follower_id;

    UPDATE users
    SET follower_count = (SELECT COUNT(*) FROM followers WHERE followee_id = NEW.followee_id AND status = 'accepted')
    WHERE id = NEW.followee_id;
END;

CREATE TRIGGER IF NOT EXISTS after_follow_accept
AFTER UPDATE ON followers
FOR EACH ROW
WHEN OLD.status = 'pending' AND NEW.status = 'accepted'
BEGIN
    INSERT INTO notifications (user_id, message, type, is_read, created_at, id_ref)
    VALUES (
        NEW.follower_id,
        (SELECT nickname FROM users WHERE id = NEW.followee_id) || ' accepted your follow request',
        'f_resp',
        FALSE,
        CURRENT_TIMESTAMP,
        NEW.followee_id
    );

    UPDATE users
    SET following_count = (SELECT COUNT(*) FROM followers WHERE follower_id = NEW.follower_id AND status = 'accepted')
    WHERE id = NEW.follower_id;

    UPDATE users
    SET follower_count = (SELECT COUNT(*) FROM followers WHERE followee_id = NEW.followee_id AND status = 'accepted')
    WHERE id = NEW.followee_id;
END;

CREATE TRIGGER IF NOT EXISTS after_follow_reject
AFTER UPDATE ON followers
FOR EACH ROW
WHEN OLD.status = 'pending' AND NEW.status = 'rejected'
BEGIN
    INSERT INTO notifications (user_id, message, type, is_read, created_at, id_ref)
    VALUES (
        NEW.follower_id,
        (SELECT nickname FROM users WHERE id = NEW.followee_id) || ' rejected your follow request',
        'f_resp',
        FALSE,
        CURRENT_TIMESTAMP,
        NEW.followee_id
    );

    DELETE FROM followers
    WHERE follower_id = NEW.follower_id
      AND followee_id = NEW.followee_id;
END;

CREATE TRIGGER IF NOT EXISTS after_follower_removal
AFTER DELETE ON followers
FOR EACH ROW
BEGIN
    UPDATE users
    SET following_count = (SELECT COUNT(*) FROM followers WHERE follower_id = OLD.follower_id AND status = 'accepted')
    WHERE id = OLD.follower_id;

    UPDATE users
    SET follower_count = (SELECT COUNT(*) FROM followers WHERE followee_id = OLD.followee_id AND status = 'accepted')
    WHERE id = OLD.followee_id;
END;