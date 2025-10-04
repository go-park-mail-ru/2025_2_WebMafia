package handler

import (
	"spotify/internal/store"
	"spotify/pkg/jwtmanager"
	"testing"
	"time"
)

func initTestEnv(t *testing.T) (*Handlers, *store.MemoryStore, *jwtmanager.Manager) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)
	return handlers, dataStore, jwtManager
}
