package handler

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAuthHandler(t *testing.T) {
	handlers, _, _ := initTestEnv(t)

	reqBody := registerRequest{
		Login:    "user_login",
		Email:    "new_user@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.RegisterHandler(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)

	var response registerResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.ID)

	cookies := rr.Result().Cookies()
	assert.Len(t, cookies, 1)
	assert.Equal(t, "session_token", cookies[0].Name)
}

func TestRegisterHandler_ValidationErrors(t *testing.T) {
	handlers, _, _ := initTestEnv(t)

	testCases := []struct {
		name    string
		request registerRequest
	}{
		{
			name: "Short login",
			request: registerRequest{
				Login:    "user",
				Email:    "new_user@test.com",
				Password: "some_password",
			},
		},
		{
			name: "Invalid email",
			request: registerRequest{
				Login:    "user_login",
				Email:    "invalid_email",
				Password: "some_password",
			},
		},
		{
			name: "Short password",
			request: registerRequest{
				Login:    "user_login",
				Email:    "new_user@test.com",
				Password: "usr",
			},
		},
		{
			name: "Empty login",
			request: registerRequest{
				Login:    "",
				Email:    "new_user@test.com",
				Password: "usr",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(tc.request)
			require.NoError(t, err)

			req := httptest.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handlers.RegisterHandler(rr, req)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
		})
	}
}

func TestRegisterHandler_DuplicateUser(t *testing.T) {
	handlers, _, _ := initTestEnv(t)

	reqBody := registerRequest{
		Login:    "user_login",
		Email:    "new_user@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.RegisterHandler(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)

	reqBody2 := registerRequest{
		Login:    "user_login",
		Email:    "new_user@test.com",
		Password: "some_password",
	}
	body2, err := json.Marshal(reqBody2)
	require.NoError(t, err)

	req2 := httptest.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(body2))
	req2.Header.Set("Content-Type", "application/json")
	rr2 := httptest.NewRecorder()
	handlers.RegisterHandler(rr2, req2)
	assert.Equal(t, http.StatusConflict, rr2.Code)
}

func TestLoginHandler(t *testing.T) {
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

	loginReq := loginRequest{
		Login:    "user_login",
		Password: "some_password",
	}
	body, err = json.Marshal(loginReq)
	require.NoError(t, err)

	req = httptest.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	handlers.LoginHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var response loginResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.ID)

	cookies := rr.Result().Cookies()
	assert.Len(t, cookies, 1)
	assert.Equal(t, "session_token", cookies[0].Name)
}

func TestLogoutHandler_Success(t *testing.T) {
	handlers, _, _ := initTestEnv(t)

	req := httptest.NewRequest("POST", "/api/v1/logout", nil)
	rr := httptest.NewRecorder()
	handlers.LogoutHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var response logoutResponse
	err := json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "ok", response.Status)

	cookies := rr.Result().Cookies()
	assert.Len(t, cookies, 1)
	assert.Equal(t, "session_token", cookies[0].Name)
	assert.Equal(t, "", cookies[0].Value)
	assert.True(t, cookies[0].Expires.Before(time.Now()))
}

func TestLoginHandler_InvalidCredentials(t *testing.T) {
	handlers, _, _ := initTestEnv(t)

	reqBody := registerRequest{
		Login:    "user_login",
		Email:    "new_user@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.RegisterHandler(rr, req)

	testCases := []struct {
		name           string
		request        loginRequest
		expectedStatus int
	}{
		{
			name: "user not exist",
			request: loginRequest{
				Login:    "no_user",
				Password: "some_password",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "wrong password",
			request: loginRequest{
				Login:    "user_login",
				Password: "wrong_password",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "empty credentials",
			request: loginRequest{
				Login:    "",
				Password: "",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(tc.request)
			require.NoError(t, err)

			req := httptest.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handlers.LoginHandler(rr, req)
			assert.Equal(t, tc.expectedStatus, rr.Code)
		})
	}
}

func TestRegisterHandler_InvalidJSON(t *testing.T) {
	handlers, _, _ := initTestEnv(t)

	req := httptest.NewRequest("POST", "/api/v1/register", bytes.NewBufferString("invalid json"))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.RegisterHandler(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestLoginHandler_InvalidJSON(t *testing.T) {
	handlers, _, _ := initTestEnv(t)

	req := httptest.NewRequest("POST", "/api/v1/login", bytes.NewBufferString("invalid json"))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.LoginHandler(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
