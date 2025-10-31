package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"spotify/internal/artist/dto"
	"spotify/internal/artist/service"
	mock_artist "spotify/internal/mocks/artist"
)

func TestArtistHandler_GetArtistByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_artist.NewMockIService(ctrl)
	handler := NewHandler(mockService)
	router := mux.NewRouter()
	router.HandleFunc("/artists/{id}", handler.GetArtistByID)

	artistDTO := &dto.Artist{
		ID:        uuid.New().String(),
		Name:      "Test Artist",
		AvatarURL: "http://example.com/avatar.jpg",
	}

	t.Run("success", func(t *testing.T) {
		id, _ := uuid.Parse(artistDTO.ID)
		mockService.EXPECT().GetArtistByID(gomock.Any(), id).Return(artistDTO, nil)

		url := fmt.Sprintf("/artists/%s", artistDTO.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var respDTO dto.Artist
		err := json.Unmarshal(rr.Body.Bytes(), &respDTO)
		require.NoError(t, err)
		assert.Equal(t, *artistDTO, respDTO)
	})

	t.Run("not found", func(t *testing.T) {
		id, _ := uuid.Parse(artistDTO.ID)
		mockService.EXPECT().GetArtistByID(gomock.Any(), id).Return(nil, service.ErrNotFound)

		url := fmt.Sprintf("/artists/%s", artistDTO.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		url := "/artists/invalid-uuid"
		req := httptest.NewRequest(http.MethodGet, url, nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestArtistHandler_GetAllArtists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_artist.NewMockIService(ctrl)
	handler := NewHandler(mockService)

	artistDTO := dto.Artist{
		ID:   uuid.New().String(),
		Name: "Test Artist",
	}

	t.Run("success", func(t *testing.T) {
		artists := []dto.Artist{artistDTO}
		mockService.EXPECT().GetAllArtists(gomock.Any(), uint64(100), uint64(0)).Return(artists, nil)

		req := httptest.NewRequest(http.MethodGet, "/artists", nil)
		rr := httptest.NewRecorder()

		handler.GetAllArtists(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var respDTOs []dto.Artist
		err := json.Unmarshal(rr.Body.Bytes(), &respDTOs)
		require.NoError(t, err)
		assert.Equal(t, artists, respDTOs)
	})

	t.Run("success with pagination", func(t *testing.T) {
		artists := []dto.Artist{artistDTO}
		mockService.EXPECT().GetAllArtists(gomock.Any(), uint64(50), uint64(10)).Return(artists, nil)

		req := httptest.NewRequest(http.MethodGet, "/artists?limit=50&offset=10", nil)
		rr := httptest.NewRecorder()

		handler.GetAllArtists(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}
