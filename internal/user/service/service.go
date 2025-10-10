package service

import (
	"context"
	"spotify/internal/user/model"
	"spotify/internal/user/tools/password"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByLogin(ctx context.Context, login string) (*model.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, login, email, userPassword string) (model.User, error) {
	if len(login) < 5 {
		return model.User{}, &IsServiceError{Type: ErrValidation, Message: "login too short (min 5 chars)"}
	}
	if len(userPassword) < 8 {
		return model.User{}, &IsServiceError{Type: ErrValidation, Message: "password too short (min 8 chars)"}
	}
	if !containsAt(email) {
		return model.User{}, &IsServiceError{Type: ErrValidation, Message: "invalid email format"}
	}

	if _, err := s.repo.GetUserByEmail(ctx, email); err == nil {
		return model.User{}, &IsServiceError{Type: ErrConflict, Message: "user with this email already exists"}
	}
	if _, err := s.repo.GetUserByLogin(ctx, login); err == nil {
		return model.User{}, &IsServiceError{Type: ErrConflict, Message: "user with this login already exists"}
	}

	hash, err := password.Hash(userPassword)
	if err != nil {
		return model.User{}, &IsServiceError{Type: ErrInternal, Message: "failed to hash password", Err: err}
	}

	user := model.User{
		ID:           uuid.New(),
		Login:        login,
		Email:        email,
		PasswordHash: hash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return model.User{}, &IsServiceError{Type: ErrInternal, Message: "failed to create user", Err: err}
	}

	return user, nil
}
func containsAt(email string) bool {
	for _, c := range email {
		if c == '@' {
			return true
		}
	}
	return false
}
