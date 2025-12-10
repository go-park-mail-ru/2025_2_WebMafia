package dto

import "time"

//go:generate easyjson $GOFILE

//easyjson:json
type Comment struct {
	ID         string    `json:"id"`
	Text       string    `json:"text"`
	TrackID    string    `json:"track_id"`
	UserID     string    `json:"user_id"`
	UserLogin  string    `json:"user_login"`
	UserAvatar string    `json:"user_avatar"`
	CreatedAt  time.Time `json:"created_at"`
}

//easyjson:json
type PostCommentRequest struct {
	TrackID string `json:"track_id"`
	Text    string `json:"text"`
}
