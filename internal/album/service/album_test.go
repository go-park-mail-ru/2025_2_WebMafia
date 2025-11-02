package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	artist_service "spotify/internal/artist/service"
	mock_album "spotify/internal/mocks/album"
	mock_artist "spotify/internal/mocks/artist"
	"spotify/internal/model"
)

func newModelAlbum(artistID uuid.UUID) *model.Album {
	return &model.Album{
		ID:          uuid.New(),
		Title:       "Test Album",
		ArtistID:    artistID,
		ReleaseDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}

func newArtistModel(id uuid.UUID) *model.Artist {
	return &model.Artist{
		ID:   id,
		Name: "Test Artist",
	}
}

func TestAlbumService_GetAlbumByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumRepo := mock_album.NewMockIRepository(ctrl)
	mockArtistRepo := mock_artist.NewMockIRepository(ctrl)
	mockTrackSvc := mock_artist.NewMockITrackService(ctrl)
	realArtistService := artist_service.New(mockArtistRepo, mockTrackSvc)
	albumService := New(mockAlbumRepo, realArtistService)

	artistModel := newArtistModel(uuid.New())
	albumModel := newModelAlbum(artistModel.ID)

	t.Run("success", func(t *testing.T) {
		mockAlbumRepo.EXPECT().GetByID(gomock.Any(), albumModel.ID).Return(albumModel, nil)
		mockArtistRepo.EXPECT().GetByID(gomock.Any(), albumModel.ArtistID).Return(artistModel, nil)
		mockTrackSvc.EXPECT().GetTotalPlaysByArtistID(gomock.Any(), artistModel.ID).Return(int64(0), nil)

		albumDTO, err := albumService.GetAlbumByID(context.Background(), albumModel.ID)

		assert.NoError(t, err)
		require.NotNil(t, albumDTO)
		assert.Equal(t, albumModel.ID.String(), albumDTO.ID)
		require.Len(t, albumDTO.Artists, 1)
		assert.Equal(t, artistModel.ID.String(), albumDTO.Artists[0].ID)
	})

	t.Run("album repo error", func(t *testing.T) {
		mockAlbumRepo.EXPECT().GetByID(gomock.Any(), albumModel.ID).Return(nil, errors.New("db error"))

		_, err := albumService.GetAlbumByID(context.Background(), albumModel.ID)
		assert.Error(t, err)
	})

	t.Run("artist service error", func(t *testing.T) {
		mockAlbumRepo.EXPECT().GetByID(gomock.Any(), albumModel.ID).Return(albumModel, nil)
		mockArtistRepo.EXPECT().GetByID(gomock.Any(), albumModel.ArtistID).Return(nil, errors.New("artist not found"))

		_, err := albumService.GetAlbumByID(context.Background(), albumModel.ID)
		assert.Error(t, err)
	})
}

func TestAlbumService_GetAllAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumRepo := mock_album.NewMockIRepository(ctrl)
	mockArtistRepo := mock_artist.NewMockIRepository(ctrl)
	mockTrackSvc := mock_artist.NewMockITrackService(ctrl)
	realArtistService := artist_service.New(mockArtistRepo, mockTrackSvc)
	albumService := New(mockAlbumRepo, realArtistService)

	artistModel := newArtistModel(uuid.New())
	albumModel := newModelAlbum(artistModel.ID)

	t.Run("success", func(t *testing.T) {
		albumModels := []model.Album{*albumModel}
		artistModels := []model.Artist{*artistModel}

		mockAlbumRepo.EXPECT().GetAll(gomock.Any(), gomock.Any(), gomock.Any()).Return(albumModels, nil)
		mockArtistRepo.EXPECT().GetByIDs(gomock.Any(), []uuid.UUID{albumModel.ArtistID}).Return(artistModels, nil)
		mockTrackSvc.EXPECT().GetTotalPlaysByArtistIDs(gomock.Any(), gomock.Any()).Return(make(map[uuid.UUID]int64), nil)

		albumDTOs, err := albumService.GetAllAlbums(context.Background(), 10, 0)

		assert.NoError(t, err)
		require.Len(t, albumDTOs, 1)
	})

	t.Run("success empty", func(t *testing.T) {
		mockAlbumRepo.EXPECT().GetAll(gomock.Any(), gomock.Any(), gomock.Any()).Return([]model.Album{}, nil)
		albumDTOs, err := albumService.GetAllAlbums(context.Background(), 10, 0)
		assert.NoError(t, err)
		assert.Len(t, albumDTOs, 0)
	})

	t.Run("album repo error", func(t *testing.T) {
		mockAlbumRepo.EXPECT().GetAll(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("db error"))
		_, err := albumService.GetAllAlbums(context.Background(), 10, 0)
		assert.Error(t, err)
	})

	t.Run("artist service error", func(t *testing.T) {
		mockAlbumRepo.EXPECT().GetAll(gomock.Any(), gomock.Any(), gomock.Any()).Return([]model.Album{*albumModel}, nil)
		mockArtistRepo.EXPECT().GetByIDs(gomock.Any(), gomock.Any()).Return(nil, errors.New("artist error"))
		_, err := albumService.GetAllAlbums(context.Background(), 10, 0)
		assert.Error(t, err)
	})
}

func TestAlbumService_GetAlbumsByArtistID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumRepo := mock_album.NewMockIRepository(ctrl)
	mockArtistRepo := mock_artist.NewMockIRepository(ctrl)
	mockTrackSvc := mock_artist.NewMockITrackService(ctrl)
	realArtistService := artist_service.New(mockArtistRepo, mockTrackSvc)
	albumService := New(mockAlbumRepo, realArtistService)

	artistModel := newArtistModel(uuid.New())
	albumModel := newModelAlbum(artistModel.ID)

	t.Run("success", func(t *testing.T) {
		mockAlbumRepo.EXPECT().GetByArtistID(gomock.Any(), artistModel.ID, gomock.Any(), gomock.Any()).Return([]model.Album{*albumModel}, nil)
		mockArtistRepo.EXPECT().GetByID(gomock.Any(), artistModel.ID).Return(artistModel, nil)
		mockTrackSvc.EXPECT().GetTotalPlaysByArtistID(gomock.Any(), artistModel.ID).Return(int64(0), nil)

		albumDTOs, err := albumService.GetAlbumsByArtistID(context.Background(), artistModel.ID, 10, 0)

		assert.NoError(t, err)
		require.Len(t, albumDTOs, 1)
	})

	t.Run("success empty", func(t *testing.T) {
		mockAlbumRepo.EXPECT().GetByArtistID(gomock.Any(), artistModel.ID, gomock.Any(), gomock.Any()).Return([]model.Album{}, nil)
		albumDTOs, err := albumService.GetAlbumsByArtistID(context.Background(), artistModel.ID, 10, 0)
		assert.NoError(t, err)
		assert.Len(t, albumDTOs, 0)
	})

	t.Run("album repo error", func(t *testing.T) {
		mockAlbumRepo.EXPECT().GetByArtistID(gomock.Any(), artistModel.ID, gomock.Any(), gomock.Any()).Return(nil, errors.New("db error"))
		_, err := albumService.GetAlbumsByArtistID(context.Background(), artistModel.ID, 10, 0)
		assert.Error(t, err)
	})

	t.Run("artist service error", func(t *testing.T) {
		mockAlbumRepo.EXPECT().GetByArtistID(gomock.Any(), artistModel.ID, gomock.Any(), gomock.Any()).Return([]model.Album{*albumModel}, nil)
		mockArtistRepo.EXPECT().GetByID(gomock.Any(), artistModel.ID).Return(nil, errors.New("artist error"))
		_, err := albumService.GetAlbumsByArtistID(context.Background(), artistModel.ID, 10, 0)
		assert.Error(t, err)
	})
}

func TestAlbumService_GetAlbumsByIDs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumRepo := mock_album.NewMockIRepository(ctrl)
	mockArtistRepo := mock_artist.NewMockIRepository(ctrl)
	mockTrackSvc := mock_artist.NewMockITrackService(ctrl)
	realArtistService := artist_service.New(mockArtistRepo, mockTrackSvc)
	albumService := New(mockAlbumRepo, realArtistService)

	artistModel := newArtistModel(uuid.New())
	albumModel := newModelAlbum(artistModel.ID)
	ids := []uuid.UUID{albumModel.ID}

	t.Run("success empty", func(t *testing.T) {
		mockAlbumRepo.EXPECT().GetByIDs(gomock.Any(), ids).Return([]model.Album{}, nil)

		albums, err := albumService.GetAlbumsByIDs(context.Background(), ids)
		assert.NoError(t, err)
		assert.Len(t, albums, 0)
	})

	t.Run("inconsistent data - no artist found", func(t *testing.T) {
		mockAlbumRepo.EXPECT().GetByIDs(gomock.Any(), ids).Return([]model.Album{*albumModel}, nil)
		mockArtistRepo.EXPECT().GetByIDs(gomock.Any(), []uuid.UUID{artistModel.ID}).Return([]model.Artist{}, nil)

		albums, err := albumService.GetAlbumsByIDs(context.Background(), ids)
		assert.NoError(t, err)
		assert.Len(t, albums, 0)
	})
}
