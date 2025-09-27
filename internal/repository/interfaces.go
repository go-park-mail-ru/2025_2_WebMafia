package repository

import (
	"context"
	"spotify/internal/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) (*model.User, error)
	GetUserByLogin(ctx context.Context, login string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error)
}
