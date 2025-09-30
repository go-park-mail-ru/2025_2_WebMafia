package handler

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"spotify/internal/store"
	"spotify/pkg/jwtmanager"
	"testing"
	"time"
)

func TestAuthHandler(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)
	reqBody := map[string]string{
		"login":    "user_login",
		"email":    "new_user@test.com",
		"password": "some_password",
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.RegisterHandler(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
	var response struct {
		ID string `json:"id"`
	}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.ID)
	cookies := rr.Result().Cookies()
	assert.Len(t, cookies, 1)
	assert.Equal(t, "session_token", cookies[0].Name)
}

func TestRegisterHandler_ValidationErrors(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)
	testCases := []struct {
		name    string
		request map[string]string
	}{
		{
			name: "Short login",
			request: map[string]string{
				"login":    "user",
				"email":    "new_user@test.com",
				"password": "some_password",
			},
		},
		{
			name: "Invalid email",
			request: map[string]string{
				"login":    "user_login",
				"email":    "invalid_email",
				"password": "some_password",
			},
		},
		{
			name: "Short password",
			request: map[string]string{
				"login":    "user_login",
				"email":    "new_user@test.com",
				"password": "usr",
			},
		},
		{
			name: "Empty login",
			request: map[string]string{
				"login":    "",
				"email":    "new_user@test.com",
				"password": "usr",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.request)
			req := httptest.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handlers.RegisterHandler(rr, req)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
		})
	}
}

func TestRegisterHandler_DuplicateUser(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)
	reqBody := map[string]string{
		"login":    "user_login",
		"email":    "new_user@test.com",
		"password": "some_password",
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.RegisterHandler(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)

	reqBody2 := map[string]string{
		"login":    "user_login",
		"email":    "new_user@test.com",
		"password": "some_password",
	}
	body2, _ := json.Marshal(reqBody2)
	req2 := httptest.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(body2))
	req2.Header.Set("Content-Type", "application/json")
	rr2 := httptest.NewRecorder()
	handlers.RegisterHandler(rr2, req2)
	assert.Equal(t, http.StatusConflict, rr2.Code)
}

func TestLoginHandler(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)

	registerReq := map[string]string{
		"login":    "user_login",
		"email":    "new_user@test.com",
		"password": "some_password",
	}
	body, _ := json.Marshal(registerReq)
	req := httptest.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.RegisterHandler(rr, req)
	loginReq := map[string]string{
		"login":    "user_login",
		"password": "some_password",
	}
	body, _ = json.Marshal(loginReq)
	req = httptest.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	handlers.LoginHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var response struct {
		ID string `json:"id"`
	}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.ID)
	cookies := rr.Result().Cookies()
	assert.Len(t, cookies, 1)
	assert.Equal(t, "session_token", cookies[0].Name)
}

func TestLogoutHandler_Success(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)
	req := httptest.NewRequest("POST", "/api/v1/logout", nil)
	rr := httptest.NewRecorder()
	handlers.LogoutHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var response struct {
		Status string `json:"status"`
	}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "ok", response.Status)
	cookies := rr.Result().Cookies()
	assert.Len(t, cookies, 1)
	assert.Equal(t, "session_token", cookies[0].Name)
	assert.Equal(t, "", cookies[0].Value)
	assert.True(t, cookies[0].Expires.Before(time.Now()))
}

func TestLoginHandler_InvalidCredentials(t *testing.T) {
	dataStore := store.NewMemoryStore()
	jwtManager := jwtmanager.NewManager("super-secret-key", time.Hour)
	handlers := NewHandler(dataStore, jwtManager)

	reqBody := map[string]string{
		"login":    "user_login",
		"email":    "new_user@test.com",
		"password": "some_password",
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.RegisterHandler(rr, req)

	testCases := []struct {
		name           string
		request        map[string]string
		expectedStatus int
	}{
		{
			name: "user nit exist",
			request: map[string]string{
				"login":    "no_user",
				"password": "some_password",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "wrong password",
			request: map[string]string{
				"login":    "user_login",
				"password": "wrong_password",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "empty credentials",
			request: map[string]string{
				"login":    "",
				"password": "",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.request)
			req := httptest.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handlers.LoginHandler(rr, req)
			assert.Equal(t, tc.expectedStatus, rr.Code)
		})
	}
}
