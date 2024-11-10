package handlers

import (
	"encoding/json"
	"kood/social-network/pkg/api"
	"kood/social-network/pkg/db/queries"
	"kood/social-network/pkg/services"
	"net/http"
)

// HandleFollowRequest manages follow and unfollow actions
// @Summary Manage follow and unfollow actions
// @Description Processes follow or unfollow requests based on the provided action type and followee ID.
// @Tags follow
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param FollowRequest body api.FollowRequest true "Follow or Unfollow request details"
// @Response 200 {object} api.Response
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /follow/request [post]
func HandleFollowRequest(w http.ResponseWriter, r *http.Request) {
	authenticated := false
	var user *api.UserResponseInfo

	user, authenticated = services.AuthenticateUser(r)
	if !authenticated {
		services.HTTPError(w, http.StatusUnauthorized, "Unauthorized", "User is not authenticated", false, user, nil)
		return
	}

	var requestBody api.FollowRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		services.HTTPError(w, http.StatusBadRequest, "Invalid request payload", "Failed to parse request", authenticated, user, nil)
		return
	}

	repo := queries.NewFollowsRepository()

	var err error
	switch requestBody.Type {
	case "follow":
		err = repo.CreateFollowRequest(user.ID, requestBody.FolloweeID)
	case "unfollow":
		err = repo.RemoveFollow(user.ID, requestBody.FolloweeID)
	default:
		services.HTTPError(w, http.StatusBadRequest, "Invalid action type", "Follow type must be 'follow' or 'unfollow'", authenticated, user, nil)
		return
	}

	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Failed to process request", err.Error(), authenticated, user, nil)
		return
	}
	services.RespondWithSuccess(w, http.StatusOK, "Follow action processed successfully", authenticated, nil, nil, user)
}

