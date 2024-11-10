package websockets

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"kood/social-network/pkg/api"
	"kood/social-network/pkg/db/queries"
)

type Client struct {
	conn           *websocket.Conn
	manager        *WebSocketManager
	user           *api.UserResponseInfo
	stopper        chan struct{}
	currentRoomHash string 
}

func (c *Client) String() string {
	return c.conn.RemoteAddr().String()
}

func (c *Client) getCurrentRoomHash() string {
	return c.currentRoomHash
}

func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClient(c)
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		var msg api.MessageResponse
		if err := json.Unmarshal(message, &msg); err != nil {
			c.sendError("Invalid message format")
			continue
		}

		c.handleMessageType(msg)
	}
}

func (c *Client) handleMessageType(msg api.MessageResponse) {
	payload, err := json.Marshal(msg.Payload)
	if err != nil {
		c.sendError("Error marshalling payload")
		return
	}
	payload, ok := msg.Payload.(json.RawMessage)
	if !ok {
		c.sendError("Invalid payload type")
		return
	}
	switch msg.Type {
	case "joinChat":
		c.handleJoinChat(payload)
	case "sendMessage":
		c.handleSendMessage(payload)
	default:
		c.sendError("Unknown message type")
	}
}

func (c *Client) handleJoinChat(payload json.RawMessage) {
	var joinRequest api.CreateChatResponse
	if err := json.Unmarshal(payload, &joinRequest); err != nil {
		c.sendError("Invalid join chat payload")
		return
	}

	c.manager.mu.Lock()
	defer c.manager.mu.Unlock()

	if c.currentRoomHash != "" {
		chatRoom, exists := c.manager.chatRooms[c.currentRoomHash]
		if exists {
			chatRoom.removeClient(c)
		}
	}

	chatRoom, exists := c.manager.chatRooms[joinRequest.Hash]
	if !exists {
		chatRoom = &ChatRoom{clients: make(map[*Client]bool)}
		c.manager.chatRooms[joinRequest.Hash] = chatRoom
	}

	chatRoom.addClient(c)
	c.currentRoomHash = joinRequest.Hash 

	messages, err := queries.NewChatsRepository().GetMessages(c.user.ID, joinRequest.Hash)
	if err != nil {
		c.sendError("Error fetching chat messages")
		return
	}

	response := api.MessageResponse{
		Type:    "chatContent",
		Payload: api.OpenChat{Hash: joinRequest.Hash, Messages: messages},
	}

	data, err := json.Marshal(response)
	if err != nil {
		c.sendError("Error creating chat content response")
		return
	}

	if err := c.sendMessage(data); err != nil {
	}
}

func (c *Client) handleSendMessage(payload json.RawMessage) {
	var sendMessageRequest api.SendMessageRequest
	if err := json.Unmarshal(payload, &sendMessageRequest); err != nil {
		c.sendError("Invalid send message payload")
		return
	}

	c.manager.mu.RLock()
	chatRoom, exists := c.manager.chatRooms[sendMessageRequest.ChatHash]
	c.manager.mu.RUnlock()

	if !exists {
		c.sendError("Chat room does not exist")
		return
	}
	err := queries.NewChatsRepository().SaveMessage(c.user.ID, &sendMessageRequest)
	if err != nil {
		c.sendError("Error saving message")
		return
	}

	response := api.MessageResponse{
		Type:    "newMessage",
		Payload: sendMessageRequest,
	}

	chatRoom.broadcastMessage(response)
}

func (c *Client) sendMessage(message []byte) error {
	if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
		return err
	}
	return nil
}

func (c *Client) sendError(errorMessage string) {
	msg := api.MessageResponse{
		Type:    "error",
		Payload: api.ErrorMessage{Message: errorMessage},
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return
	}

	if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
	}
}
