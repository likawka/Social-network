package models

import (
	"time"
)

type Session struct {
	UserID       int       `json:"user_id"`
	SessionID    string    `json:"session_id"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}
