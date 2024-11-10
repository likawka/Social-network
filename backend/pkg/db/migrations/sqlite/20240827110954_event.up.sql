CREATE TABLE IF NOT EXISTS "group_events" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "group_id" INTEGER NOT NULL,
    "creator_id" INTEGER NOT NULL,
    "title" TEXT NOT NULL,
    "description" TEXT,
    "event_time" TIMESTAMP NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "count_going" INTEGER DEFAULT 0,
    FOREIGN KEY ("group_id") REFERENCES "groups"("id") ON DELETE CASCADE,
    FOREIGN KEY ("creator_id") REFERENCES "users"("id") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "event_responses" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "event_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "response" TEXT CHECK(response IN ('going', 'not_going')) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("event_id") REFERENCES "group_events"("id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE
);

CREATE TRIGGER IF NOT EXISTS update_count_going
AFTER INSERT ON "event_responses"
FOR EACH ROW
BEGIN
    UPDATE "group_events"
    SET "count_going" = (
        SELECT COUNT(*) 
        FROM "event_responses" 
        WHERE "event_id" = NEW.event_id 
        AND "response" = 'going'
    )
    WHERE "id" = NEW.event_id;
END;

CREATE TRIGGER IF NOT EXISTS after_event_creation
AFTER INSERT ON "group_events"
FOR EACH ROW
BEGIN
    INSERT INTO "notifications" (user_id, message, type, is_read, created_at)
    SELECT 
        gm.user_id,
        'A new event titled "' || NEW.title || '" has been created in the group: ' || (SELECT title FROM "groups" WHERE id = NEW.group_id),
        'g_eve',
        FALSE,
        CURRENT_TIMESTAMP,
        NEW.id
    FROM 
        "group_memberships" gm
    WHERE 
        gm.group_id = NEW.group_id;
END;