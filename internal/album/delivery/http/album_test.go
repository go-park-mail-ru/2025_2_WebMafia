package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"spotify/internal/album/dto"
	"spotify/internal/album/service"
	mock_album "spotify/internal/mocks/album"
)

func newAlbumDTO() *dto.Album {
	return &dto.Album{
		ID:          uuid.New().String(),
		Title:       "Test Album",
		ReleaseDate: "2023-01-01",
		Artists: []dto.Artist{
			{ID: uuid.New().String(), Name: "Test Artist"},
		},
	}
}

func TestHandler_GetAlbumByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock_album.NewMockIService(ctrl)
	handler := NewHandler(mockService)
	router := mux.NewRouter()
	router.HandleFunc("/albums/{id}", handler.GetAlbumByID)

	albumDTO := newAlbumDTO()
	albumID, _ := uuid.Parse(albumDTO.ID)

	t.Run("success", func(t *testing.T) {
		mockService.EXPECT().GetAlbumByID(gomock.Any(), albumID).Return(albumDTO, nil)
		req := httptest.NewRequest(http.MethodGet, "/albums/"+albumDTO.ID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var respDTO dto.Album
		err := json.Unmarshal(rr.Body.Bytes(), &respDTO)
		require.NoError(t, err)
		assert.Equal(t, *albumDTO, respDTO)
	})

	t.Run("not found", func(t *testing.T) {
		mockService.EXPECT().GetAlbumByID(gomock.Any(), albumID).Return(nil, service.ErrNotFound)
		req := httptest.NewRequest(http.MethodGet, "/albums/"+albumDTO.ID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/albums/invalid-uuid", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestHandler_GetAllAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock_album.NewMockIService(ctrl)
	handler := NewHandler(mockService)

	t.Run("success", func(t *testing.T) {
		albums := []dto.Album{*newAlbumDTO()}
		mockService.EXPECT().GetAllAlbums(gomock.Any(), uint64(100), uint64(0)).Return(albums, nil)
		req := httptest.NewRequest(http.MethodGet, "/albums", nil)
		rr := httptest.NewRecorder()
		handler.GetAllAlbums(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("service error", func(t *testing.T) {
		mockService.EXPECT().GetAllAlbums(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("internal error"))
		req := httptest.NewRequest(http.MethodGet, "/albums", nil)
		rr := httptest.NewRecorder()
		handler.GetAllAlbums(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandler_GetAlbumsByArtistID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock_album.NewMockIService(ctrl)
	handler := NewHandler(mockService)
	router := mux.NewRouter()
	router.HandleFunc("/artists/{artistId}/albums", handler.GetAlbumsByArtistID)

	artistID := uuid.New()

	t.Run("success", func(t *testing.T) {
		albums := []dto.Album{*newAlbumDTO()}
		mockService.EXPECT().GetAlbumsByArtistID(gomock.Any(), artistID, uint64(100), uint64(0)).Return(albums, nil)
		url := fmt.Sprintf("/artists/%s/albums", artistID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid artist id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/artists/invalid-uuid/albums", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("service error", func(t *testing.T) {
		mockService.EXPECT().GetAlbumsByArtistID(gomock.Any(), artistID, gomock.Any(), gomock.Any()).Return(nil, errors.New("internal error"))
		url := fmt.Sprintf("/artists/%s/albums", artistID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
