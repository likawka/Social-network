package api

type FollowRequest struct {
	Type       string `json:"type" example:"follow"` // "follow" or "unfollow"
	FolloweeID int    `json:"followeeId" example:"1"`
}

type NotificationResponse struct {
	NotificationInfo NotificationInfo `json:"notification_info"`
	ReqestID         int              `json:"requestId" example:"1"`
	Status           string           `json:"status" example:"accepted"` // "accepted" or "rejected"
}
