package dto

type Track struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	DurationS int      `json:"duration_s"`
	FileURL   string   `json:"file_url"`
	PlayCount int64    `json:"play_count"`
	Artists   []Artist `json:"artists"`
	Album     Album    `json:"album"`
	Genres    []Genre  `json:"genres"`
}

type TrackSearch struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	DurationS int      `json:"duration_s"`
	FileURL   string   `json:"file_url"`
	PlayCount int64    `json:"play_count"`
	Artists   []Artist `json:"artists"`
	Album     Album    `json:"album"`
	Genres    []Genre  `json:"genres"`
	Rank      float32  `json:"rank"`
}

type Genre struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Album struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	AvatarURL   string   `json:"avatar_url,omitempty"`
	ReleaseDate string   `json:"release_date"`
	Artists     []Artist `json:"artists"`
}

type Artist struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url,omitempty"`
}
