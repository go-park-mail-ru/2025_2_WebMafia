package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Login        string
	Email        string
	PasswordHash string
	AvatarURL    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
