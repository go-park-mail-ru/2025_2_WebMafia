package model

import (
	"time"

	"github.com/google/uuid"
)

type FavoriteArtist struct {
	UserID    uuid.UUID
	ArtistID  uuid.UUID
	CreatedAt time.Time
}
