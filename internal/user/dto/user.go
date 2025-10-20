package dto

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
	File        []byte
	ContentType string
}
type UploadAvatarResponse struct {
	URL string `json:"avatar_url"`
}

type DeleteAvatarRequest struct {
	UserID string
}
