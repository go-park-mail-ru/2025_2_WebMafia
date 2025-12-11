package service

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"spotify/internal/mocks"
	"spotify/internal/model"
	repository_mock "spotify/microservices/catalog/mocks/repository"
)

func TestService_GetAlbumByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockAuth := mocks.NewMockAuthServiceClient(ctrl)
	svc := New(mockRepo, mockAuth)

	albumID := uuid.New()
	artistID := uuid.New()

	albumModel := &model.Album{
		ID:          albumID,
		Title:       "Test Album",
		ArtistID:    artistID,
		ReleaseDate: time.Now(),
		Description: sql.NullString{String: "Desc", Valid: true},
	}

	artistModel := &model.Artist{
		ID:   artistID,
		Name: "Test Artist",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetAlbumByID(gomock.Any(), albumID).Return(albumModel, nil)

		mockRepo.EXPECT().GetArtistByID(gomock.Any(), artistID).Return(artistModel, nil)
		mockRepo.EXPECT().GetTotalPlaysByArtistID(gomock.Any(), artistID).Return(int64(10), nil)

		res, err := svc.GetAlbumByID(context.Background(), albumID)

		assert.NoError(t, err)
		assert.Equal(t, "Test Album", res.Title)
		assert.Equal(t, "Test Artist", res.Artists[0].Name)
	})
}
