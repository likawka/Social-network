package handlers

import (
	"encoding/json"
	"kood/social-network/pkg/api"
	"kood/social-network/pkg/db/queries"
	"kood/social-network/pkg/services"
	"log"
	"net/http"
	"strconv"
)

// HandleCreatePost creates a new post.
// @Summary Create a new post
// @Description Creates a new post with the provided post details and optional image file.
// @Tags posts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body api.PostCreateRequest true "Post details"
// @Response 201 {object} api.Response{payload=api.ResponseID}
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /posts [post]
func HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	authenticated := false
	var user *api.UserResponseInfo

	user, authenticated = services.AuthenticateUser(r)
	if !authenticated {
		services.HTTPError(w, http.StatusUnauthorized, "Unauthorized", "User is not authenticated", false, user, nil)
		return
	}

	var postForm api.PostCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&postForm); err != nil {
		services.HTTPError(w, http.StatusBadRequest, "Invalid request payload", err.Error(), authenticated, nil, nil)
		return
	}
	postForm.Title = services.TrimAndNormalizeSpaces(postForm.Title)
	postForm.Content = services.TrimAndNormalizeSpaces(postForm.Content)

	validationErrors := services.ValidateOperation("post", postForm)
	if len(validationErrors) > 0 {
		services.HTTPError(w, http.StatusBadRequest, "Validation error", "Validation error", authenticated, user, validationErrors)
		return
	}
	postRepo := queries.NewPostsRepository()
	id, err := postRepo.Create(user.ID, &postForm)
	if err != nil {
		log.Println(err)
		services.HTTPError(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), authenticated, user, nil)
		return
	}
	payload := api.ResponseID{ID: id}
	services.RespondWithSuccess(w, http.StatusCreated, "Post created successfully", authenticated, payload, nil, user)
}

// HandleGetPosts godoc
// @Summary Get all posts for an authenticated user
// @Description Retrieve all public posts and those from users who have accepted follow requests
// @Tags posts
// @Accept json
// @Produce json
// @Response 200 {object} api.Response{payload=api.PostsListResponse}
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /posts [get]
func HandleGetPosts(w http.ResponseWriter, r *http.Request) {
	authenticated := false
	var user *api.UserResponseInfo

	user, authenticated = services.AuthenticateUser(r)
	var posts api.PostsListResponse
	postRepo := queries.NewPostsRepository()
	posts.Posts, _ = postRepo.GetPostsWithoutGroups(user.ID)
	payload := posts
	services.RespondWithSuccess(w, http.StatusOK, "Posts retrieved successfully", authenticated, payload, nil, user)
}

// HandleGetPost godoc
// @Summary Get a post by ID
// @Description Retrieve a post by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Response 200 {object} api.Response{payload=api.PostResponse}
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /posts/{id} [get]
func HandleGetPost(w http.ResponseWriter, r *http.Request) {
	authenticated := false
	var user *api.UserResponseInfo

	user, authenticated = services.AuthenticateUser(r)

	params := services.GetRouteParams(r)
	postID, err := strconv.Atoi(params["id"])
	if err != nil {
		services.HTTPError(w, http.StatusBadRequest, "Bad Request", "Invalid post ID", authenticated, user, nil)
		return
	}
	var post api.PostResponse
	post.Post, err = queries.NewPostsRepository().GetPostByID(user.ID, postID)
	if post.Post == nil {
		services.HTTPError(w, http.StatusNotFound, "Not Found", "Post not found", authenticated, user, nil)
		return
	}
	post.Comments, _ = queries.NewCommentsRepository().GetCommentsForPost(user.ID, postID)
	payload := post
	services.RespondWithSuccess(w, http.StatusOK, "Post retrieved successfully", authenticated, payload, nil, user)
}
