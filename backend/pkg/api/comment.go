package api

import (
	"time"
)

type CommentCreateRequest struct {
	PostID    int       `json:"postId" example:"1"`
	Content   string    `json:"content" example:"This is a comment"`
	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
}

type Comment struct {
	ID int `json:"id"`
	CommentCreateRequest
	UserResponse UserResponseInfo    `json:"user"`
	ReactionInfo ReactionInfo `json:"reactionInfo"`
}

type CommentCreateResponse struct {
	Comment Comment `json:"comments"`
}

type CommentListResponse struct {
	Comments []Comment `json:"comments"`
}
