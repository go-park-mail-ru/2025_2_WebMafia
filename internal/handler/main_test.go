package handler

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"spotify/internal/store"
	"spotify/pkg/jwtmanager"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

var (
	testServer *httptest.Server
	testClient *http.Client
)

func TestMain(m *testing.M) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", 720*time.Hour)
	handlers := NewHandler(dataStore, jwtManager)

	router := setupTestRouter(handlers)
	testServer = httptest.NewServer(router)
	testClient = testServer.Client()
	testClient.Timeout = 10 * time.Second

	code := m.Run()
	testServer.Close()
	os.Exit(code)
}

func setupTestRouter(h *Handlers) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/register", h.RegisterHandler).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/login", h.LoginHandler).Methods(http.MethodPost, http.MethodOptions)

	protected := api.PathPrefix("").Subrouter()
	protected.Use(h.AuthMiddleware)

	protected.HandleFunc("/logout", h.LogoutHandler).Methods(http.MethodPost, http.MethodOptions)
	protected.HandleFunc("/home", h.HomeHandler).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/tracks", h.GetAllTracksHandler).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/tracks/{id}", h.GetTrackByIDHandler).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/artists", h.GetAllArtistsHandler).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/artists/{id}", h.GetArtistByIDHandler).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/albums", h.GetAllAlbumsHandler).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/albums/{id}", h.GetAlbumByIDHandler).Methods(http.MethodGet, http.MethodOptions)

	return r
}

func makeRequest(t *testing.T, method, path string, body []byte, cookies ...*http.Cookie) (*http.Response, string) {
	req, err := http.NewRequest(method, testServer.URL+path, bytes.NewBuffer(body))
	require.NoError(t, err)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	resp, err := testClient.Do(req)
	require.NoError(t, err)
	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}
