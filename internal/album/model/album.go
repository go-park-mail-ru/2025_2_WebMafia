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

type Artist struct {
	ID          uuid.UUID
	Name        string
	AvatarURL   string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
