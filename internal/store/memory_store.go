package store

import (
	"context"
	"spotify/internal/model"
	"sync"
	"time"

	"github.com/google/uuid"
)

type MemoryStore struct {
	mu *sync.RWMutex

	users map[uuid.UUID]*model.User
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		mu:    &sync.RWMutex{},
		users: make(map[uuid.UUID]*model.User),
	}
}

func (s *MemoryStore) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, u := range s.users {
		if u.Login == user.Login || u.Email == user.Email {
			return nil, ErrUserAlreadyExists
		}
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	s.users[user.ID] = &user

	return &user, nil
}

func (s *MemoryStore) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, u := range s.users {
		if u.Login == login {
			return u, nil
		}
	}
	return nil, ErrUserNotFound
}

func (s *MemoryStore) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, u := range s.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, ErrUserNotFound
}

func (s *MemoryStore) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, ok := s.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}
