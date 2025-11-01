package service

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	mock_user "spotify/internal/mocks/user"
	"spotify/internal/model"
	"spotify/internal/user/dto"
	"spotify/internal/user/repository/postgres"
	"spotify/internal/user/tools"
)

func TestUserService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_user.NewMockIRepository(ctrl)
	mockStorage := mock_user.NewMockIStorage(ctrl)

	userService := NewUserService(mockRepo, mockStorage)

	loginRequest := dto.LoginRequest{
		Login:    "testuser",
		Password: "password123",
	}
	hashedPassword, err := tools.Hash(loginRequest.Password)
	require.NoError(t, err)

	userFromDB := &model.User{
		ID:           uuid.New(),
		Login:        loginRequest.Login,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	t.Run("success - valid credentials", func(t *testing.T) {
		mockRepo.EXPECT().
			GetUserByLogin(gomock.Any(), loginRequest.Login).
			Return(userFromDB, nil).
			Times(1)

		response, err := userService.Login(context.Background(), loginRequest)

		assert.NoError(t, err)
		require.NotNil(t, response)
		assert.Equal(t, userFromDB.ID.String(), response.ID)
	})

	t.Run("fail - user not found", func(t *testing.T) {
		mockRepo.EXPECT().
			GetUserByLogin(gomock.Any(), loginRequest.Login).
			Return(nil, postgres.ErrNotFound).
			Times(1)

		response, err := userService.Login(context.Background(), loginRequest)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrNotFound)
		assert.Nil(t, response)
	})

	t.Run("fail - invalid password", func(t *testing.T) {
		userWithWrongPass := &model.User{
			ID:           uuid.New(),
			Login:        loginRequest.Login,
			PasswordHash: "wrong_hash",
		}

		mockRepo.EXPECT().
			GetUserByLogin(gomock.Any(), loginRequest.Login).
			Return(userWithWrongPass, nil).
			Times(1)

		response, err := userService.Login(context.Background(), loginRequest)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrValidation)
		assert.Nil(t, response)
	})
}

func TestUserService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_user.NewMockIRepository(ctrl)
	mockStorage := mock_user.NewMockIStorage(ctrl)
	userService := NewUserService(mockRepo, mockStorage)

	registerRequest := dto.RegisterRequest{
		Login:    "newuser",
		Email:    "new@example.com",
		Password: "password123",
	}

	t.Run("success - user registered", func(t *testing.T) {
		mockRepo.EXPECT().
			CreateUser(gomock.Any(), gomock.Any()).
			Do(func(ctx context.Context, user model.User) {
				assert.Equal(t, registerRequest.Login, user.Login)
				assert.Equal(t, registerRequest.Email, user.Email)
				assert.NotEmpty(t, user.PasswordHash)
				assert.NotEqual(t, registerRequest.Password, user.PasswordHash)
			}).
			Return(nil).
			Times(1)

		response, err := userService.Register(context.Background(), registerRequest)

		assert.NoError(t, err)
		require.NotNil(t, response)
		assert.Equal(t, registerRequest.Login, response.Login)
		assert.Equal(t, registerRequest.Email, response.Email)
		assert.NotEmpty(t, response.ID)
	})

	t.Run("fail - user already exists", func(t *testing.T) {
		mockRepo.EXPECT().
			CreateUser(gomock.Any(), gomock.Any()).
			Return(postgres.ErrConflict).
			Times(1)

		response, err := userService.Register(context.Background(), registerRequest)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrConflict)
		assert.Nil(t, response)
	})
}

func TestUserService_UploadAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_user.NewMockIRepository(ctrl)
	mockStorage := mock_user.NewMockIStorage(ctrl)
	userService := NewUserService(mockRepo, mockStorage)

	req := dto.UploadAvatarRequest{
		UserID:      "user-id",
		Size:        1024,
		ContentType: "image/jpeg",
	}
	objectName := "avatar.jpg"
	url := "http://example.com/avatar.jpg"

	t.Run("success", func(t *testing.T) {
		mockStorage.EXPECT().UploadAvatar(gomock.Any(), gomock.Any(), req.Size, req.ContentType).Return(objectName, nil)
		mockRepo.EXPECT().UpdateUserAvatar(gomock.Any(), req.UserID, objectName).Return(nil)
		mockStorage.EXPECT().GetAvatarURL(gomock.Any(), objectName).Return(url, nil)

		resp, err := userService.UploadAvatar(context.Background(), req)

		assert.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, url, resp.URL)
	})
}

func TestUserService_DeleteAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_user.NewMockIRepository(ctrl)
	mockStorage := mock_user.NewMockIStorage(ctrl)
	userService := NewUserService(mockRepo, mockStorage)

	req := dto.DeleteAvatarRequest{UserID: "user-id"}
	userWithAvatar := &model.User{ID: uuid.New(), AvatarURL: "avatar.jpg"}
	userWithoutAvatar := &model.User{ID: uuid.New(), AvatarURL: ""}

	t.Run("success with existing avatar", func(t *testing.T) {
		mockRepo.EXPECT().GetUserByID(gomock.Any(), req.UserID).Return(userWithAvatar, nil)
		mockStorage.EXPECT().DeleteAvatar(gomock.Any(), userWithAvatar.AvatarURL).Return(nil)
		mockRepo.EXPECT().UpdateUserAvatar(gomock.Any(), req.UserID, "").Return(nil)

		err := userService.DeleteAvatar(context.Background(), req)
		assert.NoError(t, err)
	})

	t.Run("success user without avatar", func(t *testing.T) {
		mockRepo.EXPECT().GetUserByID(gomock.Any(), req.UserID).Return(userWithoutAvatar, nil)

		err := userService.DeleteAvatar(context.Background(), req)
		assert.NoError(t, err)
	})
}
