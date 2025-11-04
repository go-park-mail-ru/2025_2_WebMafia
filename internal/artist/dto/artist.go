package dto

type Artist struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	AvatarURL   string `json:"avatar_url,omitempty"`
	HeaderURL   string `json:"header_url,omitempty"`
	Description string `json:"description,omitempty"`
	PlayCount   int64  `json:"play_count"`
}
