package api

import "time"

type GroupResponseInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GroupCreateRequest struct {
	Title       string `json:"title" example:"Test title"`
	Description string `json:"description" example:"Test description"`
	BannerColor string `json:"bannerColor" example:"#000000"`
}

type MemberInfo struct {
	TotalMembers int    `json:"totalMembers"`
	Roles        string `json:"roles"`
}

type GroupMember struct {
	UserResponseInfo
	Role string `json:"role"`
}

type Group struct {
	ID int `json:"id"`
	GroupCreateRequest
	CreatorInfo UserResponseInfo `json:"creatorInfo"`
	CreatedAt   time.Time        `json:"createdAt" `
	MemberInfo  MemberInfo       `json:"memberInfo"`
	Members     []GroupMember    `json:"members"`
	Posts       []Post           `json:"posts"`
	Events      []GroupEvent     `json:"events"`
	ChatHash    string           `json:"chatHash"`
}

type GroupResponse struct {
	Group Group `json:"group"`
}

type GroupListResponse struct {
	Groups []Group `json:"groups"`
}

type GroupRequest struct {
	GroupID     int    `json:"groupId"`
	RequestType string `json:"requestType" example:"inv"` // "inv" or "j_req"
}

type GroupEventCreate struct {
	GroupID     int       `json:"groupId" example:"1"`
	Title       string    `json:"title" example:"Test title"`
	Description string    `json:"description" example:"Test description"`
	Date        time.Time `json:"date" example:"2021-01-01T00:00:00Z"`
	CreatedAt   time.Time `json:"createdAt" swaggerignore:"true"`
}

type GroupEvent struct {
	ID          int              `json:"id"`
	CreatorInfo UserResponseInfo `json:"creatorInfo"`
	GroupEventCreate
	UsersGoing int                `json:"usersGoing"`
	Members    []UserResponseInfo `json:"members"`
}
