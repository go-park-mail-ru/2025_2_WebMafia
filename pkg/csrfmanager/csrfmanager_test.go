package csrfmanager

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCSRFManager(t *testing.T) {
	secret := "my-super-secret-key-for-csrf"
	ttl := time.Minute * 15
	manager := NewManager(secret, ttl)

	userID := uuid.New().String()
	sessionID := uuid.New().String()

	t.Run("generate and check success", func(t *testing.T) {
		token, err := manager.Generate(userID, sessionID)
		require.NoError(t, err)
		assert.NotEmpty(t, token)

		valid, err := manager.Check(userID, sessionID, token)
		assert.NoError(t, err)
		assert.True(t, valid)
	})

	t.Run("check fails on wrong user/session", func(t *testing.T) {
		token, err := manager.Generate(userID, sessionID)
		require.NoError(t, err)

		valid, err := manager.Check("wrong-user-id", sessionID, token)
		assert.Error(t, err)
		assert.False(t, valid)

		valid, err = manager.Check(userID, "wrong-session-id", token)
		assert.Error(t, err)
		assert.False(t, valid)
	})

	t.Run("check fails on expired token", func(t *testing.T) {
		shortTTLManager := NewManager(secret, -time.Minute)
		token, err := shortTTLManager.Generate(userID, sessionID)
		require.NoError(t, err)

		valid, err := manager.Check(userID, sessionID, token)
		assert.Error(t, err)
		assert.False(t, valid)
	})

	t.Run("check fails on invalid token format", func(t *testing.T) {
		_, err := manager.Check(userID, sessionID, "invalid-token")
		assert.Error(t, err)
	})
}
