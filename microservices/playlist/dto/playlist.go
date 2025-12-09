package dto

//go:generate easyjson $GOFILE

import (
	"io"
	"time"

	"github.com/google/uuid"
)

//easyjson:json
type Playlist struct {
	ID          string    `json:"id"`
	CreatorID   string    `json:"creator_id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	IsFavorite  bool      `json:"is_favorite"`
	AvatarURL   string    `json:"avatar_url,omitempty"`
	Tracks      []Track   `json:"tracks,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
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

// любимые артисты

//easyjson:json
type FavoriteArtist struct {
	ID        string    `json:"id"`
	CreatorID string    `json:"creator_id"`
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatar_url,omitempty"`
	PlayCount int64     `json:"play_count"`
	CreatedAt time.Time `json:"created_at"`
}
type AddArtistToFavoriteRequest struct {
	UserID   uuid.UUID `json:"-"`
	ArtistID string    `json:"artist_id"`
}
type RemoveArtistFromFavoriteRequest struct {
	UserID   uuid.UUID `json:"-"`
	ArtistID string    `json:"artist_id"`
}

// любимые альбомы

//easyjson:json
type FavoriteAlbum struct {
	ID        string           `json:"id"`
	CreatorID string           `json:"creator_id"`
	Title     string           `json:"title"`
	AvatarURL string           `json:"avatar_url,omitempty"`
	Artists   []ArtistForAlbum `json:"artists"`
	Type      string           `json:"type"`
	CreatedAt time.Time        `json:"created_at"`
}

//easyjson:json
type ArtistForAlbum struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AddAlbumToFavoriteRequest struct {
	UserID  uuid.UUID `json:"-"`
	AlbumID string    `json:"album_id"`
}

type RemoveAlbumFromFavoriteRequest struct {
	UserID  uuid.UUID `json:"-"`
	AlbumID string    `json:"album_id"`
}

//easyjson:json
type GeneratedMeta struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
