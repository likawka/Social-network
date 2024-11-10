package queries

import (
	"database/sql"
	"kood/social-network/pkg/api"
	"log"
	"time"
)

type ChatsRepository struct {
	DB *sql.DB
}

func NewChatsRepository() *ChatsRepository {
	return &ChatsRepository{DB: DBWrapper}
}

func (r *ChatsRepository) GetChatHashForGroup(groupID int) (string, error) {
	var hash string
	err := r.DB.QueryRow("SELECT conversation_hash FROM conversations WHERE creator_id = ? AND conversation_type = 'group'", groupID).Scan(&hash)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return hash, nil
}

func (r *ChatsRepository) GetChatHashForUser(userID int) (string, error) {
	var hash string
	err := r.DB.QueryRow("SELECT conversation_hash FROM conversations WHERE creator_id = ? AND conversation_type = 'user'", userID).Scan(&hash)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return hash, nil
}

func (r *ChatsRepository) IndividualChats(userID int) ([]api.Chat, error) {
	rows, err := r.DB.Query(`
		SELECT c.conversation_hash, c.conversation_type, c.created_at 
		FROM conversations c
		JOIN conversation_users cu ON cu.conversation_hash = c.conversation_hash
		WHERE cu.user_id = ? AND c.conversation_type = 'user'
	`, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var chats []api.Chat
	for rows.Next() {
		var chat api.Chat
		err := rows.Scan(&chat.Hash, &chat.ChatType, &chat.CreatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		userRows, err := r.DB.Query(`
			SELECT u.id, u.nickname 
			FROM users u
			JOIN conversation_users cu ON cu.user_id = u.id
			WHERE cu.conversation_hash = ?
		`, chat.Hash)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		defer userRows.Close()

		var users []api.UserResponseInfo
		for userRows.Next() {
			var user api.UserResponseInfo
			err := userRows.Scan(&user.ID, &user.Nickname)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			users = append(users, user)

			if user.ID != userID {
				chat.ChatName = user.Nickname
			}
		}
		chat.Users = users

		chats = append(chats, chat)
	}

	return chats, nil
}
func (r *ChatsRepository) GroupChats(userID int) ([]api.Chat, error) {
	rows, err := r.DB.Query(`
		SELECT c.conversation_hash, c.conversation_type, g.title AS chat_name, c.created_at, c.last_message, c.last_message_sender_name, c.last_message_sent_at
		FROM conversations c
		JOIN conversation_users cu ON cu.conversation_hash = c.conversation_hash
		JOIN groups g ON g.id = c.creator_id
		WHERE cu.user_id = ? AND c.conversation_type = 'group'
	`, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var chats []api.Chat
	for rows.Next() {
		var chat api.Chat
		var lastMessage sql.NullString
		var lastSenderName sql.NullString
		var lastMessageSentAt sql.NullTime

		err := rows.Scan(&chat.Hash, &chat.ChatType, &chat.ChatName, &chat.CreatedAt, &lastMessage, &lastSenderName, &lastMessageSentAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		chat.LastMasege = lastMessage.String
		chat.LastSender.Nickname = lastSenderName.String
		if lastMessageSentAt.Valid {
			chat.LastMasegTime = lastMessageSentAt.Time
		} else {
			chat.LastMasegTime = time.Time{}
		}

		userRows, err := r.DB.Query(`
			SELECT u.id, u.nickname 
			FROM users u
			JOIN conversation_users cu ON cu.user_id = u.id
			WHERE cu.conversation_hash = ?
		`, chat.Hash)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		defer userRows.Close()

		var users []api.UserResponseInfo
		for userRows.Next() {
			var user api.UserResponseInfo
			err := userRows.Scan(&user.ID, &user.Nickname)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			users = append(users, user)
		}
		chat.Users = users

		chats = append(chats, chat)
	}

	return chats, nil
}

func (r *ChatsRepository) Create(userID, memberID int) (string, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var chatHash string

	err = tx.QueryRow(`
		INSERT INTO conversations (conversation_type, creator_id)
		VALUES ('individual', ?)
		RETURNING conversation_hash
	`, userID).Scan(&chatHash)

	if err != nil {
		return "", err
	}

	_, err = tx.Exec(`
		INSERT INTO conversation_users (conversation_hash, user_id)
		VALUES (?, ?), (?, ?)
	`, chatHash, userID, chatHash, memberID)

	if err != nil {
		return "", err
	}

	return chatHash, nil
}

func (r *ChatsRepository) SaveMessage(senderID int, info *api.SendMessageRequest) error {

	query := `
		INSERT INTO messages (conversation_hash, sender_id, content)
		VALUES (?, ?, ?)
		RETURNING message_hash;
	`

	err := r.DB.QueryRow(query, info.ChatHash, senderID, info.Pyload.Content).Scan(&info.Pyload.Hash)
	if err != nil {
		return err
	}
	info.Pyload.CreatedAt = time.Now().UTC()
	return nil
}

func (r *ChatsRepository) GetMessages(userID int,  conversationHash string) ([]api.Message, error) {
	rows, err := r.DB.Query(`
		SELECT m.message_hash, m.sender_id, u.nickname, m.content, m.created_at
		FROM messages m
		JOIN users u ON m.sender_id = u.id
		WHERE m.conversation_hash = ?
	`, conversationHash)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var messages []api.Message
	for rows.Next() {
		var message api.Message
		err := rows.Scan(&message.Hash, &message.Sender.ID, &message.Sender.Nickname, &message.Content, &message.CreatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}


