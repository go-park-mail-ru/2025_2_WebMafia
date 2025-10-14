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
