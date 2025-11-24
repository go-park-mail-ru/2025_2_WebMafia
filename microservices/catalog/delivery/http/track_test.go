package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"spotify/microservices/catalog/dto"
	service_mock "spotify/microservices/catalog/mocks/service"
	"spotify/microservices/catalog/service"
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

		req := httptest.NewRequest(http.MethodGet, "/tracks/"+id.String(), nil)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		rr := httptest.NewRecorder()

		handler.GetTrackByID(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("not found", func(t *testing.T) {
		mockSvc.EXPECT().GetTrackByID(gomock.Any(), id).Return(nil, service.ErrNotFound)

		req := httptest.NewRequest(http.MethodGet, "/tracks/"+id.String(), nil)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		rr := httptest.NewRecorder()

		handler.GetTrackByID(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tracks/invalid", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "invalid"})
		rr := httptest.NewRecorder()

		handler.GetTrackByID(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("missing id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tracks/", nil)
		rr := httptest.NewRecorder()

		handler.GetTrackByID(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestHandler_GetAllTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetAllTracks(gomock.Any(), uint64(100), uint64(0)).Return([]dto.Track{}, nil)
		req := httptest.NewRequest(http.MethodGet, "/tracks", nil)
		rr := httptest.NewRecorder()
		handler.GetAllTracks(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc.EXPECT().GetAllTracks(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("err"))
		req := httptest.NewRequest(http.MethodGet, "/tracks", nil)
		rr := httptest.NewRecorder()
		handler.GetAllTracks(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandler_GetTracksByArtist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetTracksByArtistID(gomock.Any(), id, gomock.Any(), gomock.Any()).
			Return([]dto.Track{}, nil)

		req := httptest.NewRequest(http.MethodGet, "/artists/"+id.String()+"/tracks", nil)
		req = mux.SetURLVars(req, map[string]string{"artistId": id.String()})
		rr := httptest.NewRecorder()

		handler.GetTracksByArtist(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/artists/invalid/tracks", nil)
		req = mux.SetURLVars(req, map[string]string{"artistId": "invalid"})
		rr := httptest.NewRecorder()

		handler.GetTracksByArtist(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc.EXPECT().GetTracksByArtistID(gomock.Any(), id, gomock.Any(), gomock.Any()).
			Return(nil, errors.New("err"))

		req := httptest.NewRequest(http.MethodGet, "/artists/"+id.String()+"/tracks", nil)
		req = mux.SetURLVars(req, map[string]string{"artistId": id.String()})
		rr := httptest.NewRecorder()

		handler.GetTracksByArtist(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandler_GetTracksByAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetTracksByAlbumID(gomock.Any(), id, gomock.Any(), gomock.Any()).
			Return([]dto.Track{}, nil)

		req := httptest.NewRequest(http.MethodGet, "/albums/"+id.String()+"/tracks", nil)
		req = mux.SetURLVars(req, map[string]string{"albumId": id.String()})
		rr := httptest.NewRecorder()

		handler.GetTracksByAlbum(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/albums/invalid/tracks", nil)
		req = mux.SetURLVars(req, map[string]string{"albumId": "invalid"})
		rr := httptest.NewRecorder()

		handler.GetTracksByAlbum(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc.EXPECT().GetTracksByAlbumID(gomock.Any(), id, gomock.Any(), gomock.Any()).
			Return(nil, errors.New("err"))

		req := httptest.NewRequest(http.MethodGet, "/albums/"+id.String()+"/tracks", nil)
		req = mux.SetURLVars(req, map[string]string{"albumId": id.String()})
		rr := httptest.NewRecorder()

		handler.GetTracksByAlbum(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandler_GetTracksByGenre(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetTracksByGenreID(gomock.Any(), id, gomock.Any(), gomock.Any()).
			Return([]dto.Track{}, nil)

		req := httptest.NewRequest(http.MethodGet, "/genres/"+id.String()+"/tracks", nil)
		req = mux.SetURLVars(req, map[string]string{"genreId": id.String()})
		rr := httptest.NewRecorder()

		handler.GetTracksByGenre(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/genres/invalid/tracks", nil)
		req = mux.SetURLVars(req, map[string]string{"genreId": "invalid"})
		rr := httptest.NewRecorder()

		handler.GetTracksByGenre(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc.EXPECT().GetTracksByGenreID(gomock.Any(), id, gomock.Any(), gomock.Any()).
			Return(nil, errors.New("err"))

		req := httptest.NewRequest(http.MethodGet, "/genres/"+id.String()+"/tracks", nil)
		req = mux.SetURLVars(req, map[string]string{"genreId": id.String()})
		rr := httptest.NewRecorder()

		handler.GetTracksByGenre(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
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

		req := httptest.NewRequest(http.MethodPost, "/tracks/"+id.String()+"/listen", nil)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		rr := httptest.NewRecorder()

		handler.RegisterPlay(rr, req)

		assert.Equal(t, http.StatusAccepted, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/tracks/invalid/listen", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "invalid"})
		rr := httptest.NewRecorder()

		handler.RegisterPlay(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc.EXPECT().RegisterPlay(gomock.Any(), id).Return(errors.New("err"))

		req := httptest.NewRequest(http.MethodPost, "/tracks/"+id.String()+"/listen", nil)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		rr := httptest.NewRecorder()

		handler.RegisterPlay(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandler_SearchTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().SearchTracks(gomock.Any(), "query", uint64(50)).
			Return([]dto.TrackSearch{}, nil)

		req := httptest.NewRequest(http.MethodGet, "/tracks/search?q=query", nil)
		rr := httptest.NewRecorder()

		handler.SearchTracks(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("empty query", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tracks/search?q=", nil)
		rr := httptest.NewRecorder()

		handler.SearchTracks(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc.EXPECT().SearchTracks(gomock.Any(), "query", uint64(50)).
			Return(nil, errors.New("err"))

		req := httptest.NewRequest(http.MethodGet, "/tracks/search?q=query", nil)
		rr := httptest.NewRecorder()

		handler.SearchTracks(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
