package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	AccessToken  string
	TokenType    string
	ExpiresIn    int
	ExpiresAt    int64
	RefreshToken string
	User         User
}

type User struct {
	ID               uuid.UUID `json:"id"`
	Aud              string    `json:"aud"`
	Role             string    `json:"role"`
	Email            string    `json:"email"`
	DisplayName      string    `json:"display_name"`
	EmailConfirmedAt time.Time `json:"email_confirmed_at"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	AppMetadata      map[string]any
	UserMetadata     map[string]any
}
