package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"spotify/internal/model"
	"spotify/microservices/catalog/dto"

	"spotify/internal/mocks"

	repository_mock "spotify/microservices/catalog/mocks/repository"
)

func TestService_GetTrackByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockAuth := mocks.NewMockAuthServiceClient(ctrl)
	svc := New(mockRepo, mockAuth)

	trackID := uuid.New()
	albumID := uuid.New()
	artistID := uuid.New()
	genreID := uuid.New()

	trackModel := &model.Track{
		ID:        trackID,
		Title:     "Hit Song",
		DurationS: 180,
		CreatedAt: time.Now(),
	}

	albumModel := model.Album{
		ID:          albumID,
		Title:       "Hit Album",
		ArtistID:    artistID,
		ReleaseDate: time.Now(),
	}

	artistModel := model.Artist{
		ID:   artistID,
		Name: "Super Star",
	}

	genreModel := model.Genre{
		ID:   genreID,
		Name: "Pop",
	}

	t.Run("success with full enrichment", func(t *testing.T) {
		mockRepo.EXPECT().GetTrackByID(gomock.Any(), trackID).Return(trackModel, nil)

		mockRepo.EXPECT().GetAlbumIDsForTracks(gomock.Any(), []uuid.UUID{trackID}).
			Return(map[uuid.UUID]uuid.UUID{trackID: albumID}, nil)

		mockRepo.EXPECT().GetAlbumsByIDs(gomock.Any(), []uuid.UUID{albumID}).
			Return([]model.Album{albumModel}, nil)

		mockRepo.EXPECT().GetArtistsByIDs(gomock.Any(), []uuid.UUID{artistID}).
			Return([]model.Artist{artistModel}, nil).Times(2)

		mockRepo.EXPECT().GetTotalPlaysByArtistIDs(gomock.Any(), []uuid.UUID{artistID}).
			Return(map[uuid.UUID]int64{artistID: 500}, nil).Times(2)

		mockRepo.EXPECT().GetArtistIDsForTracks(gomock.Any(), []uuid.UUID{trackID}).
			Return(map[uuid.UUID][]uuid.UUID{trackID: {artistID}}, nil)

		mockRepo.EXPECT().GetGenresForTracks(gomock.Any(), []uuid.UUID{trackID}).
			Return(map[uuid.UUID][]model.Genre{trackID: {genreModel}}, nil)

		res, err := svc.GetTrackByID(context.Background(), trackID)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "Hit Song", res.Title)
		assert.Equal(t, "Hit Album", res.Album.Title)
		assert.NotEmpty(t, res.Artists)
		assert.Equal(t, "Super Star", res.Artists[0].Name)
	})

	t.Run("track not found", func(t *testing.T) {
		mockRepo.EXPECT().GetTrackByID(gomock.Any(), trackID).Return(nil, ErrNotFound)

		res, err := svc.GetTrackByID(context.Background(), trackID)

		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("enrich failed", func(t *testing.T) {
		mockRepo.EXPECT().GetTrackByID(gomock.Any(), trackID).Return(trackModel, nil)
		mockRepo.EXPECT().GetAlbumIDsForTracks(gomock.Any(), gomock.Any()).Return(nil, errors.New("enrich error"))
		mockRepo.EXPECT().GetArtistIDsForTracks(gomock.Any(), gomock.Any()).Return(nil, nil)
		mockRepo.EXPECT().GetGenresForTracks(gomock.Any(), gomock.Any()).Return(nil, nil)

		_, err := svc.GetTrackByID(context.Background(), trackID)
		assert.Error(t, err)
	})
}

func TestService_RegisterPlay(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockAuth := mocks.NewMockAuthServiceClient(ctrl)
	svc := New(mockRepo, mockAuth)
	trackID := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().IncrementPlayCount(gomock.Any(), trackID).Return(nil)
		err := svc.RegisterPlay(context.Background(), trackID)
		assert.NoError(t, err)
	})
}

func TestService_GetTotalPlays(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockAuth := mocks.NewMockAuthServiceClient(ctrl)
	svc := New(mockRepo, mockAuth)
	artistID := uuid.New()

	t.Run("GetTotalPlaysByArtistID", func(t *testing.T) {
		mockRepo.EXPECT().GetTotalPlaysByArtistID(gomock.Any(), artistID).Return(int64(100), nil)
		count, err := svc.GetTotalPlaysByArtistID(context.Background(), artistID)
		assert.NoError(t, err)
		assert.Equal(t, int64(100), count)
	})

	t.Run("GetTotalPlaysByArtistIDs", func(t *testing.T) {
		mockRepo.EXPECT().GetTotalPlaysByArtistIDs(gomock.Any(), []uuid.UUID{artistID}).
			Return(map[uuid.UUID]int64{artistID: 100}, nil)
		res, err := svc.GetTotalPlaysByArtistIDs(context.Background(), []uuid.UUID{artistID})
		assert.NoError(t, err)
		assert.Equal(t, int64(100), res[artistID])
	})
}

func TestService_GetTracksByReferences(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockAuth := mocks.NewMockAuthServiceClient(ctrl)
	svc := New(mockRepo, mockAuth)

	id := uuid.New()
	trackModel := model.Track{ID: uuid.New(), Title: "T1"}

	t.Run("GetTracksByArtistID", func(t *testing.T) {
		mockRepo.EXPECT().GetTracksByArtistID(gomock.Any(), id, uint64(10), uint64(0)).
			Return([]model.Track{trackModel}, nil)

		mockRepo.EXPECT().GetAlbumIDsForTracks(gomock.Any(), gomock.Any()).Return(nil, nil)
		mockRepo.EXPECT().GetArtistIDsForTracks(gomock.Any(), gomock.Any()).Return(nil, nil)
		mockRepo.EXPECT().GetGenresForTracks(gomock.Any(), gomock.Any()).Return(nil, nil)

		res, err := svc.GetTracksByArtistID(context.Background(), id, 10, 0)
		assert.NoError(t, err)
		assert.Empty(t, res)
	})

	t.Run("GetTracksByAlbumID", func(t *testing.T) {
		mockRepo.EXPECT().GetTracksByAlbumID(gomock.Any(), id, uint64(10), uint64(0)).
			Return([]model.Track{trackModel}, nil)

		mockRepo.EXPECT().GetAlbumIDsForTracks(gomock.Any(), gomock.Any()).Return(nil, nil)
		mockRepo.EXPECT().GetArtistIDsForTracks(gomock.Any(), gomock.Any()).Return(nil, nil)
		mockRepo.EXPECT().GetGenresForTracks(gomock.Any(), gomock.Any()).Return(nil, nil)

		res, err := svc.GetTracksByAlbumID(context.Background(), id, 10, 0)
		assert.NoError(t, err)
		assert.Empty(t, res)
	})

	t.Run("GetTracksByGenreID", func(t *testing.T) {
		mockRepo.EXPECT().GetTracksByGenreID(gomock.Any(), id, uint64(10), uint64(0)).
			Return([]model.Track{trackModel}, nil)

		mockRepo.EXPECT().GetAlbumIDsForTracks(gomock.Any(), gomock.Any()).Return(nil, nil)
		mockRepo.EXPECT().GetArtistIDsForTracks(gomock.Any(), gomock.Any()).Return(nil, nil)
		mockRepo.EXPECT().GetGenresForTracks(gomock.Any(), gomock.Any()).Return(nil, nil)

		res, err := svc.GetTracksByGenreID(context.Background(), id, 10, 0)
		assert.NoError(t, err)
		assert.Empty(t, res)
	})
}

func TestService_GetAllTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockAuth := mocks.NewMockAuthServiceClient(ctrl)
	svc := New(mockRepo, mockAuth)

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetAllTracks(gomock.Any(), uint64(10), uint64(0)).
			Return([]model.Track{}, nil)

		res, err := svc.GetAllTracks(context.Background(), 10, 0)
		assert.NoError(t, err)
		assert.Empty(t, res)
	})
}

func TestService_SearchTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockAuth := mocks.NewMockAuthServiceClient(ctrl)
	svc := New(mockRepo, mockAuth)

	t.Run("success empty", func(t *testing.T) {
		mockRepo.EXPECT().SearchTracks(gomock.Any(), "query", uint64(10)).
			Return([]dto.TrackSearchResult{}, nil)

		res, err := svc.SearchTracks(context.Background(), "query", 10)
		assert.NoError(t, err)
		assert.Empty(t, res)
	})
}

func TestService_GetTracksByIDs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockAuth := mocks.NewMockAuthServiceClient(ctrl)
	svc := New(mockRepo, mockAuth)

	t.Run("empty input", func(t *testing.T) {
		res, err := svc.GetTracksByIDs(context.Background(), []uuid.UUID{})
		assert.NoError(t, err)
		assert.Empty(t, res)
	})

	t.Run("success", func(t *testing.T) {
		id := uuid.New()
		track := model.Track{ID: id, Title: "T"}
		mockRepo.EXPECT().GetTracksByIDs(gomock.Any(), []uuid.UUID{id}).Return([]model.Track{track}, nil)
		mockRepo.EXPECT().GetAlbumIDsForTracks(gomock.Any(), gomock.Any()).Return(nil, nil)
		mockRepo.EXPECT().GetArtistIDsForTracks(gomock.Any(), gomock.Any()).Return(nil, nil)
		mockRepo.EXPECT().GetGenresForTracks(gomock.Any(), gomock.Any()).Return(nil, nil)

		res, err := svc.GetTracksByIDs(context.Background(), []uuid.UUID{id})
		assert.NoError(t, err)
		assert.Empty(t, res)
	})

	t.Run("repo error", func(t *testing.T) {
		mockRepo.EXPECT().GetTracksByIDs(gomock.Any(), gomock.Any()).Return(nil, errors.New("db fail"))
		_, err := svc.GetTracksByIDs(context.Background(), []uuid.UUID{uuid.New()})
		assert.Error(t, err)
	})
}
