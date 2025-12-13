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

func TestHandler_GetArtistByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil, ws.Config{}, nil)

	artistID := uuid.New()
	dtoArtist := &dto.Artist{ID: artistID.String(), Name: "Artist"}

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetArtistByID(gomock.Any(), artistID).Return(dtoArtist, nil)

		req := httptest.NewRequest(http.MethodGet, "/artists/"+artistID.String(), nil)
		req = mux.SetURLVars(req, map[string]string{"id": artistID.String()})
		rr := httptest.NewRecorder()

		handler.GetArtistByID(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "Artist")
	})

	t.Run("not found", func(t *testing.T) {
		mockSvc.EXPECT().GetArtistByID(gomock.Any(), artistID).Return(nil, service.ErrNotFound)

		req := httptest.NewRequest(http.MethodGet, "/artists/"+artistID.String(), nil)
		req = mux.SetURLVars(req, map[string]string{"id": artistID.String()})
		rr := httptest.NewRecorder()

		handler.GetArtistByID(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("internal error", func(t *testing.T) {
		mockSvc.EXPECT().GetArtistByID(gomock.Any(), artistID).Return(nil, errors.New("db error"))

		req := httptest.NewRequest(http.MethodGet, "/artists/"+artistID.String(), nil)
		req = mux.SetURLVars(req, map[string]string{"id": artistID.String()})
		rr := httptest.NewRecorder()

		handler.GetArtistByID(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("invalid uuid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/artists/invalid", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "invalid"})
		rr := httptest.NewRecorder()

		handler.GetArtistByID(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("missing id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/artists/", nil)
		rr := httptest.NewRecorder()

		handler.GetArtistByID(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestHandler_GetAllArtists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil, ws.Config{}, nil)

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetAllArtists(gomock.Any(), uint64(100), uint64(0)).
			Return([]dto.Artist{{Name: "A1"}}, nil)

		req := httptest.NewRequest(http.MethodGet, "/artists", nil)
		rr := httptest.NewRecorder()

		handler.GetAllArtists(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("internal error", func(t *testing.T) {
		mockSvc.EXPECT().GetAllArtists(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, errors.New("err"))

		req := httptest.NewRequest(http.MethodGet, "/artists", nil)
		rr := httptest.NewRecorder()

		handler.GetAllArtists(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandler_SearchArtists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil, ws.Config{}, nil)

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().SearchArtists(gomock.Any(), "query", uint64(50)).
			Return([]dto.ArtistSearch{}, nil)

		req := httptest.NewRequest(http.MethodGet, "/artists/search?q=query", nil)
		rr := httptest.NewRecorder()

		handler.SearchArtists(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("empty query", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/artists/search?q=", nil)
		rr := httptest.NewRecorder()

		handler.SearchArtists(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("internal error", func(t *testing.T) {
		mockSvc.EXPECT().SearchArtists(gomock.Any(), "query", uint64(50)).
			Return(nil, errors.New("err"))

		req := httptest.NewRequest(http.MethodGet, "/artists/search?q=query", nil)
		rr := httptest.NewRecorder()

		handler.SearchArtists(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
