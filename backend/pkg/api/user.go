package api

import "time"

type User struct {
	UserResponseInfo
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	DateOfBirth       string    `json:"dateOfBirth"`
	AvatarPath        string    `json:"avatar"`
	AboutMe           string    `json:"aboutMe"`
	BannerColor       string    `json:"bannerColor"`
	ProfileVisibility string    `json:"profileVisibility"`
	CreatedAt         time.Time `json:"createdAt"`
	PostCount         int       `json:"postCount"`
	CommentCount      int       `json:"commentCount"`
	FollowerCount     int       `json:"followerCount"`
	FollowingCount    int       `json:"followingCount"`
	LastActive        time.Time `json:"lastActive"`
}

type UserResponseInfo struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

type UsersList struct {
	Users []User `json:"users"`
}

type UserPage struct {
	User          User               `json:"user"`
	FolowStatus   bool               `json:"followStatus"`
	PersonalPosts []Post             `json:"personalPosts"`
	GroupsPosts   []Post             `json:"groupsPosts"`
	Followers     []UserResponseInfo `json:"followers"`
	Following     []UserResponseInfo `json:"following"`
	ChatHash      string             `json:"chatHash"`
}
