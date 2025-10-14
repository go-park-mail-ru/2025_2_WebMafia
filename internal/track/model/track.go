package model

import (
	"time"

	"github.com/google/uuid"
)

type Genre struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
}

type Track struct {
	ID          uuid.UUID
	Title       string
	DurationMs  int
	FileURL     string
	Description string
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
