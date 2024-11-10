package api

import (
	"time"
)

type PostCreateRequest struct {
	Groupinfo    GroupResponseInfo   `json:"group"`
	Title     string    `json:"title" example:"Test title"`
	Content   string    `json:"content" example:"Test content"`
	Image     string    `json:"image"`
	Privacy   string    `json:"privacy" example:"public"`
	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
}

type Post struct {
	ID           int                 `json:"id"`
	Userinfo     UserResponseInfo    `json:"user"`
	PostCreateRequest
	CommentCount int                 `json:"commentCount"`
	ReactionInfo ReactionInfo `json:"reactionInfo"`
}

type PostsListResponse struct {
	Posts []Post `json:"posts"`
}

type PostResponse struct {
	Post     *Post      `json:"post"`
	Comments *[]Comment `json:"comments"`
}
