package api

import "time"

type ReactionCreateRequest struct {
	ObjectType   string `json:"objectType" example:"post"`
	ObjectID     int    `json:"objectId" example:"1"`
	ReactionType string `json:"reactionType" example:"like"`
}

type ReactionCreateResponse struct {
	ReactionInfo ReactionInfo `json:"reactionInfo"`
}

type Reaction struct {
	UserId       int       `json:"userId"`
	ReactionCreateRequest
	CreatedAt    time.Time `json:"createdAt"`
}

type ReactionInfo struct {
	Status       string `json:"status"`
	LikeCount    int    `json:"likeCount"`
	DislikeCount int    `json:"dislikeCount"`
}
