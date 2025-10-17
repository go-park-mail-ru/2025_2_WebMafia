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
