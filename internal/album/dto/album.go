package dto

type Album struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	AvatarURL   string   `json:"avatar_url,omitempty"`
	Description string   `json:"description,omitempty"`
	ReleaseDate string   `json:"release_date"`
	Artists     []Artist `json:"artists"`
}

type Artist struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url,omitempty"`
	HeaderURL string `json:"header_url,omitempty"`
}
