package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"spotify/pkg/ws"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"spotify/microservices/catalog/dto"
	"spotify/microservices/catalog/service"
	service_mock "spotify/mocks/catalog/service/http"
)

func TestHandler_GetAlbumByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil, ws.Config{}, nil)

	id := uuid.New()
	dtoAlbum := &dto.Album{ID: id.String(), Title: "Album"}

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetAlbumByID(gomock.Any(), id).Return(dtoAlbum, nil)

		req := httptest.NewRequest(http.MethodGet, "/albums/"+id.String(), nil)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		rr := httptest.NewRecorder()

		handler.GetAlbumByID(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("not found", func(t *testing.T) {
		mockSvc.EXPECT().GetAlbumByID(gomock.Any(), id).Return(nil, service.ErrNotFound)

		req := httptest.NewRequest(http.MethodGet, "/albums/"+id.String(), nil)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		rr := httptest.NewRecorder()

		handler.GetAlbumByID(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/albums/invalid", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "invalid"})
		rr := httptest.NewRecorder()

		handler.GetAlbumByID(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("missing id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/albums/", nil)
		rr := httptest.NewRecorder()

		handler.GetAlbumByID(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestHandler_GetAllAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil, ws.Config{}, nil)

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetAllAlbums(gomock.Any(), uint64(100), uint64(0)).Return([]dto.Album{}, nil)
		req := httptest.NewRequest(http.MethodGet, "/albums", nil)
		rr := httptest.NewRecorder()
		handler.GetAllAlbums(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc.EXPECT().GetAllAlbums(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("err"))
		req := httptest.NewRequest(http.MethodGet, "/albums", nil)
		rr := httptest.NewRecorder()
		handler.GetAllAlbums(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandler_GetAlbumsByArtistID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil, ws.Config{}, nil)

	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetAlbumsByArtistID(gomock.Any(), id, uint64(100), uint64(0)).
			Return([]dto.Album{}, nil)

		req := httptest.NewRequest(http.MethodGet, "/artists/"+id.String()+"/albums", nil)
		req = mux.SetURLVars(req, map[string]string{"artistId": id.String()})
		rr := httptest.NewRecorder()

		handler.GetAlbumsByArtistID(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/artists/invalid/albums", nil)
		req = mux.SetURLVars(req, map[string]string{"artistId": "invalid"})
		rr := httptest.NewRecorder()

		handler.GetAlbumsByArtistID(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("missing id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/artists//albums", nil)
		rr := httptest.NewRecorder()

		handler.GetAlbumsByArtistID(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc.EXPECT().GetAlbumsByArtistID(gomock.Any(), id, gomock.Any(), gomock.Any()).
			Return(nil, errors.New("err"))

		req := httptest.NewRequest(http.MethodGet, "/artists/"+id.String()+"/albums", nil)
		req = mux.SetURLVars(req, map[string]string{"artistId": id.String()})
		rr := httptest.NewRecorder()

		handler.GetAlbumsByArtistID(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandler_SearchAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil, ws.Config{}, nil)

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().SearchAlbums(gomock.Any(), "query", uint64(50)).
			Return([]dto.AlbumSearch{}, nil)

		req := httptest.NewRequest(http.MethodGet, "/albums/search?q=query", nil)
		rr := httptest.NewRecorder()

		handler.SearchAlbums(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("empty query", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/albums/search?q=", nil)
		rr := httptest.NewRecorder()

		handler.SearchAlbums(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc.EXPECT().SearchAlbums(gomock.Any(), "query", uint64(50)).
			Return(nil, errors.New("err"))

		req := httptest.NewRequest(http.MethodGet, "/albums/search?q=query", nil)
		rr := httptest.NewRecorder()

		handler.SearchAlbums(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
