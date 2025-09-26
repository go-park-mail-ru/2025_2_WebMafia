package model

import "time"

type Artist struct {
	ArtistID  uint64    `json:"artist_id"`
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
