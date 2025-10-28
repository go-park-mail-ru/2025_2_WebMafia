package service

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"spotify/internal/artist/repository/postgres"
	mock_artist "spotify/internal/mocks/artist"
	"spotify/internal/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func newModelArtist() *model.Artist {
	return &model.Artist{
		ID:          uuid.New(),
		Name:        "Test Artist",
		AvatarURL:   "http://example.com/avatar.jpg",
		Description: sql.NullString{String: "A test artist", Valid: true},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func TestArtistService_GetArtistByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_artist.NewMockIRepository(ctrl)
	service := New(mockRepo)
	artistModel := newModelArtist()

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetByID(gomock.Any(), artistModel.ID).Return(artistModel, nil)

		artistDTO, err := service.GetArtistByID(context.Background(), artistModel.ID)

		assert.NoError(t, err)
		require.NotNil(t, artistDTO)
		assert.Equal(t, artistModel.ID.String(), artistDTO.ID)
		assert.Equal(t, artistModel.Name, artistDTO.Name)
		assert.Equal(t, artistModel.AvatarURL, artistDTO.AvatarURL)
	})

	t.Run("not found", func(t *testing.T) {
		mockRepo.EXPECT().GetByID(gomock.Any(), artistModel.ID).Return(nil, postgres.ErrNotFound)

		artistDTO, err := service.GetArtistByID(context.Background(), artistModel.ID)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrNotFound)
		assert.Nil(t, artistDTO)
	})

	t.Run("db error", func(t *testing.T) {
		expectedError := errors.New("unexpected db error")
		mockRepo.EXPECT().GetByID(gomock.Any(), artistModel.ID).Return(nil, expectedError)

		_, err := service.GetArtistByID(context.Background(), artistModel.ID)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
	})
}

func TestArtistService_GetAllArtists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_artist.NewMockIRepository(ctrl)
	service := New(mockRepo)
	artistModels := []model.Artist{*newModelArtist(), *newModelArtist()}
	limit, offset := uint64(10), uint64(0)

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetAll(gomock.Any(), limit, offset).Return(artistModels, nil)

		artistDTOs, err := service.GetAllArtists(context.Background(), limit, offset)

		assert.NoError(t, err)
		require.Len(t, artistDTOs, 2)
		assert.Equal(t, artistModels[0].ID.String(), artistDTOs[0].ID)
		assert.Equal(t, artistModels[1].ID.String(), artistDTOs[1].ID)
	})

	t.Run("success - empty result", func(t *testing.T) {
		mockRepo.EXPECT().GetAll(gomock.Any(), limit, offset).Return([]model.Artist{}, nil)

		artistDTOs, err := service.GetAllArtists(context.Background(), limit, offset)

		assert.NoError(t, err)
		assert.Len(t, artistDTOs, 0)
	})

	t.Run("repository error", func(t *testing.T) {
		expectedError := errors.New("repository error")
		mockRepo.EXPECT().GetAll(gomock.Any(), limit, offset).Return(nil, expectedError)

		_, err := service.GetAllArtists(context.Background(), limit, offset)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
	})
}

func TestArtistService_GetArtistsByIDs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_artist.NewMockIRepository(ctrl)
	service := New(mockRepo)
	artistModel1 := newModelArtist()
	artistModel2 := newModelArtist()
	artists := []model.Artist{*artistModel1, *artistModel2}
	ids := []uuid.UUID{artistModel1.ID, artistModel2.ID}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetByIDs(gomock.Any(), ids).Return(artists, nil)

		artistDTOs, err := service.GetArtistsByIDs(context.Background(), ids)

		assert.NoError(t, err)
		require.Len(t, artistDTOs, 2)
		assert.Equal(t, artistModel1.ID.String(), artistDTOs[0].ID)
		assert.Equal(t, artistModel2.ID.String(), artistDTOs[1].ID)
	})

	t.Run("success - empty result", func(t *testing.T) {
		mockRepo.EXPECT().GetByIDs(gomock.Any(), ids).Return([]model.Artist{}, nil)

		artistDTOs, err := service.GetArtistsByIDs(context.Background(), ids)

		assert.NoError(t, err)
		assert.Len(t, artistDTOs, 0)
	})

	t.Run("repository error", func(t *testing.T) {
		expectedError := errors.New("repository error")
		mockRepo.EXPECT().GetByIDs(gomock.Any(), ids).Return(nil, expectedError)

		_, err := service.GetArtistsByIDs(context.Background(), ids)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
	})
}
