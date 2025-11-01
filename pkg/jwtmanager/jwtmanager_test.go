package jwtmanager

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJWTManager(t *testing.T) {
	secret := "my-super-secret-key-for-testing"
	ttl := time.Minute * 15
	manager := NewManager(secret, ttl)

	userID := uuid.New().String()

	t.Run("generate and validate success", func(t *testing.T) {
		token, err := manager.Generate(userID)
		require.NoError(t, err)
		assert.NotEmpty(t, token)

		claims, err := manager.Validate(token)
		require.NoError(t, err)
		require.NotNil(t, claims)
		assert.Equal(t, userID, claims.UserID)
		assert.WithinDuration(t, time.Now().Add(ttl), time.Unix(claims.Exp, 0), time.Second)
	})

	t.Run("validate expired token", func(t *testing.T) {
		shortTTLManager := NewManager(secret, -time.Minute)
		token, err := shortTTLManager.Generate(userID)
		require.NoError(t, err)

		_, err = manager.Validate(token)
		assert.Error(t, err)
		assert.EqualError(t, err, "token expired")
	})

	t.Run("validate invalid signature", func(t *testing.T) {
		token, err := manager.Generate(userID)
		require.NoError(t, err)

		invalidToken := token + "invalid"
		_, err = manager.Validate(invalidToken)
		assert.Error(t, err)

		wrongSecretManager := NewManager("wrong-secret", ttl)
		_, err = wrongSecretManager.Validate(token)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid signature")
	})
}
