package model

import (
	"time"

	"github.com/google/uuid"
)

type Playlist struct {
	ID          uuid.UUID
	Title       string
	Description string
	AvatarURL   string
	IsFavorite  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
