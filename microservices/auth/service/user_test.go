package service

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	repository_mock "spotify/microservices/auth/mocks/repository"
	storage_mock "spotify/microservices/auth/mocks/storage"

	"spotify/internal/model"
	"spotify/microservices/auth/dto"
	"spotify/microservices/auth/repository/postgres"
	"spotify/microservices/auth/tools"
)

func TestUserService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockStorage := storage_mock.NewMockIStorage(ctrl)

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
		assert.True(t, errors.Is(err, ErrNotFound) || strings.Contains(err.Error(), "not found"))
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

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockStorage := storage_mock.NewMockIStorage(ctrl)
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
		assert.True(t, errors.Is(err, ErrConflict) || strings.Contains(err.Error(), "user already exists"))
		assert.Nil(t, response)
	})
}

func TestUserService_UploadAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockStorage := storage_mock.NewMockIStorage(ctrl)
	userService := NewUserService(mockRepo, mockStorage)

	req := dto.UploadAvatarRequest{
		UserID:      "user-id",
		Size:        1024,
		ContentType: "image/jpeg",
	}
	objectName := "avatar.jpg"

	t.Run("success", func(t *testing.T) {
		mockStorage.EXPECT().UploadAvatar(gomock.Any(), gomock.Any(), req.Size, req.ContentType).Return(objectName, nil)
		mockRepo.EXPECT().UpdateUserAvatar(gomock.Any(), req.UserID, objectName).Return(nil)

		resp, err := userService.UploadAvatar(context.Background(), req)

		assert.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, objectName, resp.URL)
	})
}

func TestUserService_DeleteAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockStorage := storage_mock.NewMockIStorage(ctrl)
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

func TestUserService_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockStorage := storage_mock.NewMockIStorage(ctrl)
	userService := NewUserService(mockRepo, mockStorage)

	userID := uuid.New()
	updateRequest := dto.UpdateProfileRequest{
		UserID:   userID.String(),
		Login:    "newlogin",
		Email:    "new@email.com",
		Password: "newpassword123",
	}

	t.Run("success - with password change", func(t *testing.T) {
		mockRepo.EXPECT().
			UpdateUserProfile(gomock.Any(), gomock.Any()).
			Do(func(ctx context.Context, user model.User) {
				assert.Equal(t, updateRequest.Login, user.Login)
				assert.Equal(t, updateRequest.Email, user.Email)
				assert.NotEmpty(t, user.PasswordHash)
				assert.NoError(t, tools.Compare(user.PasswordHash, updateRequest.Password))
			}).
			Return(nil)

		response, err := userService.UpdateProfile(context.Background(), updateRequest)

		assert.NoError(t, err)
		require.NotNil(t, response)
		assert.Equal(t, updateRequest.Login, response.Login)
		assert.Equal(t, updateRequest.Email, response.Email)
	})

	t.Run("success - without password change", func(t *testing.T) {
		requestWithoutPass := updateRequest
		requestWithoutPass.Password = ""

		mockRepo.EXPECT().
			UpdateUserProfile(gomock.Any(), gomock.Any()).
			Do(func(ctx context.Context, user model.User) {
				assert.Empty(t, user.PasswordHash)
			}).
			Return(nil)

		_, err := userService.UpdateProfile(context.Background(), requestWithoutPass)
		assert.NoError(t, err)
	})

	t.Run("fail - repository error", func(t *testing.T) {
		expectedErr := errors.New("db error")
		mockRepo.EXPECT().
			UpdateUserProfile(gomock.Any(), gomock.Any()).
			Return(expectedErr)

		_, err := userService.UpdateProfile(context.Background(), updateRequest)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), expectedErr.Error())
	})
}
