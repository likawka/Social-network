package websockets

import (
	"encoding/json"
	"log"
	"time"

	"kood/social-network/pkg/api"
	"kood/social-network/pkg/db/queries"
)

const (
	notificationPeriod = 10 * time.Second
	messagePeriod      = 10 * time.Second
)

func (manager *WebSocketManager) checkForNewNotifications() {
	manager.mu.RLock()
	defer manager.mu.RUnlock()

	for client, user := range manager.authenticatedClients {
		count, err := queries.NewNotificationsRepository().GetUnreadNotificationsCount(user.ID)
		if err != nil {
			log.Printf("Error fetching unread notifications count for user %d: %v", user.ID, err)
			continue
		}

		if count > 0 {
			info := api.UpdateNotification{
				Count:   count,
				Message: "You have new notifications.",
			}

			response := api.MessageResponse{
				Type:    "notificationUpdate",
				Payload: info,
			}

			data, err := json.Marshal(response)
			if err != nil {
				log.Printf("Error marshalling notification response for user %d: %v", user.ID, err)
				continue
			}

			if err := client.sendMessage(data); err != nil {
				log.Printf("Error sending notification update to client %s: %v", client.String(), err)
			}
		}
	}
}


// func (manager *WebSocketManager) checkForNewMessages() {
// 	manager.mu.RLock()
// 	defer manager.mu.RUnlock()

// 	for client, user := range manager.authenticatedClients {
// 		count, err := queries.NewChatsRepository().GetUnreadChatCount(user.ID)
// 		if err != nil {
// 			log.Printf("Error fetching unread messages count for user %d: %v", user.ID, err)
// 			continue
// 		}

// 		if count > 0 {
// 			info := api.UpdateNotification{
// 				Count:   count,
// 				Message: "You have new notifications.",
// 			}

// 			response := api.MessageResponse{
// 				Type:    "chatUpdate",
// 				Payload: info,
// 			}

// 			data, err := json.Marshal(response)
// 			if err != nil {
// 				log.Printf("Error marshalling message update response for user %d: %v", user.ID, err)
// 				continue
// 			}

// 			if err := client.sendMessage(data); err != nil {
// 				log.Printf("Error sending message update to client %s: %v", client.String(), err)
// 			}
// 		}
// 	}
// }