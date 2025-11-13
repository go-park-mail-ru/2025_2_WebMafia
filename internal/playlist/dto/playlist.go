package dto

type Playlist struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	IsFavorite  bool     `json:"is_favorite"`
	AvatarURL   string   `json:"avatar_url,omitempty"`
	Tracks      []string `json:"tracks,omitempty"`
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
	UserID      string `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

type UpdatePlaylistRequest struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	IsFavorite  bool   `json:"is_favorite"`
}

type DeletePlaylistRequest struct {
	ID string `json:"id"`
}

type GetPlaylistRequest struct {
	ID string `json:"id"`
}

type GetPlaylistsByUserRequest struct {
	UserID string `json:"user_id"`
	Limit  uint64 `json:"limit,omitempty"`
	Offset uint64 `json:"offset,omitempty"`
}

type GetFavoritePlaylistRequest struct {
	UserID string `json:"-"`
}
type AddTrackToFavoriteRequest struct {
	UserID  string `json:"-"`
	TrackID string `json:"track_id"`
}
