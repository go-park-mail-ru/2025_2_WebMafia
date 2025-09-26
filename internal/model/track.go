package model

import "time"

type Genre struct {
	GenreID   uint64    `json:"genre_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Track struct {
	TrackID    uint64    `json:"track_id"`
	Title      string    `json:"title"`
	DurationMs int       `json:"duration_ms"`
	FileURL    string    `json:"file_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Artists    []*Artist `json:"artists,omitempty"`
	Genres     []*Genre  `json:"genres,omitempty"`
	Album      *Album    `json:"album,omitempty"`
}
