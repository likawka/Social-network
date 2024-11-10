CREATE TABLE IF NOT EXISTS "users" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "email" TEXT UNIQUE NOT NULL,
    "password_hash" TEXT NOT NULL,
    "first_name" TEXT NOT NULL,
    "last_name" TEXT NOT NULL,
    "date_of_birth" TEXT NOT NULL,
    "avatar" BLOB,
    "nickname" TEXT UNIQUE,
    "about_me" TEXT,
    "banner_color" TEXT DEFAULT '#ffffff' NOT NULL,
    "profile_visibility" TEXT CHECK(profile_visibility IN ('public', 'private')) DEFAULT 'public',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "post_count" INTEGER DEFAULT 0,
    "comment_count" INTEGER DEFAULT 0,
    "follower_count" INTEGER DEFAULT 0,
    "following_count" INTEGER DEFAULT 0,
    "last_active" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "active_sessions" (
    "user_id" INTEGER NOT NULL,
    "session_id" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "expires_at" TIMESTAMP NOT NULL,
    FOREIGN KEY("user_id") REFERENCES "users"("id")
);

CREATE TRIGGER IF NOT EXISTS delete_expired_sessions AFTER INSERT ON "active_sessions" FOR EACH ROW BEGIN
DELETE FROM "active_sessions"
WHERE
    "expires_at" < CURRENT_TIMESTAMP;
END;