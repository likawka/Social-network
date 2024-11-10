CREATE TABLE IF NOT EXISTS "groups" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "title" TEXT NOT NULL,
    "description" TEXT,
    "banner_color" TEXT DEFAULT '#ffffff' NOT NULL,
    "creator_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "member_count" INTEGER DEFAULT 0,
    FOREIGN KEY ("creator_id") REFERENCES "users"("id") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "group_invitations" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "group_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "request_type" TEXT CHECK(request_type IN ('inv', 'j_req')) NOT NULL,
    "status" TEXT CHECK(status IN ('pending', 'accepted', 'declined')) DEFAULT 'pending',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("group_id") REFERENCES "groups"("id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE
);

CREATE TRIGGER before_insert_group_invitation
BEFORE INSERT ON "group_invitations"
FOR EACH ROW
BEGIN
    DELETE FROM "group_invitations"
    WHERE group_id = NEW.group_id
    AND user_id = NEW.user_id
    AND request_type = NEW.request_type
    AND status = 'pending';
END;

CREATE TABLE IF NOT EXISTS "group_memberships" (
    "group_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "role" TEXT CHECK(role IN ('admin', 'member')) NOT NULL DEFAULT 'member', 
    "joined_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("group_id") REFERENCES "groups"("id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE,
    UNIQUE ("group_id", "user_id")
);


--

CREATE TRIGGER after_group_creation
AFTER INSERT ON groups
FOR EACH ROW
BEGIN
    INSERT INTO conversations (conversation_hash, conversation_type, creator_id, created_at)
    VALUES (
        (SELECT lower(hex(randomblob(16)))), 
        'group',
        NEW.id,
        CURRENT_TIMESTAMP
    );
    

    INSERT INTO conversation_users (conversation_hash, user_id, joined_at)
    VALUES (
        (SELECT conversation_hash FROM conversations WHERE rowid = (SELECT last_insert_rowid())),
        NEW.creator_id,
        CURRENT_TIMESTAMP
    );
    INSERT INTO "group_memberships" ("group_id", "user_id", "role", "joined_at")
    VALUES (NEW.id, NEW.creator_id, 'admin', CURRENT_TIMESTAMP);

END;

CREATE TRIGGER after_invitation_acceptance
AFTER UPDATE ON group_invitations
FOR EACH ROW
WHEN OLD.status = 'pending' AND NEW.status = 'accepted'
BEGIN
    INSERT INTO conversation_users (conversation_hash, user_id, joined_at)
    VALUES (
        (SELECT conversation_hash FROM conversations WHERE id = (SELECT group_id FROM group_memberships WHERE user_id = NEW.user_id)),
        NEW.user_id,
        CURRENT_TIMESTAMP
    );
    
    INSERT INTO notifications (user_id, message, type, is_read, created_at, id_ref)
    VALUES (
        NEW.user_id,
        'You have successfully joined the group: ' || (SELECT title FROM groups WHERE id = (SELECT group_id FROM group_memberships WHERE user_id = NEW.user_id)),
        'g_resp',
        FALSE,
        CURRENT_TIMESTAMP,
        NEW.group_id
    );
END;

--

CREATE TRIGGER IF NOT EXISTS after_invitation_creation
AFTER INSERT ON "group_invitations"
FOR EACH ROW
BEGIN
    INSERT INTO "notifications" (user_id, message, type, is_read, created_at, id_ref)
    VALUES (
        CASE 
            WHEN NEW.request_type = 'inv' THEN NEW.user_id
            WHEN NEW.request_type = 'j_req' THEN (SELECT creator_id FROM "groups" WHERE id = NEW.group_id)
        END,
        CASE 
            WHEN NEW.request_type = 'inv' THEN 'You have been invited to join the group: ' || (SELECT title FROM "groups" WHERE id = NEW.group_id)
            WHEN NEW.request_type = 'j_req' THEN (SELECT nickname FROM "users" WHERE id = NEW.user_id) || ' wants to join your group: ' || (SELECT title FROM "groups" WHERE id = NEW.group_id)
        END,
        CASE 
            WHEN NEW.request_type = 'inv' THEN 'g_inv'
            WHEN NEW.request_type = 'j_req' THEN 'g_req'
        END,
        FALSE,
        CURRENT_TIMESTAMP,
        NEW.group_id
    );
END;

CREATE TRIGGER after_user_delete
AFTER DELETE ON "users"
FOR EACH ROW
BEGIN
    UPDATE "groups"
    SET "member_count" = "member_count" - 1
    WHERE "id" IN (SELECT "group_id" FROM "group_memberships" WHERE "user_id" = OLD.id);
    DELETE FROM "group_memberships" WHERE "user_id" = OLD.id;
END;

CREATE TRIGGER after_membership_insert
AFTER INSERT ON "group_memberships"
FOR EACH ROW
BEGIN
    UPDATE "groups"
    SET "member_count" = "member_count" + 1
    WHERE "id" = NEW.group_id;
END;

CREATE TRIGGER after_membership_delete
AFTER DELETE ON group_memberships
FOR EACH ROW
BEGIN
    UPDATE groups
    SET member_count = member_count - 1
    WHERE id = OLD.group_id;
END;


CREATE TRIGGER after_group_membership_delete
AFTER DELETE ON group_memberships
FOR EACH ROW
BEGIN
    DELETE FROM conversation_users
    WHERE conversation_hash = (
        SELECT conversation_hash
        FROM conversations
        WHERE id = OLD.group_id
    ) AND user_id = OLD.user_id;
END;