package router

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"spotify/internal/handler"
	"spotify/internal/store"
	"spotify/pkg/jwtmanager"
	"testing"
	"time"
)

func TestRouter_PublicRoutes(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("test-secret", time.Hour)
	handlers := handler.NewHandler(dataStore, jwtManager)
	corsConfig := handler.CORSConfig{
		AllowedOrigins:   []string{"http://localhost:8090"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
	}
	router := NewRouter(handlers, corsConfig)
	ts := httptest.NewServer(router)
	defer ts.Close()
	client := &http.Client{}
	req, _ := http.NewRequest("OPTIONS", ts.URL+"/api/v1/register", nil)
	resp, err := client.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	req, _ = http.NewRequest("GET", ts.URL+"/api/v1/tracks", nil)
	resp, err = client.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.NotEqual(t, http.StatusNotFound, resp.StatusCode)
}

func TestRouter_ProtectedRoutes(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("test-secret", time.Hour)
	handlers := handler.NewHandler(dataStore, jwtManager)
	corsConfig := handler.CORSConfig{
		AllowedOrigins:   []string{"http://localhost:8090"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
	}
	router := NewRouter(handlers, corsConfig)
	ts := httptest.NewServer(router)
	defer ts.Close()
	client := &http.Client{}
	req, _ := http.NewRequest("GET", ts.URL+"/api/v1/home", nil)
	resp, err := client.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	req, _ = http.NewRequest("GET", ts.URL+"/api/v1/tracks", nil)
	req.AddCookie(&http.Cookie{Name: "session_token", Value: "invalid-token"})
	resp, err = client.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
