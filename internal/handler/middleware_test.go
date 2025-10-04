package handler

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
	handlers, _, _ := initTestEnv(t)

	registerReq := registerRequest{
		Login:    "user_login",
		Email:    "new_user@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(registerReq)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.RegisterHandler(rr, req)

	cookies := rr.Result().Cookies()
	var authCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "session_token" {
			authCookie = cookie
			break
		}
	}
	assert.NotNil(t, authCookie)

	req2 := httptest.NewRequest("GET", "/api/v1/protected", nil)
	req2.AddCookie(authCookie)
	rr2 := httptest.NewRecorder()

	var userIDInContext interface{}
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIDInContext = r.Context().Value(userIDKey)
		w.WriteHeader(http.StatusOK)
	})

	middleware := handlers.AuthMiddleware(testHandler)
	middleware.ServeHTTP(rr2, req2)

	assert.Equal(t, http.StatusOK, rr2.Code)
	assert.NotNil(t, userIDInContext)
}

func TestAuthMiddleware_NoToken(t *testing.T) {
	handlers, _, _ := initTestEnv(t)

	req := httptest.NewRequest("GET", "/api/v1/protected", nil)
	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	middleware := handlers.AuthMiddleware(testHandler)
	middleware.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	handlers, _, _ := initTestEnv(t)

	req := httptest.NewRequest("GET", "/api/v1/protected", nil)
	req.AddCookie(&http.Cookie{
		Name:  "session_token",
		Value: "invalid-token",
	})
	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	middleware := handlers.AuthMiddleware(testHandler)
	middleware.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}
