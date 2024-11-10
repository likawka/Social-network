package websockets

import (
	"encoding/json"
	"sync"

	"kood/social-network/pkg/api"
)

type ChatRoom struct {
	clients map[*Client]bool
	mu      sync.RWMutex
}

func (room *ChatRoom) addClient(client *Client) {
	room.mu.Lock()
	defer room.mu.Unlock()
	room.clients[client] = true
}

func (room *ChatRoom) removeClient(client *Client) {
	room.mu.Lock()
	defer room.mu.Unlock()

	if room.clients[client] {
		delete(room.clients, client)

		if len(room.clients) == 0 {
			for hash, r := range client.manager.chatRooms {
				if r == room {
					delete(client.manager.chatRooms, hash)
					break
				}
			}
		}
	}
}

func (room *ChatRoom) broadcastMessage(message api.MessageResponse) {
	room.mu.RLock()
	defer room.mu.RUnlock()

	data, err := json.Marshal(message)
	if err != nil {
		return
	}

	for client := range room.clients {
		if err := client.sendMessage(data); err != nil {
		}
	}
}
