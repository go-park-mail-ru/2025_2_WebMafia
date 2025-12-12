package service

import (
	"context"
	"errors"
	"spotify/microservices/playlist/ai"
	ai_mock "spotify/microservices/playlist/mocks/ai"
	"testing"
	"time"

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
	svc := New(mockRepo, nil, nil, nil)

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

func TestService_GetPlaylistsByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo, nil, nil, nil)
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
	svc := New(mockRepo, nil, nil, nil)
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
	svc := New(mockRepo, nil, nil, nil)
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
	svc := New(mockRepo, mockStorage, nil, nil)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetByID(gomock.Any(), id).Return(&model.Playlist{ID: id}, nil)
		mockStorage.EXPECT().UploadAvatar(gomock.Any(), gomock.Any(), int64(10), "img/png").Return("url", nil)
		mockRepo.EXPECT().UpdatePlaylistAvatar(gomock.Any(), id, "url").Return(nil)

		res, err := svc.UploadPlaylistAvatar(context.Background(), dto.UploadPlaylistAvatarRequest{
			PlaylistID:  id,
			Size:        10,
			ContentType: "img/png",
		})
		assert.NoError(t, err)
		assert.Equal(t, "url", res.URL)
	})
}

func TestService_DeletePlaylistAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockStorage := storage_mock.NewMockIStorage(ctrl)
	svc := New(mockRepo, mockStorage, nil, nil)
	id := uuid.New()

	mockRepo.EXPECT().GetByID(gomock.Any(), id).Return(&model.Playlist{
		ID:        id,
		AvatarURL: "old.jpg",
	}, nil)
	mockStorage.EXPECT().DeleteAvatar(gomock.Any(), "old.jpg").Return(nil)
	mockRepo.EXPECT().UpdatePlaylistAvatar(gomock.Any(), id, "").Return(nil)

	err := svc.DeletePlaylistAvatar(context.Background(), dto.DeletePlaylistAvatarRequest{PlaylistID: id})
	assert.NoError(t, err)
}

func TestService_RemoveTrack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo, nil, nil, nil)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().RemoveTrackFromPlaylist(gomock.Any(), id, "t1").Return(nil)
		err := svc.RemoveTrackFromPlaylist(context.Background(), dto.RemoveTrackFromPlaylistRequest{
			PlaylistID: id,
			TrackID:    "t1",
		})
		assert.NoError(t, err)
	})

	t.Run("empty track", func(t *testing.T) {
		err := svc.RemoveTrackFromPlaylist(context.Background(), dto.RemoveTrackFromPlaylistRequest{
			PlaylistID: id,
			TrackID:    "",
		})
		assert.Error(t, err)
	})
}

func TestService_AddTrackToPlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo, nil, nil, nil)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().AddTrackToPlaylist(gomock.Any(), id, "t1").Return(nil)
		err := svc.AddTrackToPlaylist(context.Background(), dto.AddTrackToPlaylistRequest{
			PlaylistID: id,
			TrackID:    "t1",
		})
		assert.NoError(t, err)
	})
}

func TestService_GetFavoritePlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo, nil, nil, nil)
	uid := uuid.New()

	t.Run("exists", func(t *testing.T) {
		mockRepo.EXPECT().GetFavoritePlaylist(gomock.Any(), uid).
			Return(&model.Playlist{Title: "Fav"}, nil)

		res, err := svc.GetFavoritePlaylist(context.Background(), dto.GetFavoritePlaylistRequest{UserID: uid})
		assert.NoError(t, err)
		assert.Equal(t, "Fav", res.Title)
	})

	t.Run("create new", func(t *testing.T) {
		mockRepo.EXPECT().GetFavoritePlaylist(gomock.Any(), uid).
			Return(nil, postgres.ErrNotFound)
		mockRepo.EXPECT().CreatePlaylist(gomock.Any(), gomock.Any(), uid).
			Return(nil)

		res, err := svc.GetFavoritePlaylist(context.Background(), dto.GetFavoritePlaylistRequest{UserID: uid})
		assert.NoError(t, err)
		assert.True(t, res.IsFavorite)
	})
}

func TestService_AddTrackToFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo, nil, nil, nil)
	uid := uuid.New()

	favPlaylistID := uuid.New()

	mockRepo.EXPECT().
		GetFavoritePlaylist(gomock.Any(), uid).
		Return(&model.Playlist{ID: favPlaylistID}, nil)

	mockRepo.EXPECT().
		AddTrackToPlaylist(gomock.Any(), favPlaylistID, "t1").
		Return(nil)

	err := svc.AddTrackToFavorite(context.Background(), dto.AddTrackToFavoriteRequest{
		UserID:  uid,
		TrackID: "t1",
	})
	assert.NoError(t, err)
}

func TestService_GetPlaylistWithTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockCatalog := catalog_mock.NewMockCatalogServiceClient(ctrl)
	svc := New(mockRepo, nil, mockCatalog, nil)
	id := uuid.New()

	mockRepo.EXPECT().
		GetByID(gomock.Any(), id).
		Return(&model.Playlist{ID: id}, nil)

	mockRepo.EXPECT().
		GetTracksByPlaylist(gomock.Any(), id).
		Return([]string{"t1"}, nil)

	mockCatalog.EXPECT().
		GetTracksByIDs(gomock.Any(), &pbCatalog.GetTracksByIDsRequest{
			Ids: []string{"t1"},
		}).
		Return(&pbCatalog.GetTracksByIDsResponse{
			Tracks: []*pbCatalog.Track{
				{
					Id:        "t1",
					Title:     "S",
					DurationS: 123,
					FileUrl:   "url",
					Album: &pbCatalog.AlbumForTrack{
						Id:        "a1",
						Title:     "Album",
						AvatarUrl: "img",
					},
					Artists: []*pbCatalog.ArtistForTrack{
						{
							Id:   "art1",
							Name: "Artist",
						},
					},
				},
			},
		}, nil)

	res, err := svc.GetPlaylistWithTracks(context.Background(), id)

	assert.NoError(t, err)
	assert.Len(t, res.Tracks, 1)
	assert.Equal(t, "t1", res.Tracks[0].ID)
}

func TestService_AddAlbumToFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	svc := New(repo, nil, nil, nil)

	uid := uuid.New()

	repo.EXPECT().
		AddAlbumToFavorite(gomock.Any(), uid, "a1").
		Return(nil)

	err := svc.AddAlbumToFavorite(context.Background(), dto.AddAlbumToFavoriteRequest{
		UserID:  uid,
		AlbumID: "a1",
	})
	assert.NoError(t, err)
}

func TestService_RemoveAlbumFromFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	svc := New(repo, nil, nil, nil)

	uid := uuid.New()

	repo.EXPECT().
		RemoveAlbumFromFavorite(gomock.Any(), uid, "a1").
		Return(nil)

	err := svc.RemoveAlbumFromFavorite(context.Background(), dto.RemoveAlbumFromFavoriteRequest{
		UserID:  uid,
		AlbumID: "a1",
	})
	assert.NoError(t, err)
}

func TestService_AddArtistToFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	svc := New(repo, nil, nil, nil)

	uid := uuid.New()

	repo.EXPECT().
		AddArtistToFavorite(gomock.Any(), uid, "art1").
		Return(nil)

	err := svc.AddArtistToFavorite(context.Background(), dto.AddArtistToFavoriteRequest{
		UserID:   uid,
		ArtistID: "art1",
	})
	assert.NoError(t, err)
}

func TestService_RemoveArtistFromFavorite_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	svc := New(repo, nil, nil, nil)

	uid := uuid.New()

	repo.EXPECT().
		RemoveArtistFromFavorite(gomock.Any(), uid, "art1").
		Return(nil)

	err := svc.RemoveArtistFromFavorite(context.Background(), dto.RemoveArtistFromFavoriteRequest{
		UserID:   uid,
		ArtistID: "art1",
	})
	assert.NoError(t, err)
}

func TestService_GetFavoriteArtists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	cat := catalog_mock.NewMockCatalogServiceClient(ctrl)
	svc := New(repo, nil, cat, nil)

	uid := uuid.New()

	artistID := uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb")
	now := time.Now()

	repo.EXPECT().
		GetFavoriteArtistIDs(gomock.Any(), uid).
		Return([]model.FavoriteArtist{
			{UserID: uid, ArtistID: artistID, CreatedAt: now},
		}, nil)

	cat.EXPECT().
		GetArtistsByIDs(gomock.Any(), &pbCatalog.GetArtistsByIDsRequest{
			Ids: []string{artistID.String()},
		}).
		Return(&pbCatalog.GetArtistsByIDsResponse{
			Artists: []*pbCatalog.Artist{
				{
					Id:        artistID.String(),
					Name:      "A1",
					AvatarUrl: "img1",
				},
			},
		}, nil)

	res, err := svc.GetFavoriteArtists(context.Background(), uid)

	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, artistID.String(), res[0].ID)
	assert.Equal(t, "A1", res[0].Name)
}

func TestService_AddAlbumToFavorite_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	svc := New(repo, nil, nil, nil)
	uid := uuid.New()

	repo.EXPECT().
		AddAlbumToFavorite(gomock.Any(), uid, "a1").
		Return(errors.New("err"))

	err := svc.AddAlbumToFavorite(context.Background(), dto.AddAlbumToFavoriteRequest{
		UserID:  uid,
		AlbumID: "a1",
	})
	assert.Error(t, err)
}

func TestService_GetFavoriteAlbums_CatalogError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	cat := catalog_mock.NewMockCatalogServiceClient(ctrl)
	svc := New(repo, nil, cat, nil)

	uid := uuid.New()
	albumID := uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")

	repo.EXPECT().
		GetFavoriteAlbumIDs(gomock.Any(), uid).
		Return([]model.FavoriteAlbum{
			{UserID: uid, AlbumID: albumID, CreatedAt: time.Now()},
		}, nil)

	cat.EXPECT().
		GetAlbumsByIDs(gomock.Any(), &pbCatalog.GetAlbumsByIDsRequest{
			Ids: []string{albumID.String()},
		}).
		Return(nil, errors.New("cat"))

	_, err := svc.GetFavoriteAlbums(context.Background(), uid)
	assert.Error(t, err)
}

func TestService_GetFavoriteArtists_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	svc := New(repo, nil, nil, nil)
	uid := uuid.New()

	repo.EXPECT().
		GetFavoriteArtistIDs(gomock.Any(), uid).
		Return(nil, errors.New("db"))

	_, err := svc.GetFavoriteArtists(context.Background(), uid)
	assert.Error(t, err)
}

func TestService_GetFavoriteArtists_CatalogError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	cat := catalog_mock.NewMockCatalogServiceClient(ctrl)
	svc := New(repo, nil, cat, nil)

	uid := uuid.New()
	artistID := uuid.New()

	repo.EXPECT().
		GetFavoriteArtistIDs(gomock.Any(), uid).
		Return([]model.FavoriteArtist{
			{UserID: uid, ArtistID: artistID, CreatedAt: time.Now()},
		}, nil)

	cat.EXPECT().
		GetArtistsByIDs(gomock.Any(), &pbCatalog.GetArtistsByIDsRequest{
			Ids: []string{artistID.String()},
		}).
		Return(nil, errors.New("cat"))

	_, err := svc.GetFavoriteArtists(context.Background(), uid)
	assert.Error(t, err)
}

func TestService_UpdatePlaylist_GetByID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo, nil, nil, nil)
	id := uuid.New()

	mockRepo.EXPECT().
		GetByID(gomock.Any(), id).
		Return(nil, errors.New("repo error"))

	_, err := svc.UpdatePlaylist(context.Background(), dto.UpdatePlaylistRequest{ID: id})
	assert.Error(t, err)
}

func TestService_UpdatePlaylist_Update_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	svc := New(mockRepo, nil, nil, nil)
	id := uuid.New()
	title := "X"

	mockRepo.EXPECT().
		GetByID(gomock.Any(), id).
		Return(&model.Playlist{ID: id}, nil)

	mockRepo.EXPECT().
		UpdatePlaylist(gomock.Any(), id, gomock.Any()).
		Return(errors.New("update err"))

	_, err := svc.UpdatePlaylist(context.Background(), dto.UpdatePlaylistRequest{ID: id, Title: &title})
	assert.Error(t, err)
}

func TestService_UploadPlaylistAvatar_GetByID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockStorage := storage_mock.NewMockIStorage(ctrl)
	svc := New(mockRepo, mockStorage, nil, nil)

	id := uuid.New()

	mockRepo.EXPECT().
		GetByID(gomock.Any(), id).
		Return(nil, errors.New("not found"))

	_, err := svc.UploadPlaylistAvatar(context.Background(), dto.UploadPlaylistAvatarRequest{
		PlaylistID: id,
		Size:       5,
	})
	assert.Error(t, err)
}

func TestService_UploadPlaylistAvatar_Upload_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockStorage := storage_mock.NewMockIStorage(ctrl)
	svc := New(mockRepo, mockStorage, nil, nil)
	id := uuid.New()

	mockRepo.EXPECT().GetByID(gomock.Any(), id).
		Return(&model.Playlist{ID: id}, nil)

	mockStorage.EXPECT().
		UploadAvatar(gomock.Any(), gomock.Any(), int64(10), "png").
		Return("", errors.New("upload err"))

	_, err := svc.UploadPlaylistAvatar(context.Background(), dto.UploadPlaylistAvatarRequest{
		PlaylistID:  id,
		Size:        10,
		ContentType: "png",
	})
	assert.Error(t, err)
}

func TestService_GetPlaylistWithTracks_GetTracks_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockCatalog := catalog_mock.NewMockCatalogServiceClient(ctrl)
	svc := New(mockRepo, nil, mockCatalog, nil)

	id := uuid.New()

	mockRepo.EXPECT().GetByID(gomock.Any(), id).
		Return(&model.Playlist{ID: id}, nil)

	mockRepo.EXPECT().GetTracksByPlaylist(gomock.Any(), id).
		Return(nil, errors.New("tracks err"))

	_, err := svc.GetPlaylistWithTracks(context.Background(), id)
	assert.Error(t, err)
}

func TestService_GetFavoriteAlbums_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockIRepository(ctrl)
	mockCatalog := catalog_mock.NewMockCatalogServiceClient(ctrl)
	svc := New(mockRepo, nil, mockCatalog, nil)

	uid := uuid.New()

	mockRepo.EXPECT().
		GetFavoriteAlbumIDs(gomock.Any(), uid).
		Return(nil, errors.New("db err"))

	_, err := svc.GetFavoriteAlbums(context.Background(), uid)
	assert.Error(t, err)
}

func TestService_GeneratePlaylistMeta_AI_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	cat := catalog_mock.NewMockCatalogServiceClient(ctrl)
	aiMock := ai_mock.NewMockIAIGenerator(ctrl)

	svc := New(repo, nil, cat, aiMock)

	pid := uuid.New()

	repo.EXPECT().
		GetByID(gomock.Any(), pid).
		Return(&model.Playlist{ID: pid, Title: "Old"}, nil)

	repo.EXPECT().
		GetTracksByPlaylist(gomock.Any(), pid).
		Return([]string{"t1"}, nil)

	cat.EXPECT().
		GetTracksByIDs(gomock.Any(), gomock.Any()).
		Return(&pbCatalog.GetTracksByIDsResponse{
			Tracks: []*pbCatalog.Track{
				{
					Id:    "t1",
					Title: "Song",
					Album: &pbCatalog.AlbumForTrack{
						Id:    "a1",
						Title: "Album",
					},
					Artists: []*pbCatalog.ArtistForTrack{
						{
							Id:   "art1",
							Name: "Artist",
						},
					},
				},
			},
		}, nil)

	aiMock.EXPECT().
		GeneratePlaylistMeta(gomock.Any(), gomock.Any()).
		Return("AI title", "AI desc", nil)

	meta, err := svc.GeneratePlaylistMeta(context.Background(), pid)

	assert.NoError(t, err)
	assert.Equal(t, "AI title", meta.Title)
	assert.Equal(t, "ai", meta.Source)
}

func TestService_GeneratePlaylistMeta_Fallback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	cat := catalog_mock.NewMockCatalogServiceClient(ctrl)
	aiMock := ai_mock.NewMockIAIGenerator(ctrl)

	svc := New(repo, nil, cat, aiMock)

	pid := uuid.New()

	repo.EXPECT().
		GetByID(gomock.Any(), pid).
		Return(&model.Playlist{ID: pid}, nil)

	repo.EXPECT().
		GetTracksByPlaylist(gomock.Any(), pid).
		Return([]string{"t1"}, nil)

	cat.EXPECT().
		GetTracksByIDs(gomock.Any(), gomock.Any()).
		Return(&pbCatalog.GetTracksByIDsResponse{
			Tracks: []*pbCatalog.Track{
				{
					Id:    "t1",
					Title: "Song",
					Album: &pbCatalog.AlbumForTrack{
						Id:    "a1",
						Title: "Album",
					},
					Artists: []*pbCatalog.ArtistForTrack{
						{
							Id:   "art1",
							Name: "Artist",
						},
					},
				},
			},
		}, nil)

	aiMock.EXPECT().
		GeneratePlaylistMeta(gomock.Any(), gomock.Any()).
		Return("", "", ai.ErrAIRateLimit)

	meta, err := svc.GeneratePlaylistMeta(context.Background(), pid)

	assert.NoError(t, err)
	assert.Equal(t, "fallback", meta.Source)
}

func TestService_ConfirmPlaylistMeta_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	svc := New(repo, nil, nil, nil)

	pid := uuid.New()

	repo.EXPECT().
		UpdatePlaylist(gomock.Any(), pid, gomock.Any()).
		Return(nil)

	err := svc.ConfirmPlaylistMeta(
		context.Background(),
		pid,
		"title",
		"desc",
	)

	assert.NoError(t, err)
}

func TestService_ConfirmPlaylistMeta_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	svc := New(repo, nil, nil, nil)

	pid := uuid.New()

	repo.EXPECT().
		UpdatePlaylist(gomock.Any(), pid, gomock.Any()).
		Return(errors.New("db err"))

	err := svc.ConfirmPlaylistMeta(
		context.Background(),
		pid,
		"title",
		"desc",
	)

	assert.Error(t, err)
}

func TestService_GeneratePlaylistMeta_EmptyTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	svc := New(repo, nil, nil, nil)

	pid := uuid.New()

	repo.EXPECT().
		GetByID(gomock.Any(), pid).
		Return(&model.Playlist{
			ID:          pid,
			Title:       "My",
			Description: "Desc",
		}, nil)

	repo.EXPECT().
		GetTracksByPlaylist(gomock.Any(), pid).
		Return([]string{}, nil)

	meta, err := svc.GeneratePlaylistMeta(context.Background(), pid)

	assert.NoError(t, err)
	assert.Equal(t, "My", meta.Title)
	assert.Equal(t, "Desc", meta.Description)
}

func TestService_GeneratePlaylistMeta_CatalogError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	cat := catalog_mock.NewMockCatalogServiceClient(ctrl)

	svc := New(repo, nil, cat, nil)

	pid := uuid.New()

	repo.EXPECT().
		GetByID(gomock.Any(), pid).
		Return(&model.Playlist{ID: pid}, nil)

	repo.EXPECT().
		GetTracksByPlaylist(gomock.Any(), pid).
		Return([]string{"t1"}, nil)

	cat.EXPECT().
		GetTracksByIDs(gomock.Any(), gomock.Any()).
		Return(nil, errors.New("catalog down"))

	_, err := svc.GeneratePlaylistMeta(context.Background(), pid)

	assert.Error(t, err)
}

func TestService_GeneratePlaylistMeta_AIAuthError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository_mock.NewMockIRepository(ctrl)
	cat := catalog_mock.NewMockCatalogServiceClient(ctrl)
	aiMock := ai_mock.NewMockIAIGenerator(ctrl)

	svc := New(repo, nil, cat, aiMock)

	pid := uuid.New()

	repo.EXPECT().GetByID(gomock.Any(), pid).
		Return(&model.Playlist{ID: pid}, nil)

	repo.EXPECT().GetTracksByPlaylist(gomock.Any(), pid).
		Return([]string{"t1"}, nil)

	cat.EXPECT().GetTracksByIDs(gomock.Any(), gomock.Any()).
		Return(&pbCatalog.GetTracksByIDsResponse{
			Tracks: []*pbCatalog.Track{{Id: "t1", Title: "Song"}},
		}, nil)

	aiMock.EXPECT().
		GeneratePlaylistMeta(gomock.Any(), gomock.Any()).
		Return("", "", ai.ErrAIAuth)

	_, err := svc.GeneratePlaylistMeta(context.Background(), pid)

	assert.Error(t, err)
}
