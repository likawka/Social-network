package api

import (
	"time"
)

type RegistrationRequest struct {
	FirstName string ` json:"firstName" example:"John"`
	LastName  string ` json:"lastName" example:"Doe"`
	LoginRequest
	DateOfBirth string    ` json:"dateOfBirth" example:"20.02.2002"`
	Nickname    string    `json:"nickname" example:"test"`
	AboutMe     string    ` json:"aboutMe" example:"I am a software engineer"`
	Avatar      string    `json:"image" example:""`
	CreatedAt   time.Time `json:"createdAt" swaggerignore:"true"`
}

type LoginRequest struct {
	Email    string `json:"email" example:"test@test.com"`
	Password string `json:"password" example:"!QAZ2wsx"`
}
