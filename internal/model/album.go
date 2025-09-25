package model

import "time"

type Album struct {
	AlbumID     string    `json:"album_id"`
	Title       string    `json:"title"`
	AvatarURL   string    `json:"avatar_url"`
	ArtistID    string    `json:"artist_id"`
	Artist      *Artist   `json:"artist,omitempty"`
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Tracks      []*Track  `json:"tracks,omitempty"`
}
