package websockets

import (
	"sync"

	"kood/social-network/pkg/api"
	"kood/social-network/pkg/db/queries"
	"kood/social-network/pkg/services"
	"net/http"
	"github.com/gorilla/websocket"
	"time"
	"encoding/json"
)

type WebSocketManager struct {
	authenticatedClients map[*Client]*api.UserResponseInfo
	chatRooms            map[string]*ChatRoom
	mu                   sync.RWMutex
	userRepo             *queries.UsersRepository
}

func NewWebSocketManager(userRepo *queries.UsersRepository) *WebSocketManager {
	manager := &WebSocketManager{
		authenticatedClients: make(map[*Client]*api.UserResponseInfo),
		chatRooms:            make(map[string]*ChatRoom),
		userRepo:             userRepo,
	}
	go manager.startNotificationChecker()
	// go manager.startMessageChecker()
	return manager
}

func (manager *WebSocketManager) addClient(client *Client, user *api.UserResponseInfo) {
	manager.mu.Lock()
	defer manager.mu.Unlock()
	client.user = user
	manager.authenticatedClients[client] = user
}

func (manager *WebSocketManager) removeClient(client *Client) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	if client.user != nil {
		if err := manager.userRepo.UpdateLastActivity(client.user.ID); err != nil {
		}
		delete(manager.authenticatedClients, client)
	}

	for _, room := range manager.chatRooms {
		if room.clients[client] {
			room.removeClient(client)
		}
	}
}

func (manager *WebSocketManager) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	user, status := services.AuthenticateUser(r)
	if !status {
		conn.Close()
		return
	}

	client := &Client{
		conn:    conn,
		manager: manager,
		user:    user,
		stopper: make(chan struct{}),
	}

	manager.addClient(client, user)
	client.readMessages()
}

func (manager *WebSocketManager) startNotificationChecker() {
	ticker := time.NewTicker(notificationPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			manager.checkForNewNotifications()
		}
	}
}

// func (manager *WebSocketManager) startMessageChecker() {
// 	ticker := time.NewTicker(messagePeriod)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			manager.checkForNewMessages()
// 		}
// 	}
// }

func (manager *WebSocketManager) GetOnlineUsers(userIds []int) map[int]bool {
	manager.mu.RLock()
	defer manager.mu.RUnlock()

	onlineUsers := make(map[int]bool)
	for client := range manager.authenticatedClients {
		for _, userId := range userIds {
			if client.user.ID == userId {
				onlineUsers[userId] = true
			}
		}
	}

	return onlineUsers
}


// UserStatusHandler godoc
// @Summary Get online status of users
// @Description Get online status of users
// @Tags websockets
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param users body []api.UserResponseInfo true "Users"
// @Success 200 {object} map[int]bool
// @Failure 400 {object} api.ErrorDetails
// @Failure 500 {object} api.ErrorDetails
// @Router /users/status [post]
func (manager *WebSocketManager) UserStatusHandler(w http.ResponseWriter, r *http.Request) {
	var users []api.UserResponseInfo
	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var userIds []int
	for _, user := range users {
		userIds = append(userIds, user.ID)
	}

	onlineUsers := manager.GetOnlineUsers(userIds)

	response := make(map[int]bool)
	for _, user := range users {
		response[user.ID] = onlineUsers[user.ID]
	}

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}