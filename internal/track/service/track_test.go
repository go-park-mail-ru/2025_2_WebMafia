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

	album_service "spotify/internal/album/service"
	artist_service "spotify/internal/artist/service"
	mock_album_repo "spotify/internal/mocks/album"
	mock_artist "spotify/internal/mocks/artist"
	mock_track_repo "spotify/internal/mocks/track"
	"spotify/internal/model"
)

func TestTrackService_GetTrackByID(t *testing.T) {
	ctx := context.Background()

	trackID := uuid.New()
	albumID := uuid.New()
	artistID := uuid.New()

	track := &model.Track{ID: trackID, Title: "Test Track"}
	album := &model.Album{ID: albumID, Title: "Test Album", ArtistID: artistID, ReleaseDate: time.Now()}
	artist := &model.Artist{ID: artistID, Name: "Test Artist"}
	genre := model.Genre{ID: uuid.New(), Name: "Rock"}

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockTrackRepo := mock_track_repo.NewMockIRepository(ctrl)
		mockAlbumRepo := mock_album_repo.NewMockIRepository(ctrl)
		mockArtistRepo := mock_artist.NewMockIRepository(ctrl)
		mockTrackSvcForArtist := mock_artist.NewMockITrackService(ctrl)
		realArtistService := artist_service.New(mockArtistRepo, mockTrackSvcForArtist)
		realAlbumService := album_service.New(mockAlbumRepo, realArtistService)
		trackService := New(mockTrackRepo, realAlbumService, realArtistService)

		mockTrackRepo.EXPECT().GetByID(ctx, trackID).Return(track, nil)
		mockTrackRepo.EXPECT().GetAlbumIDsForTracks(ctx, []uuid.UUID{trackID}).Return(map[uuid.UUID]uuid.UUID{trackID: albumID}, nil)
		mockTrackRepo.EXPECT().GetArtistIDsForTracks(ctx, []uuid.UUID{trackID}).Return(map[uuid.UUID][]uuid.UUID{trackID: {artistID}}, nil)
		mockTrackRepo.EXPECT().GetGenresForTracks(ctx, []uuid.UUID{trackID}).Return(map[uuid.UUID][]model.Genre{trackID: {genre}}, nil)
		mockAlbumRepo.EXPECT().GetByIDs(ctx, []uuid.UUID{albumID}).Return([]model.Album{*album}, nil)
		mockArtistRepo.EXPECT().GetByIDs(ctx, gomock.Any()).Return([]model.Artist{*artist}, nil).AnyTimes()
		mockTrackSvcForArtist.EXPECT().GetTotalPlaysByArtistIDs(gomock.Any(), gomock.Any()).Return(make(map[uuid.UUID]int64), nil).AnyTimes()

		trackDTO, err := trackService.GetTrackByID(ctx, trackID)

		assert.NoError(t, err)
		require.NotNil(t, trackDTO)
		assert.Equal(t, trackID.String(), trackDTO.ID)
		assert.Equal(t, albumID.String(), trackDTO.Album.ID)
		require.Len(t, trackDTO.Artists, 1)
		assert.Equal(t, artistID.String(), trackDTO.Artists[0].ID)
		require.Len(t, trackDTO.Genres, 1)
		assert.Equal(t, genre.ID.String(), trackDTO.Genres[0].ID)
	})

	t.Run("track not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockTrackRepo := mock_track_repo.NewMockIRepository(ctrl)
		trackService := New(mockTrackRepo, nil, nil)

		mockTrackRepo.EXPECT().GetByID(ctx, trackID).Return(nil, errors.New("not found"))

		_, err := trackService.GetTrackByID(ctx, trackID)
		assert.Error(t, err)
	})

	t.Run("enrichment fails - album error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockTrackRepo := mock_track_repo.NewMockIRepository(ctrl)
		mockAlbumRepo := mock_album_repo.NewMockIRepository(ctrl)
		mockArtistRepo := mock_artist.NewMockIRepository(ctrl)
		mockTrackSvcForArtist := mock_artist.NewMockITrackService(ctrl)
		realArtistService := artist_service.New(mockArtistRepo, mockTrackSvcForArtist)
		realAlbumService := album_service.New(mockAlbumRepo, realArtistService)
		trackService := New(mockTrackRepo, realAlbumService, realArtistService)

		mockTrackRepo.EXPECT().GetByID(ctx, trackID).Return(track, nil)

		mockTrackRepo.EXPECT().GetAlbumIDsForTracks(ctx, gomock.Any()).Return(nil, errors.New("db error"))

		mockTrackRepo.EXPECT().GetArtistIDsForTracks(ctx, gomock.Any()).Return(nil, nil).AnyTimes()
		mockTrackRepo.EXPECT().GetGenresForTracks(ctx, gomock.Any()).Return(nil, nil).AnyTimes()

		_, err := trackService.GetTrackByID(ctx, trackID)
		assert.Error(t, err)
	})

	t.Run("enrichment fails - inconsistent data (no album)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockTrackRepo := mock_track_repo.NewMockIRepository(ctrl)
		mockArtistRepo := mock_artist.NewMockIRepository(ctrl)
		mockTrackSvcForArtist := mock_artist.NewMockITrackService(ctrl)
		realArtistService := artist_service.New(mockArtistRepo, mockTrackSvcForArtist)
		trackService := New(mockTrackRepo, nil, realArtistService)

		mockTrackRepo.EXPECT().GetByID(ctx, trackID).Return(track, nil)
		mockTrackRepo.EXPECT().GetAlbumIDsForTracks(ctx, gomock.Any()).Return(map[uuid.UUID]uuid.UUID{}, nil)
		mockTrackRepo.EXPECT().GetArtistIDsForTracks(ctx, gomock.Any()).Return(map[uuid.UUID][]uuid.UUID{trackID: {artistID}}, nil).AnyTimes()
		mockTrackRepo.EXPECT().GetGenresForTracks(ctx, gomock.Any()).Return(map[uuid.UUID][]model.Genre{trackID: {genre}}, nil).AnyTimes()
		mockArtistRepo.EXPECT().GetByIDs(ctx, gomock.Any()).Return([]model.Artist{*artist}, nil).AnyTimes()
		mockTrackSvcForArtist.EXPECT().GetTotalPlaysByArtistIDs(gomock.Any(), gomock.Any()).Return(make(map[uuid.UUID]int64), nil).AnyTimes()

		_, err := trackService.GetTrackByID(ctx, trackID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "track data is inconsistent")
	})
}

func TestTrackService_GetTracksMethods(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	limit, offset := uint64(10), uint64(0)

	t.Run("GetAllTracks - repo error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockTrackRepo := mock_track_repo.NewMockIRepository(ctrl)
		trackService := New(mockTrackRepo, nil, nil)
		mockTrackRepo.EXPECT().GetAll(ctx, limit, offset).Return(nil, errors.New("db error"))
		_, err := trackService.GetAllTracks(ctx, limit, offset)
		assert.Error(t, err)
	})

	t.Run("GetAllTracks - success empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockTrackRepo := mock_track_repo.NewMockIRepository(ctrl)
		trackService := New(mockTrackRepo, nil, nil)
		mockTrackRepo.EXPECT().GetAll(ctx, limit, offset).Return([]model.Track{}, nil)
		tracks, err := trackService.GetAllTracks(ctx, limit, offset)
		assert.NoError(t, err)
		assert.Len(t, tracks, 0)
	})

	t.Run("GetTracksByArtistID - repo error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockTrackRepo := mock_track_repo.NewMockIRepository(ctrl)
		trackService := New(mockTrackRepo, nil, nil)
		mockTrackRepo.EXPECT().GetByArtistID(ctx, id, limit, offset).Return(nil, errors.New("db error"))
		_, err := trackService.GetTracksByArtistID(ctx, id, limit, offset)
		assert.Error(t, err)
	})

	t.Run("GetTracksByAlbumID - repo error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockTrackRepo := mock_track_repo.NewMockIRepository(ctrl)
		trackService := New(mockTrackRepo, nil, nil)
		mockTrackRepo.EXPECT().GetByAlbumID(ctx, id, limit, offset).Return(nil, errors.New("db error"))
		_, err := trackService.GetTracksByAlbumID(ctx, id, limit, offset)
		assert.Error(t, err)
	})

	t.Run("GetTracksByGenreID - repo error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockTrackRepo := mock_track_repo.NewMockIRepository(ctrl)
		trackService := New(mockTrackRepo, nil, nil)
		mockTrackRepo.EXPECT().GetByGenreID(ctx, id, limit, offset).Return(nil, errors.New("db error"))
		_, err := trackService.GetTracksByGenreID(ctx, id, limit, offset)
		assert.Error(t, err)
	})
}

func TestTrackService_EnrichmentFailures(t *testing.T) {
	ctx := context.Background()
	tracks := []model.Track{{ID: uuid.New(), Title: "Test Track"}}
	trackIDs := []uuid.UUID{tracks[0].ID}

	t.Run("enrich - artist service fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockTrackRepo := mock_track_repo.NewMockIRepository(ctrl)
		mockAlbumRepo := mock_album_repo.NewMockIRepository(ctrl)
		mockArtistRepo := mock_artist.NewMockIRepository(ctrl)
		mockTrackSvcForArtist := mock_artist.NewMockITrackService(ctrl)
		realArtistService := artist_service.New(mockArtistRepo, mockTrackSvcForArtist)
		realAlbumService := album_service.New(mockAlbumRepo, realArtistService)
		trackService := New(mockTrackRepo, realAlbumService, realArtistService)

		artistID := uuid.New()
		expectedError := errors.New("artist service error")

		mockTrackRepo.EXPECT().GetAll(ctx, gomock.Any(), gomock.Any()).Return(tracks, nil)

		mockTrackRepo.EXPECT().GetAlbumIDsForTracks(ctx, trackIDs).Return(map[uuid.UUID]uuid.UUID{}, nil).AnyTimes()
		mockTrackRepo.EXPECT().GetArtistIDsForTracks(ctx, trackIDs).Return(map[uuid.UUID][]uuid.UUID{trackIDs[0]: {artistID}}, nil).AnyTimes()
		mockTrackRepo.EXPECT().GetGenresForTracks(ctx, trackIDs).Return(map[uuid.UUID][]model.Genre{}, nil).AnyTimes()

		mockArtistRepo.EXPECT().GetByIDs(ctx, []uuid.UUID{artistID}).Return(nil, expectedError)

		_, err := trackService.GetAllTracks(ctx, 10, 0)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), expectedError.Error())
	})

	t.Run("getAlbumsForTracks - success empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockTrackRepo := mock_track_repo.NewMockIRepository(ctrl)
		trackService := New(mockTrackRepo, nil, nil)

		mockTrackRepo.EXPECT().GetAlbumIDsForTracks(ctx, trackIDs).Return(map[uuid.UUID]uuid.UUID{}, nil)
		albums, err := trackService.getAlbumsForTracks(ctx, trackIDs)
		assert.NoError(t, err)
		assert.Empty(t, albums)
	})

	t.Run("getArtistsForTracks - success empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockTrackRepo := mock_track_repo.NewMockIRepository(ctrl)
		trackService := New(mockTrackRepo, nil, nil)

		mockTrackRepo.EXPECT().GetArtistIDsForTracks(ctx, trackIDs).Return(map[uuid.UUID][]uuid.UUID{}, nil)
		artists, err := trackService.getArtistsForTracks(ctx, trackIDs)
		assert.NoError(t, err)
		assert.Empty(t, artists)
	})
}
