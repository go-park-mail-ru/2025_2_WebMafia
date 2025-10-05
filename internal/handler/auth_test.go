package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthHandler(t *testing.T) {
	reqBody := registerRequest{
		Login:    "user_login",
		Email:    "new_user@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var response registerResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.ID)

	cookies := resp.Cookies()
	assert.Len(t, cookies, 1)
	assert.Equal(t, "session_token", cookies[0].Name)
}

func TestRegisterHandler_ValidationErrors(t *testing.T) {
	testCases := []struct {
		name    string
		request registerRequest
	}{
		{
			name: "Short login",
			request: registerRequest{
				Login:    "user",
				Email:    "new_user_short@test.com",
				Password: "some_password",
			},
		},
		{
			name: "Invalid email",
			request: registerRequest{
				Login:    "user_login_1",
				Email:    "invalid_email",
				Password: "some_password",
			},
		},
		{
			name: "Short password",
			request: registerRequest{
				Login:    "user_login_2",
				Email:    "new_user_2@test.com",
				Password: "usr",
			},
		},
		{
			name: "Empty login",
			request: registerRequest{
				Login:    "",
				Email:    "new_user_empty@test.com",
				Password: "usr",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(tc.request)
			require.NoError(t, err)

			resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
			require.NoError(t, err)
			defer resp.Body.Close()
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		})
	}
}

func TestRegisterHandler_DuplicateUser(t *testing.T) {
	reqBody := registerRequest{
		Login:    "user_login_duplicate",
		Email:    "duplicate_user@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	resp2, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp2.Body.Close()
	assert.Equal(t, http.StatusConflict, resp2.StatusCode)
}

func TestLoginHandler(t *testing.T) {
	registerReq := registerRequest{
		Login:    "user_login_3",
		Email:    "new_user_3@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(registerReq)
	require.NoError(t, err)

	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	resp.Body.Close()

	loginReq := loginRequest{
		Login:    "user_login_3",
		Password: "some_password",
	}
	body, err = json.Marshal(loginReq)
	require.NoError(t, err)

	resp, err = testClient.Post(testServer.URL+"/api/v1/login", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response loginResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.ID)

	cookies := resp.Cookies()
	assert.Len(t, cookies, 1)
	assert.Equal(t, "session_token", cookies[0].Name)
}

func TestLogoutHandler_Success(t *testing.T) {
	registerReq := registerRequest{
		Login:    "logout_test_user",
		Email:    "logout_test@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(registerReq)
	require.NoError(t, err)

	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	resp.Body.Close()

	loginReq := loginRequest{
		Login:    "logout_test_user",
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

	req, err := http.NewRequest("POST", testServer.URL+"/api/v1/logout", nil)
	require.NoError(t, err)
	req.AddCookie(sessionCookie)

	resp, err = testClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response logoutResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "ok", response.Status)

	logoutCookies := resp.Cookies()
	assert.Len(t, logoutCookies, 1)
	assert.Equal(t, "session_token", logoutCookies[0].Name)
	assert.Equal(t, "", logoutCookies[0].Value)
	assert.True(t, logoutCookies[0].Expires.Before(time.Now()))
}

func TestLoginHandler_InvalidCredentials(t *testing.T) {
	reqBody := registerRequest{
		Login:    "user_login_4",
		Email:    "new_user_4@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	resp.Body.Close()

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
				Login:    "user_login_4",
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

			resp, err := testClient.Post(testServer.URL+"/api/v1/login", "application/json", bytes.NewBuffer(body))
			require.NoError(t, err)
			defer resp.Body.Close()
			assert.Equal(t, tc.expectedStatus, resp.StatusCode)
		})
	}
}

func TestRegisterHandler_InvalidJSON(t *testing.T) {
	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBufferString("invalid json"))
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestLoginHandler_InvalidJSON(t *testing.T) {
	resp, err := testClient.Post(testServer.URL+"/api/v1/login", "application/json", bytes.NewBufferString("invalid json"))
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
