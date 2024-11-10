package main

import (
	"kood/social-network/config"
	"kood/social-network/pkg/db/queries"
	"kood/social-network/pkg/db/sqlite"
	"kood/social-network/pkg/handlers"
	"kood/social-network/pkg/middleware"
	"log"
	"net/http"
	// "os"
	// "path/filepath"
	// "strings"

	"github.com/gorilla/mux"
	"kood/social-network/pkg/websockets"
)

// @swagger 2.0
// @title           kood-social-network API
// @version         0.1.0
// @description     This is a sample server for kood-social-network.
// @host            localhost:8080
// @BasePath        /api
func main() {
	config.LoadConfig()
	dbWrapper := sqlite.NewDBWrapper()
	dbWrapper.InitDB(config.AppConfig.Database.Path, config.AppConfig.Database.MigrationsPath)
	defer dbWrapper.Close()
	queries.NewDBWrapper(dbWrapper.DB)

	r := mux.NewRouter()

	r.PathPrefix("/swagger/{any:.*}").Handler(handlers.SwaggerHandler())
	// r.PathPrefix("/api/uploads/").Handler(customFileServer("../backend/pkg/db/data/uploads"))

	// Define API routes
	api := r.PathPrefix("/api").Subrouter()

	// Authentication
	api.HandleFunc("/auth/register", handlers.HandleRegister).Methods("POST")
	api.HandleFunc("/auth/login", handlers.HandleLogin).Methods("POST")
	api.HandleFunc("/auth/logout", handlers.HandleLogout).Methods("DELETE")

	// Posts
	api.HandleFunc("/posts", handlers.HandleCreatePost).Methods("POST")
	api.HandleFunc("/posts", handlers.HandleGetPosts).Methods("GET")
	api.HandleFunc("/posts/{id:[0-9]+}", handlers.HandleGetPost).Methods("GET")

	// Comments
	api.HandleFunc("/comments", handlers.HandleCreateComment).Methods("POST")

	// Reactions
	api.HandleFunc("/reactions", handlers.HandleCreateReaction).Methods("PATCH")

	// Notifications
	api.HandleFunc("/notification/response", handlers.HandleNotificationResponse).Methods("POST")

	// Follow
	api.HandleFunc("/follow/request", handlers.HandleFollowRequest).Methods("POST")

	// Groups
	api.HandleFunc("/groups", handlers.HandleCreateGroup).Methods("POST")
	api.HandleFunc("/groups", handlers.HandleGetGroups).Methods("GET")
	api.HandleFunc("/groups/{id:[0-9]+}", handlers.HandleGetGroup).Methods("GET")
	api.HandleFunc("/groups/request", handlers.HandleGroupRequest).Methods("POST")
	api.HandleFunc("/groups/event", handlers.HandleCreateGroupEvent).Methods("POST")

	// Users
	api.HandleFunc("/users", handlers.HandleGetUsers).Methods("GET")
	api.HandleFunc("/users/{id:[0-9]+}", handlers.HandleGetUser).Methods("GET")

	// Chats
	api.HandleFunc("/chats", handlers.HandleCreateChat).Methods("POST")
	api.HandleFunc("/chats", handlers.HandleGetChats).Methods("GET")


	// Websockets
	manager := websockets.NewWebSocketManager(queries.NewUsersRepository())
	api.HandleFunc("/ws", manager.WebSocketHandler).Methods("GET")
	api.HandleFunc("/users/status", manager.UserStatusHandler).Methods("POST")


	
	// Start the server
	log.Printf("Server running on http://%s\n", config.AppConfig.PortNumber)
	log.Println("To stop the server press `Ctrl + C`")
	log.Fatal(http.ListenAndServe(config.AppConfig.PortNumber, middleware.ConfigureCORS(r)))
}

// func customFileServer(dir string) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		path := r.URL.Path
// 		if !strings.HasPrefix(path, "/api/uploads/") {
// 			http.NotFound(w, r)
// 			return
// 		}

// 		filePath := filepath.Join(dir, strings.TrimPrefix(path, "/api/uploads/"))

// 		fileInfo, err := os.Stat(filePath)
// 		if err != nil || fileInfo.IsDir() {
// 			http.NotFound(w, r)
// 			return
// 		}

// 		http.ServeFile(w, r, filePath)
// 	}
// }
