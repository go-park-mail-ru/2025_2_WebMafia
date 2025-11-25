package service

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"spotify/internal/model"
	"spotify/microservices/playlist/dto"
	catalog_mock "spotify/microservices/playlist/mocks/catalog"
	repository_mock "spotify/microservices/playlist/mocks/repository"
	storage_mock "spotify/microservices/playlist/mocks/storage"
	"spotify/microservices/playlist/repository/postgres"
	pbCatalog "spotify/proto/catalog"
)

func TestService_CreatePlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo, nil, nil)

	userID := uuid.New()
	req := dto.CreatePlaylistRequest{
		UserID:      userID,
		Title:       "New Playlist",
		Description: "Desc",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().CreatePlaylist(gomock.Any(), gomock.Any(), userID).Return(nil)
		res, err := svc.CreatePlaylist(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, req.Title, res.Title)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().CreatePlaylist(gomock.Any(), gomock.Any(), userID).Return(errors.New("err"))
		_, err := svc.CreatePlaylist(context.Background(), req)
		assert.Error(t, err)
	})
}

func TestService_GetPlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo, nil, nil)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetByID(gomock.Any(), id).Return(&model.Playlist{ID: id, Title: "P"}, nil)
		res, err := svc.GetPlaylist(context.Background(), dto.GetPlaylistRequest{ID: id})
		assert.NoError(t, err)
		assert.Equal(t, "P", res.Title)
	})
}

func TestService_GetPlaylistsByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo, nil, nil)
	uid := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetAllByUser(gomock.Any(), uid, uint64(10), uint64(0)).
			Return([]model.Playlist{{Title: "P1"}}, nil)
		res, err := svc.GetPlaylistsByUser(context.Background(), dto.GetPlaylistsByUserRequest{UserID: uid, Limit: 10})
		assert.NoError(t, err)
		assert.Len(t, res, 1)
	})
}

func TestService_UpdatePlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo, nil, nil)
	id := uuid.New()
	title := "New"

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetByID(gomock.Any(), id).Return(&model.Playlist{ID: id}, nil)
		mockRepo.EXPECT().UpdatePlaylist(gomock.Any(), id, gomock.Any()).Return(nil)

		res, err := svc.UpdatePlaylist(context.Background(), dto.UpdatePlaylistRequest{ID: id, Title: &title})
		assert.NoError(t, err)
		assert.Equal(t, title, res.Title)
	})
}

func TestService_DeletePlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo, nil, nil)
	id := uuid.New()

	mockRepo.EXPECT().DeletePlaylist(gomock.Any(), id).Return(nil)
	err := svc.DeletePlaylist(context.Background(), dto.DeletePlaylistRequest{ID: id})
	assert.NoError(t, err)
}

func TestService_UploadPlaylistAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockStorage := storage_mock.NewMockIStorage(ctrl)
	svc := New(mockRepo, mockStorage, nil)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetByID(gomock.Any(), id).Return(&model.Playlist{ID: id}, nil)
		mockStorage.EXPECT().UploadAvatar(gomock.Any(), gomock.Any(), int64(10), "img/png").Return("url", nil)
		mockRepo.EXPECT().UpdatePlaylistAvatar(gomock.Any(), id, "url").Return(nil)

		res, err := svc.UploadPlaylistAvatar(context.Background(), dto.UploadPlaylistAvatarRequest{PlaylistID: id, Size: 10, ContentType: "img/png"})
		assert.NoError(t, err)
		assert.Equal(t, "url", res.URL)
	})
}

func TestService_DeletePlaylistAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockStorage := storage_mock.NewMockIStorage(ctrl)
	svc := New(mockRepo, mockStorage, nil)
	id := uuid.New()

	mockRepo.EXPECT().GetByID(gomock.Any(), id).Return(&model.Playlist{ID: id, AvatarURL: "old.jpg"}, nil)
	mockStorage.EXPECT().DeleteAvatar(gomock.Any(), "old.jpg").Return(nil)
	mockRepo.EXPECT().UpdatePlaylistAvatar(gomock.Any(), id, "").Return(nil)

	err := svc.DeletePlaylistAvatar(context.Background(), dto.DeletePlaylistAvatarRequest{PlaylistID: id})
	assert.NoError(t, err)
}

func TestService_RemoveTrack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo, nil, nil)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().RemoveTrackFromPlaylist(gomock.Any(), id, "t1").Return(nil)
		err := svc.RemoveTrackFromPlaylist(context.Background(), dto.RemoveTrackFromPlaylistRequest{PlaylistID: id, TrackID: "t1"})
		assert.NoError(t, err)
	})

	t.Run("empty track", func(t *testing.T) {
		err := svc.RemoveTrackFromPlaylist(context.Background(), dto.RemoveTrackFromPlaylistRequest{PlaylistID: id, TrackID: ""})
		assert.Error(t, err)
	})
}

func TestService_AddTrackToPlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockCatalog := catalog_mock.NewMockCatalogServiceClient(ctrl)
	svc := New(mockRepo, nil, mockCatalog)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockCatalog.EXPECT().GetTrackByID(gomock.Any(), gomock.Any()).Return(&pbCatalog.Track{}, nil)
		mockRepo.EXPECT().AddTrackToPlaylist(gomock.Any(), id, "t1").Return(nil)
		err := svc.AddTrackToPlaylist(context.Background(), dto.AddTrackToPlaylistRequest{PlaylistID: id, TrackID: "t1"})
		assert.NoError(t, err)
	})
}

func TestService_GetFavoritePlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo, nil, nil)
	uid := uuid.New()

	t.Run("exists", func(t *testing.T) {
		mockRepo.EXPECT().GetFavoritePlaylist(gomock.Any(), uid).Return(&model.Playlist{Title: "Fav"}, nil)
		res, err := svc.GetFavoritePlaylist(context.Background(), dto.GetFavoritePlaylistRequest{UserID: uid})
		assert.NoError(t, err)
		assert.Equal(t, "Fav", res.Title)
	})

	t.Run("create new", func(t *testing.T) {
		mockRepo.EXPECT().GetFavoritePlaylist(gomock.Any(), uid).Return(nil, postgres.ErrNotFound)
		mockRepo.EXPECT().CreatePlaylist(gomock.Any(), gomock.Any(), uid).Return(nil)
		res, err := svc.GetFavoritePlaylist(context.Background(), dto.GetFavoritePlaylistRequest{UserID: uid})
		assert.NoError(t, err)
		assert.True(t, res.IsFavorite)
	})
}

func TestService_AddTrackToFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockCatalog := catalog_mock.NewMockCatalogServiceClient(ctrl)
	svc := New(mockRepo, nil, mockCatalog)
	uid := uuid.New()

	mockCatalog.EXPECT().GetTrackByID(gomock.Any(), gomock.Any()).Return(&pbCatalog.Track{}, nil)
	mockRepo.EXPECT().GetFavoritePlaylist(gomock.Any(), uid).Return(&model.Playlist{ID: uid}, nil)
	mockRepo.EXPECT().AddTrackToPlaylist(gomock.Any(), uid, "t1").Return(nil)

	err := svc.AddTrackToFavorite(context.Background(), dto.AddTrackToFavoriteRequest{UserID: uid, TrackID: "t1"})
	assert.NoError(t, err)
}

func TestService_GetPlaylistWithTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockCatalog := catalog_mock.NewMockCatalogServiceClient(ctrl)
	svc := New(mockRepo, nil, mockCatalog)
	id := uuid.New()

	mockRepo.EXPECT().GetByID(gomock.Any(), id).Return(&model.Playlist{ID: id}, nil)
	mockRepo.EXPECT().GetTracksByPlaylist(gomock.Any(), id).Return([]string{"t1"}, nil)
	mockCatalog.EXPECT().GetTrackByID(gomock.Any(), gomock.Any()).Return(&pbCatalog.Track{Id: "t1", Title: "S", Album: &pbCatalog.AlbumForTrack{}}, nil)

	res, err := svc.GetPlaylistWithTracks(context.Background(), id)
	assert.NoError(t, err)
	assert.Len(t, res.Tracks, 1)
}
