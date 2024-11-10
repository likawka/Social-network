package handlers

import (
	"encoding/json"
	"kood/social-network/pkg/api"
	"kood/social-network/pkg/db/queries"
	"kood/social-network/pkg/services"
	"net/http"
	"strconv"
)

// HandleCreateGroup creates a new group.
// @Summary Create a new group
// @Description Creates a new group with the provided title and description.
// @Tags groups
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Param GroupCreateRequest body api.GroupCreateRequest true "Group details"
// @Response 200 {object} api.Response{payload=api.ResponseID}
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /groups [post]
func HandleCreateGroup(w http.ResponseWriter, r *http.Request) {
	user, authenticated := services.AuthenticateUser(r)
	if !authenticated {
		services.HTTPError(w, http.StatusUnauthorized, "Unauthorized", "User is not authenticated", false, user, nil)
		return
	}

	var request api.GroupCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		services.HTTPError(w, http.StatusBadRequest, "Invalid request", "Failed to parse request body", false, user, nil)
		return
	}

	request.Title = services.TrimAndNormalizeSpaces(request.Title)
	request.Description = services.TrimAndNormalizeSpaces(request.Description)

	validationErrors := services.ValidateOperation("group", request)
	if len(validationErrors) > 0 {
		services.HTTPError(w, http.StatusBadRequest, "Validation error", "Validation error", authenticated, user, validationErrors)
		return
	}
	groupRepo := queries.NewGroupsRepository()
	id, err := groupRepo.Create(user.ID, &request)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), authenticated, user, nil)
		return
	}

	payload := api.ResponseID{ID: id}

	services.RespondWithSuccess(w, http.StatusCreated, "Group created successfully", authenticated, payload, nil, user)

}

// HandleGetGroups fetches all groups.
// @Summary Get all groups
// @Description Fetches all groups.
// @Tags groups
// @Accept  json
// @Produce json
// @Response 200 {object} api.Response{payload=api.CommentCreateResponse}
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /groups [get]
func HandleGetGroups(w http.ResponseWriter, r *http.Request) {
	user, authenticated := services.AuthenticateUser(r)
	groupRepo := queries.NewGroupsRepository()
	groups, err := groupRepo.GetGroups(user.ID)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), authenticated, user, nil)
		return
	}
	payload := api.GroupListResponse{Groups: groups}
	services.RespondWithSuccess(w, http.StatusOK, "Groups fetched successfully", authenticated, payload, nil, user)
}

// HandleGroupRequest processes a group request.
// @Summary Process a group request
// @Description Processes a group request, either accepting or rejecting it.
// @Tags groups
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Param GroupRequest body api.GroupRequest true "Group request details"
// @Response 200 {object} api.Response
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /groups/request [post]
func HandleGroupRequest(w http.ResponseWriter, r *http.Request) {
	authenticated := false
	var user *api.UserResponseInfo

	user, authenticated = services.AuthenticateUser(r)
	if !authenticated {
		services.HTTPError(w, http.StatusUnauthorized, "Unauthorized", "User is not authenticated", false, nil, nil)
		return
	}

	var requestBody api.GroupRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		services.HTTPError(w, http.StatusBadRequest, "Invalid request payload", "Failed to parse request", authenticated, nil, nil)
		return
	}

	repo := queries.NewGroupsRepository()
	err := repo.GroupRequest(user.ID, &requestBody)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Failed to process request", err.Error(), authenticated, nil, nil)
		return
	}

	services.RespondWithSuccess(w, http.StatusOK, "Request processed successfully", authenticated, nil, nil, user)
}

// HandleCreateGroupEvent creates a new group event.
// @Summary Create a new group event
// @Description Creates a new group event with the provided title, description, and date.
// @Tags groups
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Param GroupEventCreateRequest body api.GroupEventCreate true "Group event details"
// @Response 200 {object} api.Response{payload=api.GroupEvent}
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /groups/event [post]
func HandleCreateGroupEvent(w http.ResponseWriter, r *http.Request) {
	var request api.GroupEventCreate

	user, authenticated := services.AuthenticateUser(r)
	if !authenticated {
		services.HTTPError(w, http.StatusUnauthorized, "Unauthorized", "User is not authenticated", false, nil, nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		services.HTTPError(w, http.StatusBadRequest, "Invalid request", "Failed to parse request body", false, nil, nil)
		return
	}

	request.Title = services.TrimAndNormalizeSpaces(request.Title)
	request.Description = services.TrimAndNormalizeSpaces(request.Description)

	validationErrors := services.ValidateOperation("event", request)
	if len(validationErrors) > 0 {
		services.HTTPError(w, http.StatusBadRequest, "Validation error", "Validation error", authenticated, nil, validationErrors)
		return
	}
	groupRepo := queries.NewGroupsRepository()

	eventID, err := groupRepo.CreateGroupEvent(user.ID, &request)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), authenticated, user, nil)
		return
	}

	payload := api.GroupEvent{
		ID:               eventID,
		GroupEventCreate: request,
		CreatorInfo:      *user,
	}

	services.RespondWithSuccess(w, http.StatusCreated, "Group event created successfully", authenticated, payload, nil, user)
}

// HandleGetGroup fetches a group by ID.
// @Summary Get a group by ID
// @Description Fetches a group by its ID.
// @Tags groups
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Group ID"
// @Response 200 {object} api.Response{payload=api.GroupResponse}
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /groups/{id} [get]
func HandleGetGroup(w http.ResponseWriter, r *http.Request) {
	user, authenticated := services.AuthenticateUser(r)
	if !authenticated {
		services.HTTPError(w, http.StatusUnauthorized, "Unauthorized", "User is not authenticated", false, user, nil)
		return
	}

	params := services.GetRouteParams(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		services.HTTPError(w, http.StatusBadRequest, "Bad Request", "Invalid post ID", authenticated, user, nil)
		return
	}
	var group *api.Group
	groupRepo := queries.NewGroupsRepository()
	group, err = groupRepo.GetGroup(id)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), authenticated, user, nil)
		return
	}
	if group == nil {
		services.HTTPError(w, http.StatusNotFound, "Not Found", "Group not found", authenticated, user, nil)
		return
	}
	group.ID = id
	if err := groupRepo.GetGroupMembers(group); err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), authenticated, user, nil)
		return
	}

	if err := groupRepo.GetGroupEvents(group); err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), authenticated, user, nil)
		return
	}

	postsRepo := queries.NewPostsRepository()
	posts, err := postsRepo.GetPostsForGroup(id)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), authenticated, user, nil)
		return
	}
	group.Posts = posts
	if is_member, _ := groupRepo.IsUserMemberOfGroup(user.ID, id); is_member {
		group.ChatHash, _ = queries.NewChatsRepository().GetChatHashForGroup(id)
	}

	payload := api.GroupResponse{Group: *group}
	services.RespondWithSuccess(w, http.StatusOK, "Group fetched successfully", authenticated, payload, nil, user)
}
