package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllTracksHandler(t *testing.T) {
	registerReq := registerRequest{
		Login:    "tracks_user",
		Email:    "tracks@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(registerReq)
	require.NoError(t, err)

	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	resp.Body.Close()

	loginReq := loginRequest{
		Login:    "tracks_user",
		Password: "some_password",
	}
	body, err = json.Marshal(loginReq)
	require.NoError(t, err)

	resp, err = testClient.Post(testServer.URL+"/api/v1/login", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	cookies := resp.Cookies()
	require.Len(t, cookies, 1)
	sessionCookie := cookies[0]

	req, err := http.NewRequest("GET", testServer.URL+"/api/v1/tracks", nil)
	require.NoError(t, err)
	req.AddCookie(sessionCookie)

	resp, err = testClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response TracksResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Greater(t, len(response.Tracks), 0)

	for _, track := range response.Tracks {
		assert.NotEmpty(t, track.Title)
		assert.Greater(t, track.DurationMs, 0)
		assert.NotEmpty(t, track.FileURL)
	}
}

func TestGetAllArtistsHandler(t *testing.T) {
	registerReq := registerRequest{
		Login:    "artists_user",
		Email:    "artists@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(registerReq)
	require.NoError(t, err)

	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	resp.Body.Close()

	loginReq := loginRequest{
		Login:    "artists_user",
		Password: "some_password",
	}
	body, err = json.Marshal(loginReq)
	require.NoError(t, err)

	resp, err = testClient.Post(testServer.URL+"/api/v1/login", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	cookies := resp.Cookies()
	require.Len(t, cookies, 1)
	sessionCookie := cookies[0]

	req, err := http.NewRequest("GET", testServer.URL+"/api/v1/artists", nil)
	require.NoError(t, err)
	req.AddCookie(sessionCookie)

	resp, err = testClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response ArtistsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Greater(t, len(response.Artists), 0)

	for _, artist := range response.Artists {
		assert.NotEmpty(t, artist.Name)
		assert.NotZero(t, artist.ArtistID)
	}
}

func TestGetAllAlbumsHandler(t *testing.T) {
	registerReq := registerRequest{
		Login:    "albums_user",
		Email:    "albums@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(registerReq)
	require.NoError(t, err)

	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	resp.Body.Close()

	loginReq := loginRequest{
		Login:    "albums_user",
		Password: "some_password",
	}
	body, err = json.Marshal(loginReq)
	require.NoError(t, err)

	resp, err = testClient.Post(testServer.URL+"/api/v1/login", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	cookies := resp.Cookies()
	require.Len(t, cookies, 1)
	sessionCookie := cookies[0]

	req, err := http.NewRequest("GET", testServer.URL+"/api/v1/albums", nil)
	require.NoError(t, err)
	req.AddCookie(sessionCookie)

	resp, err = testClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response AlbumsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Greater(t, len(response.Albums), 0)

	for _, album := range response.Albums {
		assert.NotEmpty(t, album.Title)
		assert.NotZero(t, album.AlbumID)
		assert.NotZero(t, album.ReleaseDate)
	}
}

func TestGetTrackByIDHandler(t *testing.T) {
	registerReq := registerRequest{
		Login:    "track_by_id_user",
		Email:    "track_by_id@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(registerReq)
	require.NoError(t, err)

	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	resp.Body.Close()

	loginReq := loginRequest{
		Login:    "track_by_id_user",
		Password: "some_password",
	}
	body, err = json.Marshal(loginReq)
	require.NoError(t, err)

	resp, err = testClient.Post(testServer.URL+"/api/v1/login", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	cookies := resp.Cookies()
	require.Len(t, cookies, 1)
	sessionCookie := cookies[0]

	req, err := http.NewRequest("GET", testServer.URL+"/api/v1/tracks", nil)
	require.NoError(t, err)
	req.AddCookie(sessionCookie)

	resp, err = testClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	var tracksResponse TracksResponse
	err = json.NewDecoder(resp.Body).Decode(&tracksResponse)
	require.NoError(t, err)
	require.Greater(t, len(tracksResponse.Tracks), 0, "Should have tracks in store")

	trackID := tracksResponse.Tracks[0].TrackID

	req, err = http.NewRequest("GET", testServer.URL+"/api/v1/tracks/"+strconv.FormatUint(uint64(trackID), 10), nil)
	require.NoError(t, err)
	req.AddCookie(sessionCookie)

	resp, err = testClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response TrackResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, trackID, response.Track.TrackID)
	assert.NotEmpty(t, response.Track.Title)
}

func TestGetArtistByIDHandler(t *testing.T) {
	registerReq := registerRequest{
		Login:    "artist_by_id_user",
		Email:    "artist_by_id@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(registerReq)
	require.NoError(t, err)

	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	resp.Body.Close()

	loginReq := loginRequest{
		Login:    "artist_by_id_user",
		Password: "some_password",
	}
	body, err = json.Marshal(loginReq)
	require.NoError(t, err)

	resp, err = testClient.Post(testServer.URL+"/api/v1/login", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	cookies := resp.Cookies()
	require.Len(t, cookies, 1)
	sessionCookie := cookies[0]

	req, err := http.NewRequest("GET", testServer.URL+"/api/v1/artists", nil)
	require.NoError(t, err)
	req.AddCookie(sessionCookie)

	resp, err = testClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	var artistsResponse ArtistsResponse
	err = json.NewDecoder(resp.Body).Decode(&artistsResponse)
	require.NoError(t, err)
	require.Greater(t, len(artistsResponse.Artists), 0, "Should have artists in store")

	artistID := artistsResponse.Artists[0].ArtistID

	req, err = http.NewRequest("GET", testServer.URL+"/api/v1/artists/"+strconv.FormatUint(uint64(artistID), 10), nil)
	require.NoError(t, err)
	req.AddCookie(sessionCookie)

	resp, err = testClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetAlbumByIDHandler(t *testing.T) {
	registerReq := registerRequest{
		Login:    "album_by_id_user",
		Email:    "album_by_id@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(registerReq)
	require.NoError(t, err)

	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	resp.Body.Close()

	loginReq := loginRequest{
		Login:    "album_by_id_user",
		Password: "some_password",
	}
	body, err = json.Marshal(loginReq)
	require.NoError(t, err)

	resp, err = testClient.Post(testServer.URL+"/api/v1/login", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	cookies := resp.Cookies()
	require.Len(t, cookies, 1)
	sessionCookie := cookies[0]

	req, err := http.NewRequest("GET", testServer.URL+"/api/v1/albums", nil)
	require.NoError(t, err)
	req.AddCookie(sessionCookie)

	resp, err = testClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	var albumsResponse AlbumsResponse
	err = json.NewDecoder(resp.Body).Decode(&albumsResponse)
	require.NoError(t, err)
	require.Greater(t, len(albumsResponse.Albums), 0, "Should have albums in store")

	albumID := albumsResponse.Albums[0].AlbumID

	req, err = http.NewRequest("GET", testServer.URL+"/api/v1/albums/"+strconv.FormatUint(uint64(albumID), 10), nil)
	require.NoError(t, err)
	req.AddCookie(sessionCookie)

	resp, err = testClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetTrackByIDHandler_NotFound(t *testing.T) {
	registerReq := registerRequest{
		Login:    "track_not_found_user",
		Email:    "track_not_found@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(registerReq)
	require.NoError(t, err)

	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	resp.Body.Close()

	loginReq := loginRequest{
		Login:    "track_not_found_user",
		Password: "some_password",
	}
	body, err = json.Marshal(loginReq)
	require.NoError(t, err)

	resp, err = testClient.Post(testServer.URL+"/api/v1/login", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	cookies := resp.Cookies()
	require.Len(t, cookies, 1)
	sessionCookie := cookies[0]

	req, err := http.NewRequest("GET", testServer.URL+"/api/v1/tracks/9999999", nil)
	require.NoError(t, err)
	req.AddCookie(sessionCookie)

	resp, err = testClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGetArtistByIDHandler_NotFound(t *testing.T) {
	registerReq := registerRequest{
		Login:    "artist_not_found_user",
		Email:    "artist_not_found@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(registerReq)
	require.NoError(t, err)

	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	resp.Body.Close()

	loginReq := loginRequest{
		Login:    "artist_not_found_user",
		Password: "some_password",
	}
	body, err = json.Marshal(loginReq)
	require.NoError(t, err)

	resp, err = testClient.Post(testServer.URL+"/api/v1/login", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	cookies := resp.Cookies()
	require.Len(t, cookies, 1)
	sessionCookie := cookies[0]

	req, err := http.NewRequest("GET", testServer.URL+"/api/v1/artists/9999999", nil)
	require.NoError(t, err)
	req.AddCookie(sessionCookie)

	resp, err = testClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGetAlbumByIDHandler_NotFound(t *testing.T) {
	registerReq := registerRequest{
		Login:    "album_not_found_user",
		Email:    "album_not_found@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(registerReq)
	require.NoError(t, err)

	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	resp.Body.Close()

	loginReq := loginRequest{
		Login:    "album_not_found_user",
		Password: "some_password",
	}
	body, err = json.Marshal(loginReq)
	require.NoError(t, err)

	resp, err = testClient.Post(testServer.URL+"/api/v1/login", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	cookies := resp.Cookies()
	require.Len(t, cookies, 1)
	sessionCookie := cookies[0]

	req, err := http.NewRequest("GET", testServer.URL+"/api/v1/albums/9999999", nil)
	require.NoError(t, err)
	req.AddCookie(sessionCookie)

	resp, err = testClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
