package handlers

import (
	"kood/social-network/pkg/api"
	"kood/social-network/pkg/db/queries"
	"kood/social-network/pkg/services"
	"net/http"
	"strconv"
)

// HandleGetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Response 200 {object} api.Response{payload=api.UsersList}
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /users [get]
func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	authenticated := false
	var user *api.UserResponseInfo

	user, authenticated = services.AuthenticateUser(r)
	if !authenticated {
		services.HTTPError(w, http.StatusUnauthorized, "Unauthorized", "User is not authenticated", false, user, nil)
		return
	}
	var users []api.User
	repo := queries.NewUsersRepository()

	if err := repo.GetUsers(user.ID, &users); err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Failed to get users", err.Error(), authenticated, user, nil)
		return
	}
	payload := api.UsersList{users}
	services.RespondWithSuccess(w, http.StatusOK, "Users retrieved successfully", authenticated, payload, nil, user)
}

// HandleGetUser godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Response 200 {object} api.Response{payload=api.UserPage}
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /users/{id} [get]
func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	authenticated := false
	var user *api.UserResponseInfo

	user, authenticated = services.AuthenticateUser(r)
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
	var userPage *api.UserPage
	repoUser := queries.NewUsersRepository()
	userPage, err = repoUser.GetUserByID(user.ID, id)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Failed to get user", err.Error(), authenticated, user, nil)
		return
	}
	if userPage == nil {
		services.HTTPError(w, http.StatusNotFound, "Not Found", "User not found", authenticated, user, nil)
		return
	}
	userPage.FolowStatus = queries.NewFollowsRepository().CheckFollow(user.ID, id)
	userPage.Following, _ = queries.NewFollowsRepository().GetFollows(id, "following")
	userPage.Followers, _ = queries.NewFollowsRepository().GetFollows(id, "followers")
	userPage.PersonalPosts, _ = queries.NewPostsRepository().GetPostsForUser(user.ID, id)
	userPage.GroupsPosts, _ = queries.NewPostsRepository().GetGroupPostsForUser(user.ID, id)
	if is_member := queries.NewFollowsRepository().CheckFollow(user.ID, id); is_member {
		userPage.ChatHash, _ = queries.NewChatsRepository().GetChatHashForUser(id)
	}
	payload := userPage
	services.RespondWithSuccess(w, http.StatusOK, "User retrieved successfully", authenticated, payload, nil, user)
}
