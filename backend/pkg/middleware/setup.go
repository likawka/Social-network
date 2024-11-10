package middleware

import (
  "github.com/gorilla/handlers"
	"net/http"
)

func ConfigureCORS(handler http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)(handler)
}
