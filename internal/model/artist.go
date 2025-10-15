package model

import (
	"time"

	"github.com/google/uuid"
)

type Artist struct {
	ID          uuid.UUID
	Name        string
	AvatarURL   string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
