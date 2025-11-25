package grpc

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"spotify/pkg/csrfmanager"
	"spotify/pkg/jwtmanager"
	pb "spotify/proto/auth"
)

func TestHandler_ValidateToken(t *testing.T) {
	jwtMgr := jwtmanager.NewManager("secret", time.Hour)
	csrfMgr := csrfmanager.NewManager("secret", time.Hour)
	handler := NewHandler(jwtMgr, csrfMgr)

	userID := "user123"
	token, err := jwtMgr.Generate(userID)
	require.NoError(t, err)

	t.Run("valid token", func(t *testing.T) {
		resp, err := handler.ValidateToken(context.Background(), &pb.ValidateTokenRequest{Token: token})
		assert.NoError(t, err)
		assert.True(t, resp.IsValid)
		assert.Equal(t, userID, resp.UserId)
		assert.NotEmpty(t, resp.SessionId)
	})

	t.Run("invalid token", func(t *testing.T) {
		resp, err := handler.ValidateToken(context.Background(), &pb.ValidateTokenRequest{Token: "bad.token"})
		assert.NoError(t, err)
		assert.False(t, resp.IsValid)
	})
}

func TestHandler_CheckCSRF(t *testing.T) {
	jwtMgr := jwtmanager.NewManager("secret", time.Hour)
	csrfMgr := csrfmanager.NewManager("secret", time.Hour)
	handler := NewHandler(jwtMgr, csrfMgr)

	userID := "user123"
	sessionID := "sess1"

	token, err := csrfMgr.Generate(userID, sessionID)
	require.NoError(t, err)

	t.Run("valid csrf", func(t *testing.T) {
		resp, err := handler.CheckCSRF(context.Background(), &pb.CheckCSRFRequest{
			UserId:    userID,
			SessionId: sessionID,
			CsrfToken: token,
		})
		assert.NoError(t, err)
		assert.True(t, resp.IsValid)
	})

	t.Run("invalid csrf", func(t *testing.T) {
		resp, err := handler.CheckCSRF(context.Background(), &pb.CheckCSRFRequest{
			UserId:    userID,
			SessionId: sessionID,
			CsrfToken: "bad_token",
		})
		assert.NoError(t, err)
		assert.False(t, resp.IsValid)
	})
}
