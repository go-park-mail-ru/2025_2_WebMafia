package grpc

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"spotify/microservices/auth/dto"
	authmocks "spotify/microservices/auth/mocks/grpc_service"
	"spotify/pkg/csrfmanager"
	"spotify/pkg/jwtmanager"
	pb "spotify/proto/auth"
)

func TestHandler_ValidateToken(t *testing.T) {
	jwtMgr := jwtmanager.NewManager("secret", time.Hour)
	csrfMgr := csrfmanager.NewManager("secret", time.Hour)

	handler := NewHandler(jwtMgr, csrfMgr, nil)

	userID := "user123"
	token, err := jwtMgr.Generate(userID)
	require.NoError(t, err)

	t.Run("valid token", func(t *testing.T) {
		resp, err := handler.ValidateToken(context.Background(),
			&pb.ValidateTokenRequest{Token: token},
		)
		assert.NoError(t, err)
		assert.True(t, resp.IsValid)
		assert.Equal(t, userID, resp.UserId)
		assert.NotEmpty(t, resp.SessionId)
	})

	t.Run("invalid token", func(t *testing.T) {
		resp, err := handler.ValidateToken(context.Background(),
			&pb.ValidateTokenRequest{Token: "bad.token"},
		)
		assert.NoError(t, err)
		assert.False(t, resp.IsValid)
	})
}

func TestHandler_CheckCSRF(t *testing.T) {
	jwtMgr := jwtmanager.NewManager("secret", time.Hour)
	csrfMgr := csrfmanager.NewManager("secret", time.Hour)

	handler := NewHandler(jwtMgr, csrfMgr, nil)

	userID := "user123"
	sessionID := "sess1"

	token, err := csrfMgr.Generate(userID, sessionID)
	require.NoError(t, err)

	t.Run("valid csrf", func(t *testing.T) {
		resp, err := handler.CheckCSRF(context.Background(),
			&pb.CheckCSRFRequest{
				UserId:    userID,
				SessionId: sessionID,
				CsrfToken: token,
			},
		)
		assert.NoError(t, err)
		assert.True(t, resp.IsValid)
	})

	t.Run("invalid csrf", func(t *testing.T) {
		resp, err := handler.CheckCSRF(context.Background(),
			&pb.CheckCSRFRequest{
				UserId:    userID,
				SessionId: sessionID,
				CsrfToken: "bad_token",
			},
		)
		assert.NoError(t, err)
		assert.False(t, resp.IsValid)
	})
}

func TestHandler_GetUsers(t *testing.T) {
	jwtMgr := jwtmanager.NewManager("secret", time.Hour)
	csrfMgr := csrfmanager.NewManager("secret", time.Hour)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := authmocks.NewMockIUserService(ctrl)
	handler := NewHandler(jwtMgr, csrfMgr, mockSvc)

	t.Run("empty request", func(t *testing.T) {
		resp, err := handler.GetUsers(context.Background(),
			&pb.GetUsersRequest{UserIds: []string{}},
		)
		require.NoError(t, err)
		assert.Empty(t, resp.Users)
	})

	t.Run("service returns users", func(t *testing.T) {
		input := []string{"u1", "u2"}

		mockSvc.EXPECT().
			GetUsersByIDs(gomock.Any(), input).
			Return([]dto.GetProfileResponse{
				{ID: "u1", Login: "alice", AvatarURL: "a.png"},
				{ID: "u2", Login: "bob", AvatarURL: "b.png"},
			}, nil)

		resp, err := handler.GetUsers(context.Background(),
			&pb.GetUsersRequest{UserIds: input},
		)

		require.NoError(t, err)
		require.Len(t, resp.Users, 2)

		assert.Equal(t, "u1", resp.Users[0].UserId)
		assert.Equal(t, "alice", resp.Users[0].Login)

		assert.Equal(t, "u2", resp.Users[1].UserId)
		assert.Equal(t, "bob", resp.Users[1].Login)
	})

	t.Run("service error", func(t *testing.T) {
		input := []string{"broken"}

		mockSvc.EXPECT().
			GetUsersByIDs(gomock.Any(), input).
			Return(nil, errors.New("db error"))

		resp, err := handler.GetUsers(context.Background(),
			&pb.GetUsersRequest{UserIds: input},
		)

		require.Nil(t, resp)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to fetch users")
	})
}
