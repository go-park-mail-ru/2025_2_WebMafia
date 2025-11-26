package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"spotify/microservices/catalog/dto"
	service_mock "spotify/microservices/catalog/mocks/grpc_service"
	pb "spotify/proto/catalog"
)

func TestHandler_GetTrackByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)

	id := uuid.New()
	dtoTrack := &dto.Track{ID: id.String(), Title: "Track"}

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetTrackByID(gomock.Any(), id).Return(dtoTrack, nil)

		resp, err := handler.GetTrackByID(context.Background(), &pb.GetTrackByIDRequest{Id: id.String()})
		assert.NoError(t, err)
		assert.Equal(t, dtoTrack.Title, resp.Title)
	})

	t.Run("invalid id", func(t *testing.T) {
		_, err := handler.GetTrackByID(context.Background(), &pb.GetTrackByIDRequest{Id: "invalid"})
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("not found", func(t *testing.T) {
		mockSvc.EXPECT().GetTrackByID(gomock.Any(), id).Return(nil, errors.New("not found"))
		_, err := handler.GetTrackByID(context.Background(), &pb.GetTrackByIDRequest{Id: id.String()})
		assert.Error(t, err)
		assert.Equal(t, codes.NotFound, status.Code(err))
	})
}

func TestHandler_GetAllTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetAllTracks(gomock.Any(), uint64(10), uint64(0)).
			Return([]dto.Track{{Title: "T1"}}, nil)

		resp, err := handler.GetAllTracks(context.Background(), &pb.GetAllTracksRequest{Limit: 10, Offset: 0})
		assert.NoError(t, err)
		assert.Len(t, resp.Tracks, 1)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc.EXPECT().GetAllTracks(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, errors.New("err"))
		_, err := handler.GetAllTracks(context.Background(), &pb.GetAllTracksRequest{})
		assert.Error(t, err)
		assert.Equal(t, codes.Internal, status.Code(err))
	})
}

func TestHandler_GetArtistByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetArtistByID(gomock.Any(), id).Return(&dto.Artist{Name: "A1"}, nil)
		resp, err := handler.GetArtistByID(context.Background(), &pb.GetArtistByIDRequest{Id: id.String()})
		assert.NoError(t, err)
		assert.Equal(t, "A1", resp.Name)
	})

	t.Run("invalid id", func(t *testing.T) {
		_, err := handler.GetArtistByID(context.Background(), &pb.GetArtistByIDRequest{Id: "bad"})
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
	})
}

func TestHandler_GetAllArtists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetAllArtists(gomock.Any(), gomock.Any(), gomock.Any()).Return([]dto.Artist{}, nil)
		_, err := handler.GetAllArtists(context.Background(), &pb.GetAllArtistsRequest{})
		assert.NoError(t, err)
	})
}

func TestHandler_GetAlbumByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetAlbumByID(gomock.Any(), id).Return(&dto.Album{Title: "Al1"}, nil)
		resp, err := handler.GetAlbumByID(context.Background(), &pb.GetAlbumByIDRequest{Id: id.String()})
		assert.NoError(t, err)
		assert.Equal(t, "Al1", resp.Title)
	})

	t.Run("invalid id", func(t *testing.T) {
		_, err := handler.GetAlbumByID(context.Background(), &pb.GetAlbumByIDRequest{Id: "x"})
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
	})
}

func TestHandler_GetAllAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetAllAlbums(gomock.Any(), gomock.Any(), gomock.Any()).Return([]dto.Album{}, nil)
		_, err := handler.GetAllAlbums(context.Background(), &pb.GetAllAlbumsRequest{})
		assert.NoError(t, err)
	})
}

func TestHandler_RegisterPlay(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().RegisterPlay(gomock.Any(), id).Return(nil)
		_, err := handler.RegisterPlay(context.Background(), &pb.RegisterPlayRequest{TrackId: id.String()})
		assert.NoError(t, err)
	})

	t.Run("invalid id", func(t *testing.T) {
		_, err := handler.RegisterPlay(context.Background(), &pb.RegisterPlayRequest{TrackId: "x"})
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
	})
}

func TestHandler_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)

	t.Run("SearchTracks", func(t *testing.T) {
		mockSvc.EXPECT().SearchTracks(gomock.Any(), "q", uint64(10)).Return([]dto.TrackSearch{}, nil)
		_, err := handler.SearchTracks(context.Background(), &pb.SearchTracksRequest{Query: "q", Limit: 10})
		assert.NoError(t, err)
	})

	t.Run("SearchArtists", func(t *testing.T) {
		mockSvc.EXPECT().SearchArtists(gomock.Any(), "q", uint64(10)).Return([]dto.ArtistSearch{}, nil)
		_, err := handler.SearchArtists(context.Background(), &pb.SearchArtistsRequest{Query: "q", Limit: 10})
		assert.NoError(t, err)
	})

	t.Run("SearchAlbums", func(t *testing.T) {
		mockSvc.EXPECT().SearchAlbums(gomock.Any(), "q", uint64(10)).Return([]dto.AlbumSearch{}, nil)
		_, err := handler.SearchAlbums(context.Background(), &pb.SearchAlbumsRequest{Query: "q", Limit: 10})
		assert.NoError(t, err)
	})
}

func TestHandler_GetTracksBy(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)
	id := uuid.New()

	t.Run("ByArtist success", func(t *testing.T) {
		mockSvc.EXPECT().GetTracksByArtistID(gomock.Any(), id, gomock.Any(), gomock.Any()).
			Return([]dto.Track{}, nil)
		_, err := handler.GetTracksByArtistID(context.Background(), &pb.GetTracksByArtistIDRequest{ArtistId: id.String()})
		assert.NoError(t, err)
	})

	t.Run("ByAlbum success", func(t *testing.T) {
		mockSvc.EXPECT().GetTracksByAlbumID(gomock.Any(), id, gomock.Any(), gomock.Any()).
			Return([]dto.Track{}, nil)
		_, err := handler.GetTracksByAlbumID(context.Background(), &pb.GetTracksByAlbumIDRequest{AlbumId: id.String()})
		assert.NoError(t, err)
	})

	t.Run("ByGenre success", func(t *testing.T) {
		mockSvc.EXPECT().GetTracksByGenreID(gomock.Any(), id, gomock.Any(), gomock.Any()).
			Return([]dto.Track{}, nil)
		_, err := handler.GetTracksByGenreID(context.Background(), &pb.GetTracksByGenreIDRequest{GenreId: id.String()})
		assert.NoError(t, err)
	})

	t.Run("ByArtist invalid ID", func(t *testing.T) {
		_, err := handler.GetTracksByArtistID(context.Background(), &pb.GetTracksByArtistIDRequest{ArtistId: "bad"})
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("ByAlbum error", func(t *testing.T) {
		mockSvc.EXPECT().GetTracksByAlbumID(gomock.Any(), id, gomock.Any(), gomock.Any()).
			Return(nil, errors.New("err"))
		_, err := handler.GetTracksByAlbumID(context.Background(), &pb.GetTracksByAlbumIDRequest{AlbumId: id.String()})
		assert.Error(t, err)
		assert.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("ByGenre invalid ID", func(t *testing.T) {
		_, err := handler.GetTracksByGenreID(context.Background(), &pb.GetTracksByGenreIDRequest{GenreId: "bad"})
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("ByGenre error", func(t *testing.T) {
		mockSvc.EXPECT().GetTracksByGenreID(gomock.Any(), id, gomock.Any(), gomock.Any()).
			Return(nil, errors.New("err"))
		_, err := handler.GetTracksByGenreID(context.Background(), &pb.GetTracksByGenreIDRequest{GenreId: id.String()})
		assert.Error(t, err)
		assert.Equal(t, codes.Internal, status.Code(err))
	})
}

func TestHandler_GetArtistsByIDs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetArtistsByIDs(gomock.Any(), []uuid.UUID{id}).Return([]dto.Artist{}, nil)
		_, err := handler.GetArtistsByIDs(context.Background(), &pb.GetArtistsByIDsRequest{Ids: []string{id.String()}})
		assert.NoError(t, err)
	})

	t.Run("empty", func(t *testing.T) {
		resp, err := handler.GetArtistsByIDs(context.Background(), &pb.GetArtistsByIDsRequest{Ids: []string{}})
		assert.NoError(t, err)
		assert.Empty(t, resp.Artists)
	})

	t.Run("invalid id", func(t *testing.T) {
		_, err := handler.GetArtistsByIDs(context.Background(), &pb.GetArtistsByIDsRequest{Ids: []string{"bad"}})
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("error", func(t *testing.T) {
		mockSvc.EXPECT().GetArtistsByIDs(gomock.Any(), []uuid.UUID{id}).Return(nil, errors.New("err"))
		_, err := handler.GetArtistsByIDs(context.Background(), &pb.GetArtistsByIDsRequest{Ids: []string{id.String()}})
		assert.Error(t, err)
		assert.Equal(t, codes.Internal, status.Code(err))
	})
}

func TestHandler_GetAlbumsByIDs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetAlbumsByIDs(gomock.Any(), []uuid.UUID{id}).Return([]dto.Album{}, nil)
		_, err := handler.GetAlbumsByIDs(context.Background(), &pb.GetAlbumsByIDsRequest{Ids: []string{id.String()}})
		assert.NoError(t, err)
	})

	t.Run("empty", func(t *testing.T) {
		resp, err := handler.GetAlbumsByIDs(context.Background(), &pb.GetAlbumsByIDsRequest{Ids: []string{}})
		assert.NoError(t, err)
		assert.Empty(t, resp.Albums)
	})

	t.Run("invalid id", func(t *testing.T) {
		_, err := handler.GetAlbumsByIDs(context.Background(), &pb.GetAlbumsByIDsRequest{Ids: []string{"bad"}})
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("error", func(t *testing.T) {
		mockSvc.EXPECT().GetAlbumsByIDs(gomock.Any(), []uuid.UUID{id}).Return(nil, errors.New("err"))
		_, err := handler.GetAlbumsByIDs(context.Background(), &pb.GetAlbumsByIDsRequest{Ids: []string{id.String()}})
		assert.Error(t, err)
		assert.Equal(t, codes.Internal, status.Code(err))
	})
}
