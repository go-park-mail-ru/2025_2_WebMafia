package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Artist struct {
	ID          uuid.UUID
	Name        string
	AvatarURL   string
	HeaderURL   string
	Description sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type FavoriteArtist struct {
	UserID    uuid.UUID
	ArtistID  uuid.UUID
	CreatedAt time.Time
}
