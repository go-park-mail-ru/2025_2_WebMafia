package dto

type Album struct {
	ID          string           `json:"id"`
	Title       string           `json:"title"`
	Type        string           `json:"type"`
	AvatarURL   string           `json:"avatar_url,omitempty"`
	Description string           `json:"description,omitempty"`
	ReleaseDate string           `json:"release_date"`
	Artists     []ArtistForAlbum `json:"artists"`
}

type ArtistForAlbum struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url,omitempty"`
	HeaderURL string `json:"header_url,omitempty"`
}

type AlbumSearch struct {
	Album
	Rank float32 `json:"rank"`
}
