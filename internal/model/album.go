package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Album struct {
	ID          uuid.UUID
	Title       string
	AvatarURL   string
	ArtistID    uuid.UUID
	Description sql.NullString
	ReleaseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
