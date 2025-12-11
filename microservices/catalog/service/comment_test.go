package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"spotify/internal/mocks"
	"spotify/internal/model"
	"spotify/microservices/catalog/dto"
	repository_mock "spotify/microservices/catalog/mocks/repository"
	pbAuth "spotify/proto/auth"
)

func TestService_PostComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockAuth := mocks.NewMockAuthServiceClient(ctrl)
	svc := New(mockRepo, mockAuth)

	userID := uuid.New()
	trackID := uuid.New()
	req := dto.PostCommentRequest{
		TrackID: trackID.String(),
		Text:    "Great track",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().CreateComment(gomock.Any(), gomock.Any()).Return(nil)
		mockAuth.EXPECT().GetUsers(gomock.Any(), &pbAuth.GetUsersRequest{
			UserIds: []string{userID.String()},
		}).Return(&pbAuth.GetUsersResponse{
			Users: []*pbAuth.UserInfo{{UserId: userID.String(), Login: "test"}},
		}, nil)

		res, err := svc.PostComment(context.Background(), userID, req)
		assert.NoError(t, err)
		assert.Equal(t, "Great track", res.Text)
		assert.Equal(t, "test", res.UserLogin)
	})

	t.Run("invalid track uuid", func(t *testing.T) {
		invalidReq := dto.PostCommentRequest{TrackID: "bad", Text: "txt"}
		_, err := svc.PostComment(context.Background(), userID, invalidReq)
		assert.Error(t, err)
	})

	t.Run("repo error", func(t *testing.T) {
		mockRepo.EXPECT().CreateComment(gomock.Any(), gomock.Any()).Return(errors.New("db fail"))
		_, err := svc.PostComment(context.Background(), userID, req)
		assert.Error(t, err)
	})

	t.Run("auth error", func(t *testing.T) {
		mockRepo.EXPECT().CreateComment(gomock.Any(), gomock.Any()).Return(nil)
		mockAuth.EXPECT().GetUsers(gomock.Any(), gomock.Any()).Return(nil, errors.New("grpc fail"))
		_, err := svc.PostComment(context.Background(), userID, req)
		assert.Error(t, err)
	})
}

func TestService_GetCommentsByTrackID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockAuth := mocks.NewMockAuthServiceClient(ctrl)
	svc := New(mockRepo, mockAuth)

	trackID := uuid.New()
	userID := uuid.New()

	commentModel := model.Comment{
		ID:        uuid.New(),
		TrackID:   trackID,
		UserID:    userID,
		Text:      "Text",
		CreatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetCommentsByTrackID(gomock.Any(), trackID, uint64(10), uint64(0)).
			Return([]model.Comment{commentModel}, nil)

		mockAuth.EXPECT().GetUsers(gomock.Any(), &pbAuth.GetUsersRequest{
			UserIds: []string{userID.String()},
		}).Return(&pbAuth.GetUsersResponse{
			Users: []*pbAuth.UserInfo{{UserId: userID.String(), Login: "User1"}},
		}, nil)

		res, err := svc.GetCommentsByTrackID(context.Background(), trackID, 10, 0)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
		assert.Equal(t, "User1", res[0].UserLogin)
	})

	t.Run("empty", func(t *testing.T) {
		mockRepo.EXPECT().GetCommentsByTrackID(gomock.Any(), trackID, uint64(10), uint64(0)).
			Return([]model.Comment{}, nil)

		res, err := svc.GetCommentsByTrackID(context.Background(), trackID, 10, 0)
		assert.NoError(t, err)
		assert.Empty(t, res)
	})

	t.Run("repo error", func(t *testing.T) {
		mockRepo.EXPECT().GetCommentsByTrackID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, errors.New("db error"))

		_, err := svc.GetCommentsByTrackID(context.Background(), trackID, 10, 0)
		assert.Error(t, err)
	})

	t.Run("auth error", func(t *testing.T) {
		mockRepo.EXPECT().GetCommentsByTrackID(gomock.Any(), trackID, uint64(10), uint64(0)).
			Return([]model.Comment{commentModel}, nil)

		mockAuth.EXPECT().GetUsers(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("grpc fail"))

		_, err := svc.GetCommentsByTrackID(context.Background(), trackID, 10, 0)
		assert.Error(t, err)
	})
}
