package dto

//go:generate easyjson $GOFILE

import (
	"io"

	"github.com/google/uuid"
)

//easyjson:json
type Playlist struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description,omitempty"`
	IsFavorite  bool    `json:"is_favorite"`
	AvatarURL   string  `json:"avatar_url,omitempty"`
	Tracks      []Track `json:"tracks,omitempty"`
}

//easyjson:json
type Track struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	DurationS int      `json:"duration_s"`
	FileURL   string   `json:"file_url"`
	Artists   []Artist `json:"artists"`
	Album     Album    `json:"album"`
}

//easyjson:json
type Album struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

//easyjson:json
type Artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//easyjson:json
type CreatePlaylistRequest struct {
	UserID      uuid.UUID `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

//easyjson:json
type UpdatePlaylistRequest struct {
	ID          uuid.UUID `json:"-"`
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	IsFavorite  *bool     `json:"is_favorite"`
}

type DeletePlaylistRequest struct {
	ID uuid.UUID
}

type GetPlaylistRequest struct {
	ID uuid.UUID
}

type GetPlaylistsByUserRequest struct {
	UserID uuid.UUID
	Limit  uint64
	Offset uint64
}

type GetFavoritePlaylistRequest struct {
	UserID uuid.UUID
}

//easyjson:json
type AddTrackToFavoriteRequest struct {
	UserID  uuid.UUID `json:"-"`
	TrackID string    `json:"track_id"`
}

type UploadPlaylistAvatarRequest struct {
	PlaylistID  uuid.UUID `json:"-"`
	File        io.Reader `json:"-"`
	Size        int64     `json:"-"`
	ContentType string    `json:"-"`
}

//easyjson:json
type UploadPlaylistAvatarResponse struct {
	URL string `json:"avatar_url"`
}

type DeletePlaylistAvatarRequest struct {
	PlaylistID uuid.UUID `json:"-"`
}

//easyjson:json
type AddTrackToPlaylistRequest struct {
	PlaylistID uuid.UUID `json:"-"`
	TrackID    string    `json:"track_id"`
}

//easyjson:json
type RemoveTrackFromPlaylistRequest struct {
	PlaylistID uuid.UUID `json:"-"`
	TrackID    string    `json:"track_id"`
}
