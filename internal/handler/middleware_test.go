package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthMiddleware(t *testing.T) {
	registerReq := registerRequest{
		Login:    "middleware_user",
		Email:    "middleware@test.com",
		Password: "some_password",
	}
	body, err := json.Marshal(registerReq)
	require.NoError(t, err)

	resp, err := testClient.Post(testServer.URL+"/api/v1/register", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	resp.Body.Close()

	loginReq := loginRequest{
		Login:    "middleware_user",
		Password: "some_password",
	}
	body, err = json.Marshal(loginReq)
	require.NoError(t, err)

	resp, err = testClient.Post(testServer.URL+"/api/v1/login", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	cookies := resp.Cookies()
	var authCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "session_token" {
			authCookie = cookie
			break
		}
	}
	assert.NotNil(t, authCookie)

	req, err := http.NewRequest("GET", testServer.URL+"/api/v1/home", nil)
	require.NoError(t, err)
	req.AddCookie(authCookie)

	resp, err = testClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestAuthMiddleware_NoToken(t *testing.T) {
	resp, err := testClient.Get(testServer.URL + "/api/v1/home")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	req, err := http.NewRequest("GET", testServer.URL+"/api/v1/home", nil)
	require.NoError(t, err)
	req.AddCookie(&http.Cookie{
		Name:  "session_token",
		Value: "invalid-token",
	})

	resp, err := testClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
