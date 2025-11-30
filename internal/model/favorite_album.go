package model

import (
	"github.com/google/uuid"
	"time"
)

type FavoriteAlbum struct {
	UserID    uuid.UUID
	AlbumID   uuid.UUID
	CreatedAt time.Time
}
