package dto

import "io"

type RegisterRequest struct {
	Login    string
	Email    string
	Password string
}
type RegisterResponse struct {
	ID    string
	Login string
	Email string
}

type LoginRequest struct {
	Login    string
	Password string
}

type LoginResponse struct {
	ID string
}

type UploadAvatarRequest struct {
	UserID      string
	File        io.Reader
	Size        int64
	ContentType string
}
type UploadAvatarResponse struct {
	URL string `json:"avatar_url"`
}

type DeleteAvatarRequest struct {
	UserID string
}

type UpdateProfileRequest struct {
	UserID   string
	Login    string
	Email    string
	Password string
}

type UpdateProfileResponse struct {
	ID    string
	Login string
	Email string
}

type GetProfileRequest struct {
	UserID string
}
type GetProfileResponse struct {
	ID        string
	Login     string
	Email     string
	AvatarURL string
}

type UpdateRoleRequest struct {
	UserID string
	Role   string
}

type UpdateRoleResponse struct {
	ID   string
	Role string
}
