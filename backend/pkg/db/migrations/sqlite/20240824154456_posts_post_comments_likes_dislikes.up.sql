CREATE TABLE IF NOT EXISTS "posts" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "user_id" INTEGER NOT NULL,
    "group_id" INTEGER,
    "title" TEXT,
    "content" TEXT,
    "image" BLOB,
    "privacy" TEXT CHECK(privacy IN ('public', 'private', 'followers')) DEFAULT 'public',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "comment_count" INTEGER DEFAULT 0,
    "like_count" INTEGER DEFAULT 0,
    "dislike_count" INTEGER DEFAULT 0,
    FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE,
    FOREIGN KEY ("group_id") REFERENCES "groups"("id") ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS "likes_dislikes" (
    "user_id" INTEGER NOT NULL,
    "object_type" TEXT CHECK(object_type IN ('post', 'comment')) NOT NULL,
    "object_id" INTEGER NOT NULL,
    "reaction" TEXT CHECK(reaction IN ('like', 'dislike')) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("user_id") REFERENCES "users"("id"),
    UNIQUE ("user_id", "object_type", "object_id")
);

CREATE INDEX IF NOT EXISTS "idx_object_type_object_id" 
ON "likes_dislikes" ("object_type", "object_id");

CREATE INDEX IF NOT EXISTS "idx_user_id" 
ON "likes_dislikes" ("user_id");

CREATE TABLE IF NOT EXISTS "post_comments" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "post_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "content" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "like_count" INTEGER DEFAULT 0,
    "dislike_count" INTEGER DEFAULT 0,
    FOREIGN KEY ("post_id") REFERENCES "posts"("id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE
);

CREATE TRIGGER IF NOT EXISTS update_user_counts_on_post_insert 
AFTER INSERT ON "posts"
FOR EACH ROW 
BEGIN
    UPDATE "users"
    SET "post_count" = "post_count" + 1
    WHERE "id" = NEW."user_id";
END;

CREATE TRIGGER IF NOT EXISTS update_user_counts_on_post_delete 
AFTER DELETE ON "posts"
FOR EACH ROW 
BEGIN
    UPDATE "users"
    SET "post_count" = "post_count" - 1
    WHERE "id" = OLD."user_id";
END;

CREATE TRIGGER IF NOT EXISTS update_user_counts_on_comment_insert 
AFTER INSERT ON "post_comments"
FOR EACH ROW 
BEGIN
    UPDATE "users"
    SET "comment_count" = "comment_count" + 1
    WHERE "id" = NEW."user_id";
END;

CREATE TRIGGER IF NOT EXISTS update_user_counts_on_comment_delete 
AFTER DELETE ON "post_comments"
FOR EACH ROW 
BEGIN
    UPDATE "users"
    SET "comment_count" = "comment_count" - 1
    WHERE "id" = OLD."user_id";
END;

CREATE TRIGGER IF NOT EXISTS update_post_comment_count_on_comment_insert 
AFTER INSERT ON "post_comments"
FOR EACH ROW 
BEGIN
    UPDATE "posts"
    SET "comment_count" = "comment_count" + 1
    WHERE "id" = NEW."post_id";
END;

CREATE TRIGGER IF NOT EXISTS update_post_comment_count_on_comment_delete 
AFTER DELETE ON "post_comments"
FOR EACH ROW 
BEGIN
    UPDATE "posts"
    SET "comment_count" = "comment_count" - 1
    WHERE "id" = OLD."post_id";
END;

CREATE TRIGGER IF NOT EXISTS delete_posts_and_comments_on_user_delete 
AFTER DELETE ON "users"
FOR EACH ROW 
BEGIN
    DELETE FROM "posts"
    WHERE "user_id" = OLD."id";
    
    DELETE FROM "post_comments"
    WHERE "user_id" = OLD."id";
END;

