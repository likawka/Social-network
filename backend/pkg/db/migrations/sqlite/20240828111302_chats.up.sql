-- Table to manage conversations
CREATE TABLE IF NOT EXISTS "conversations" (
    "conversation_hash" TEXT PRIMARY KEY,
    "conversation_type" TEXT CHECK (conversation_type IN ('individual', 'group')) NOT NULL,
    "creator_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "last_message" TEXT,
    "last_message_sender_name" TEXT,
    "last_message_sent_at" TIMESTAMP
);

-- Create a messages table for all chats
CREATE TABLE IF NOT EXISTS "messages" (
    "message_hash" TEXT PRIMARY KEY,
    "conversation_hash" TEXT NOT NULL,
    "sender_id" INTEGER NOT NULL,
    "content" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("conversation_hash") REFERENCES "conversations" ("conversation_hash") ON DELETE CASCADE,
    FOREIGN KEY ("sender_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

-- Table to map users to conversations
CREATE TABLE IF NOT EXISTS "conversation_users" (
    "conversation_hash" TEXT NOT NULL,
    "user_id" INTEGER NOT NULL,
    "joined_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "last_activity" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("conversation_hash", "user_id"),
    FOREIGN KEY ("conversation_hash") REFERENCES "conversations" ("conversation_hash") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

-- Trigger to generate conversation hash
CREATE TRIGGER generate_conversation_hash
BEFORE INSERT ON conversations
FOR EACH ROW
WHEN NEW.conversation_hash IS NULL
BEGIN
    UPDATE conversations
    SET conversation_hash = lower(hex(randomblob(16)))
    WHERE rowid = NEW.rowid;
END;

-- Trigger to generate message hash
CREATE TRIGGER generate_message_hash
BEFORE INSERT ON messages
FOR EACH ROW
WHEN NEW.message_hash IS NULL
BEGIN
    UPDATE messages
    SET message_hash = lower(hex(randomblob(16)))
    WHERE rowid = NEW.rowid;
END;