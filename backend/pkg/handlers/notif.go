package handlers

import (
	"encoding/json"
	"kood/social-network/pkg/api"
	"kood/social-network/pkg/db/queries"
	"kood/social-network/pkg/services"
	"net/http"
)

// HandleNotificationResponse processes the notification response from the authenticated user.
// @Summary Process a notification response
// @Description Processes the response to a notification, which could involve actions such as accepting or rejecting various types of notifications.
// @Tags notifications
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param NotificationResponse body api.NotificationResponse true "Notification response details"
// @Success 200 {object} api.Response
// @Router /notification/response [post]
func HandleNotificationResponse(w http.ResponseWriter, r *http.Request) {
	authenticated := false
	var user *api.UserResponseInfo

	user, authenticated = services.AuthenticateUser(r)
	if !authenticated {
		services.HTTPError(w, http.StatusUnauthorized, "Unauthorized", "User is not authenticated", false, user, nil)
		return
	}

	var requestBody api.NotificationResponse
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		services.HTTPError(w, http.StatusBadRequest, "Invalid request payload", "Failed to parse request", authenticated, user, nil)
		return
	}

	var table string
	switch requestBody.NotificationInfo.Type {
	case "f_req":
		table = "followers"
	case "g_req", "g_inv":
		table = "group_members"
	case "g_eve":
		table = "group_events"
	default:
		services.HTTPError(w, http.StatusBadRequest, "Invalid notification type", "Invalid notification type", authenticated, user, nil)
		return
	}
	repo := queries.NewFollowsRepository()
	err := repo.ProcessNotificationResponse(user.ID, &requestBody, table)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Failed to process follow response", err.Error(), authenticated, user, nil)
		return
	}
	services.RespondWithSuccess(w, http.StatusOK, "Follow response processed successfully", authenticated, nil, nil, user)
}
