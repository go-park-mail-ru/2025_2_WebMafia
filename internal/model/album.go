package model

import (
	"time"

	"github.com/google/uuid"
)

type Album struct {
	ID          uuid.UUID
	Title       string
	AvatarURL   string
	ArtistID    uuid.UUID
	Description string
	ReleaseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

