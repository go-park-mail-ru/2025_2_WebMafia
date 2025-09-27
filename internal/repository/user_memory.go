package repository

import (
	"context"
	"spotify/internal/model"
	"sync"
	"time"

	"github.com/google/uuid"
)

type UserMemoryRepository struct {
	mu    *sync.RWMutex
	users map[uuid.UUID]*model.User
}

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		users: make(map[uuid.UUID]*model.User),
		mu:    &sync.RWMutex{},
	}
}

func (r *UserMemoryRepository) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, u := range r.users {
		if u.Login == user.Login || u.Email == user.Email {
			return &model.User{}, ErrUserAlreadyExists
		}
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	r.users[user.ID] = &user

	return &user, nil
}

func (r *UserMemoryRepository) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.users {
		if u.Login == login {
			return u, nil
		}
	}
	return &model.User{}, ErrUserNotFound
}

func (r *UserMemoryRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return &model.User{}, ErrUserNotFound
}

func (r *UserMemoryRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return &model.User{}, ErrUserNotFound
	}
	return user, nil
}
