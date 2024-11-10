package handlers

import (
	"encoding/json"
	"kood/social-network/pkg/api"
	"kood/social-network/pkg/db/queries"
	"kood/social-network/pkg/services"
	"net/http"
)

// HandleCreateComment godoc
// @Summary Create a new comment on a post
// @Description Creates a new comment for a specified post
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body api.CommentCreateRequest true "Comment Data"
// @Security BearerAuth
// @Response 200 {object} api.Response{payload=api.CommentCreateResponse}
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /comments [post]
func HandleCreateComment(w http.ResponseWriter, r *http.Request) {
	authenticated := false
	var user *api.UserResponseInfo

	user, authenticated = services.AuthenticateUser(r)
	if !authenticated {
		services.HTTPError(w, http.StatusUnauthorized, "Unauthorized", "User is not authenticated", false, user, nil)
		return
	}

	var commentForm api.CommentCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&commentForm); err != nil {
		services.HTTPError(w, http.StatusBadRequest, "Bad Request", "Invalid request payload", authenticated, user, nil)
		return
	}

	commentForm.Content = services.TrimAndNormalizeSpaces(commentForm.Content)

	validationErrors := services.ValidateOperation("comment", commentForm)
	if len(validationErrors) > 0 {
		services.HTTPError(w, http.StatusBadRequest, "Validation error", "Validation error", authenticated, user, validationErrors)
		return
	}

	commentRepo := queries.NewCommentsRepository()
	commentID, err := commentRepo.Create(user.ID, &commentForm)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), authenticated, user, nil)
		return
	}

	response := api.Comment{
		ID:                   commentID,
		CommentCreateRequest: commentForm,
		UserResponse:         *user,
	}
	services.RespondWithSuccess(w, http.StatusCreated, "Comment created successfully", authenticated, api.CommentCreateResponse{response}, nil, user)
}
