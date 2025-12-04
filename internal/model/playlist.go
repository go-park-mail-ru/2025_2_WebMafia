package model

import (
	"github.com/google/uuid"
	"time"
)

type Playlist struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Title       string
	Description string
	AvatarURL   string
	IsFavorite  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
