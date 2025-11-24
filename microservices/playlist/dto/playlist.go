package dto

import (
	"github.com/google/uuid"
	"io"
)

type Playlist struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description,omitempty"`
	IsFavorite  bool    `json:"is_favorite"`
	AvatarURL   string  `json:"avatar_url,omitempty"`
	Tracks      []Track `json:"tracks,omitempty"`
}

type Track struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	DurationS int      `json:"duration_s"`
	FileURL   string   `json:"file_url"`
	Artists   []Artist `json:"artists"`
	Album     Album    `json:"album"`
}

type Album struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

type Artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreatePlaylistRequest struct {
	UserID      uuid.UUID `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

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

type UploadPlaylistAvatarResponse struct {
	URL string `json:"avatar_url"`
}

type DeletePlaylistAvatarRequest struct {
	PlaylistID uuid.UUID `json:"-"`
}

type AddTrackToPlaylistRequest struct {
	PlaylistID uuid.UUID `json:"-"`
	TrackID    string    `json:"track_id"`
}

type RemoveTrackFromPlaylistRequest struct {
	PlaylistID uuid.UUID `json:"-"`
	TrackID    string    `json:"track_id"`
}
