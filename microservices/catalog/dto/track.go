package dto

//go:generate easyjson $GOFILE

//easyjson:json
type Track struct {
	ID        string           `json:"id"`
	Title     string           `json:"title"`
	DurationS int              `json:"duration_s"`
	FileURL   string           `json:"file_url"`
	PlayCount int64            `json:"play_count"`
	Artists   []ArtistForTrack `json:"artists"`
	Album     AlbumForTrack    `json:"album"`
	Genres    []Genre          `json:"genres"`
}

//easyjson:json
type ArtistForTrack struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

//easyjson:json
type AlbumForTrack struct {
	ID          string           `json:"id"`
	Title       string           `json:"title"`
	AvatarURL   string           `json:"avatar_url,omitempty"`
	ReleaseDate string           `json:"release_date"`
	Artists     []ArtistForTrack `json:"artists"`
}

//easyjson:json
type Genre struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//easyjson:json
type TrackSearch struct {
	Track
	Rank float32 `json:"rank"`
}
