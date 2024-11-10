package handlers

import (
	"encoding/json"
	"kood/social-network/pkg/api"
	"kood/social-network/pkg/db/queries"
	"kood/social-network/pkg/services"
	"net/http"
)

// HandleCreateReaction godoc
// @Summary Create or update a reaction (like/dislike) on a post or comment
// @Description Allows a user to create or update a reaction (like or dislike) for a specified post or comment
// @Tags reactions
// @Accept json
// @Produce json
// @Param reaction body api.ReactionCreateRequest true "Reaction Data"
// @Security BearerAuth
// @Response 200 {object} api.Response{payload=api.ReactionCreateResponse}
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /reactions [PATCH]
func HandleCreateReaction(w http.ResponseWriter, r *http.Request) {
	user, authenticated := services.AuthenticateUser(r)
	if !authenticated {
		services.HTTPError(w, http.StatusUnauthorized, "Unauthorized", "User is not authenticated", false, user, nil)
		return
	}

	var request api.ReactionCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		services.HTTPError(w, http.StatusBadRequest, "Invalid request", "Failed to parse request body", false, user, nil)
		return
	}

	if request.ObjectType != "post" && request.ObjectType != "comment" {
		services.HTTPError(w, http.StatusBadRequest, "Invalid object type", "Object type must be 'post' or 'comment'", false, user, nil)
		return
	}
	if request.ReactionType != "like" && request.ReactionType != "dislike" {
		services.HTTPError(w, http.StatusBadRequest, "Invalid reaction type", "Reaction type must be 'like' or 'dislike'", false, user, nil)
		return
	}

	reactionRepo := queries.NewReactionsRepository()

	err := reactionRepo.CreateOrUpdateReaction(user.ID, request.ObjectType, request.ObjectID, request.ReactionType)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), authenticated, user, nil)
		return
	}

	likes, dislikes, err := reactionRepo.GetReactionCounts(request.ObjectType, request.ObjectID)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), authenticated, user, nil)
		return
	}
	payload := api.ReactionCreateResponse{ReactionInfo: api.ReactionInfo{LikeCount: likes, DislikeCount: dislikes, Status: request.ReactionType}}
	services.RespondWithSuccess(w, http.StatusCreated, "Reaction created successfully", authenticated, payload, nil, user)
}
