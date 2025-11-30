package model

import (
	"github.com/google/uuid"
	"time"
)

type FavoriteArtist struct {
	UserID    uuid.UUID
	ArtistID  uuid.UUID
	CreatedAt time.Time
}
