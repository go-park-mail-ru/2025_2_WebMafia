package dto

//go:generate easyjson $GOFILE

import "io"

//easyjson:json
type RegisterRequest struct {
	Login    string
	Email    string
	Password string
}

//easyjson:json
type RegisterResponse struct {
	ID    string
	Login string
	Email string
}

//easyjson:json
type LoginRequest struct {
	Login    string
	Password string
}

//easyjson:json
type LoginResponse struct {
	ID string
}

type UploadAvatarRequest struct {
	UserID      string
	File        io.Reader
	Size        int64
	ContentType string
}

//easyjson:json
type UploadAvatarResponse struct {
	URL string `json:"avatar_url"`
}

type DeleteAvatarRequest struct {
	UserID string
}

//easyjson:json
type UpdateProfileRequest struct {
	UserID   string
	Login    string
	Email    string
	Password string
}

//easyjson:json
type UpdateProfileResponse struct {
	ID    string
	Login string
	Email string
}

type GetProfileRequest struct {
	UserID string
}

//easyjson:json
type GetProfileResponse struct {
	ID        string
	Login     string
	Email     string
	AvatarURL string
}
