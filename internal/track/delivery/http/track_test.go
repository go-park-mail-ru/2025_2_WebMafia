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

	mock_track "spotify/internal/mocks/track"
	"spotify/internal/track/dto"
	"spotify/internal/track/service"
)

func newTrackDTO() *dto.Track {
	return &dto.Track{
		ID:        uuid.New().String(),
		Title:     "Test Track",
		DurationS: 180,
		FileURL:   "/static/track.mp3",
		Album:     dto.Album{ID: uuid.New().String(), Title: "Test Album"},
		Artists: []dto.Artist{
			{ID: uuid.New().String(), Name: "Test Artist"},
		},
		Genres: []dto.Genre{
			{ID: uuid.New().String(), Name: "Rock"},
		},
	}
}

func TestHandler_GetTrackByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock_track.NewMockIService(ctrl)
	handler := NewHandler(mockService)
	router := mux.NewRouter()
	router.HandleFunc("/tracks/{id}", handler.GetTrackByID)

	trackDTO := newTrackDTO()
	trackID, _ := uuid.Parse(trackDTO.ID)

	t.Run("success", func(t *testing.T) {
		mockService.EXPECT().GetTrackByID(gomock.Any(), trackID).Return(trackDTO, nil)
		req := httptest.NewRequest(http.MethodGet, "/tracks/"+trackDTO.ID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var respDTO dto.Track
		err := json.Unmarshal(rr.Body.Bytes(), &respDTO)
		require.NoError(t, err)
		assert.Equal(t, *trackDTO, respDTO)
	})

	t.Run("not found", func(t *testing.T) {
		mockService.EXPECT().GetTrackByID(gomock.Any(), trackID).Return(nil, service.ErrNotFound)
		req := httptest.NewRequest(http.MethodGet, "/tracks/"+trackDTO.ID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("service internal error", func(t *testing.T) {
		mockService.EXPECT().GetTrackByID(gomock.Any(), trackID).Return(nil, errors.New("internal error"))
		req := httptest.NewRequest(http.MethodGet, "/tracks/"+trackDTO.ID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("invalid uuid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tracks/invalid-uuid", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestHandler_GetAllTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock_track.NewMockIService(ctrl)
	handler := NewHandler(mockService)

	t.Run("success", func(t *testing.T) {
		tracks := []dto.Track{*newTrackDTO()}
		mockService.EXPECT().GetAllTracks(gomock.Any(), uint64(100), uint64(0)).Return(tracks, nil)

		req := httptest.NewRequest(http.MethodGet, "/tracks", nil)
		rr := httptest.NewRecorder()
		handler.GetAllTracks(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("service error", func(t *testing.T) {
		mockService.EXPECT().GetAllTracks(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("internal error"))
		req := httptest.NewRequest(http.MethodGet, "/tracks", nil)
		rr := httptest.NewRecorder()
		handler.GetAllTracks(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandler_GetTracksByArtist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock_track.NewMockIService(ctrl)
	handler := NewHandler(mockService)
	router := mux.NewRouter()
	router.HandleFunc("/artists/{artistId}/tracks", handler.GetTracksByArtist)

	artistID := uuid.New()

	t.Run("success", func(t *testing.T) {
		tracks := []dto.Track{*newTrackDTO()}
		mockService.EXPECT().GetTracksByArtistID(gomock.Any(), artistID, uint64(100), uint64(0)).Return(tracks, nil)
		url := fmt.Sprintf("/artists/%s/tracks", artistID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid artist id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/artists/invalid-uuid/tracks", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestHandler_GetTracksByAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock_track.NewMockIService(ctrl)
	handler := NewHandler(mockService)
	router := mux.NewRouter()
	router.HandleFunc("/albums/{albumId}/tracks", handler.GetTracksByAlbum)

	albumID := uuid.New()

	t.Run("success", func(t *testing.T) {
		tracks := []dto.Track{*newTrackDTO()}
		mockService.EXPECT().GetTracksByAlbumID(gomock.Any(), albumID, uint64(100), uint64(0)).Return(tracks, nil)
		url := fmt.Sprintf("/albums/%s/tracks", albumID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid album id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/albums/invalid-uuid/tracks", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestHandler_GetTracksByGenre(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock_track.NewMockIService(ctrl)
	handler := NewHandler(mockService)
	router := mux.NewRouter()
	router.HandleFunc("/genres/{genreId}/tracks", handler.GetTracksByGenre)

	genreID := uuid.New()

	t.Run("success", func(t *testing.T) {
		tracks := []dto.Track{*newTrackDTO()}
		mockService.EXPECT().GetTracksByGenreID(gomock.Any(), genreID, uint64(100), uint64(0)).Return(tracks, nil)
		url := fmt.Sprintf("/genres/%s/tracks", genreID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid genre id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/genres/invalid-uuid/tracks", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestHandler_RegisterPlay(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock_track.NewMockIService(ctrl)
	handler := NewHandler(mockService)
	router := mux.NewRouter()
	router.HandleFunc("/tracks/{id}/listen", handler.RegisterPlay)

	trackID := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockService.EXPECT().RegisterPlay(gomock.Any(), trackID).Return(nil)
		url := fmt.Sprintf("/tracks/%s/listen", trackID)
		req := httptest.NewRequest(http.MethodPost, url, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusAccepted, rr.Code)
	})

	t.Run("not found", func(t *testing.T) {
		mockService.EXPECT().RegisterPlay(gomock.Any(), trackID).Return(service.ErrNotFound)
		url := fmt.Sprintf("/tracks/%s/listen", trackID)
		req := httptest.NewRequest(http.MethodPost, url, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		url := "/tracks/invalid-uuid/listen"
		req := httptest.NewRequest(http.MethodPost, url, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
