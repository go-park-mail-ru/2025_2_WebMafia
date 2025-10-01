package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"spotify/internal/store"
	"spotify/pkg/jwtmanager"
	"strconv"
	"testing"
	"time"
)

func TestGetAllTracksHandler(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)

	req := httptest.NewRequest("GET", "/api/v1/tracks", nil)
	rr := httptest.NewRecorder()
	handlers.GetAllTracksHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var response TracksResponse
	err := json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Greater(t, len(response.Tracks), 0)

	for _, track := range response.Tracks {
		assert.NotEmpty(t, track.Title)
		assert.Greater(t, track.DurationMs, 0)
		assert.NotEmpty(t, track.FileURL)
	}
}

func TestGetAllArtistsHandler(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)

	req := httptest.NewRequest("GET", "/api/v1/artists", nil)
	rr := httptest.NewRecorder()
	handlers.GetAllArtistsHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var response ArtistsResponse
	err := json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Greater(t, len(response.Artists), 0)

	for _, artist := range response.Artists {
		assert.NotEmpty(t, artist.Name)
		assert.NotZero(t, artist.ArtistID)
	}
}

func TestGetAllAlbumsHandler(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)

	req := httptest.NewRequest("GET", "/api/v1/albums", nil)
	rr := httptest.NewRecorder()
	handlers.GetAllAlbumsHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var response AlbumsResponse
	err := json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Greater(t, len(response.Albums), 0)

	for _, album := range response.Albums {
		assert.NotEmpty(t, album.Title)
		assert.NotZero(t, album.AlbumID)
		assert.NotZero(t, album.ReleaseDate)
	}
}

func TestGetTrackByIDHandler(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)

	tracks, err := dataStore.GetAllTracks()
	require.NoError(t, err)
	assert.Greater(t, len(tracks), 0, "Should have tracks in store")

	trackID := tracks[0].TrackID
	req := httptest.NewRequest("GET", "/api/v1/tracks/"+strconv.FormatUint(uint64(trackID), 10), nil)
	req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(uint64(trackID), 10)})
	rr := httptest.NewRecorder()

	handlers.GetTrackByIDHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var response TrackResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, trackID, response.Track.TrackID)
	assert.NotEmpty(t, response.Track.Title)
}

func TestGetArtistByIDHandler(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)

	artists, err := dataStore.GetAllArtists()
	require.NoError(t, err)
	assert.Greater(t, len(artists), 0, "Should have artists in store")

	artistID := artists[0].ArtistID
	req := httptest.NewRequest("GET", "/api/v1/artists/"+strconv.FormatUint(uint64(artistID), 10), nil)
	req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(uint64(artistID), 10)})
	rr := httptest.NewRecorder()

	handlers.GetArtistByIDHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetAlbumByIDHandler(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)

	albums, err := dataStore.GetAllAlbums()
	require.NoError(t, err)
	assert.Greater(t, len(albums), 0, "Should have albums in store")

	albumID := albums[0].AlbumID
	req := httptest.NewRequest("GET", "/api/v1/albums/"+strconv.FormatUint(uint64(albumID), 10), nil)
	req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(uint64(albumID), 10)})
	rr := httptest.NewRecorder()

	handlers.GetAlbumByIDHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetTrackByIDHandler_NotFound(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)

	req := httptest.NewRequest("GET", "/api/v1/tracks/9999999", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "9999999"})
	rr := httptest.NewRecorder()

	handlers.GetTrackByIDHandler(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestGetArtistByIDHandler_NotFound(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)

	req := httptest.NewRequest("GET", "/api/v1/artists/9999999", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "9999999"})
	rr := httptest.NewRecorder()

	handlers.GetArtistByIDHandler(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestGetAlbumByIDHandler_NotFound(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)

	req := httptest.NewRequest("GET", "/api/v1/albums/9999999", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "9999999"})
	rr := httptest.NewRecorder()

	handlers.GetAlbumByIDHandler(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}
