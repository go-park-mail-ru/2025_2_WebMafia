package model

import (
	"time"

	"github.com/google/uuid"
)

type FavoriteAlbum struct {
	UserID    uuid.UUID
	AlbumID   uuid.UUID
	CreatedAt time.Time
}
