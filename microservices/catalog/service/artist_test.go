package service

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"spotify/internal/model"
	repository_mock "spotify/microservices/catalog/mocks/repository"
)

func TestService_GetArtistByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo)

	artistID := uuid.New()
	expectedArtist := &model.Artist{
		ID:        artistID,
		Name:      "Test Artist",
		AvatarURL: "avatar.jpg",
		HeaderURL: "header.jpg",
		Description: sql.NullString{
			String: "Best artist",
			Valid:  true,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetArtistByID(gomock.Any(), artistID).Return(expectedArtist, nil)
		mockRepo.EXPECT().GetTotalPlaysByArtistID(gomock.Any(), artistID).Return(int64(100), nil)

		res, err := svc.GetArtistByID(context.Background(), artistID)

		assert.NoError(t, err)
		assert.Equal(t, expectedArtist.Name, res.Name)
		assert.Equal(t, int64(100), res.PlayCount)
	})

	t.Run("not found", func(t *testing.T) {
		mockRepo.EXPECT().GetArtistByID(gomock.Any(), artistID).Return(nil, ErrNotFound)

		res, err := svc.GetArtistByID(context.Background(), artistID)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestService_GetAllArtists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo)

	artist1 := model.Artist{ID: uuid.New(), Name: "A1"}
	artist2 := model.Artist{ID: uuid.New(), Name: "A2"}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetAllArtists(gomock.Any(), uint64(10), uint64(0)).
			Return([]model.Artist{artist1, artist2}, nil)

		mockRepo.EXPECT().GetTotalPlaysByArtistIDs(gomock.Any(), gomock.Any()).
			Return(map[uuid.UUID]int64{artist1.ID: 10, artist2.ID: 20}, nil)

		res, err := svc.GetAllArtists(context.Background(), 10, 0)

		assert.NoError(t, err)
		assert.Len(t, res, 2)
		assert.Equal(t, "A1", res[0].Name)
	})
}
