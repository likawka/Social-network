package services

import (
	"encoding/json"
	"kood/social-network/pkg/api"
	"net/http"
	"time"
)

var metadata = api.Metadata{
	Timestamp: time.Now().UTC(),
	Version:   "0.1.0",
}

func HTTPError(w http.ResponseWriter, status int, userMessage, logMessage string, authenticated bool, user *api.UserResponseInfo, fieldErrors []api.ValidationError) {
	response := api.Response{
		Status:  "error",
		Message: userMessage,
		Error: &api.ErrorDetails{
			Code:    status,
			Message: logMessage,
			Details: fieldErrors,
		},
		Authenticated: authenticated,
		User:          user,
		Metadata:      metadata,
	}
	RespondWithJSON(w, status, response)
}

func RespondWithSuccess(w http.ResponseWriter, status int, message string, authenticated bool, payload, pagination interface{}, user *api.UserResponseInfo) {
	response := api.Response{
		Status:        "success",
		Message:       message,
		Payload:       payload,
		Pagination:    pagination,
		Authenticated: authenticated,
		User:          user,
		Metadata:      metadata,
	}
	RespondWithJSON(w, status, response)
}

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
