package handler

import (
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

	resp, _ := makeRequest(t, "POST", "/api/v1/register", body)
	resp.Body.Close()

	loginReq := loginRequest{
		Login:    "middleware_user",
		Password: "some_password",
	}
	body, err = json.Marshal(loginReq)
	require.NoError(t, err)

	resp, _ = makeRequest(t, "POST", "/api/v1/login", body)
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

	resp, _ = makeRequest(t, "GET", "/api/v1/home", nil, authCookie)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestAuthMiddleware_NoToken(t *testing.T) {
	resp, _ := makeRequest(t, "GET", "/api/v1/home", nil)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	invalidCookie := &http.Cookie{
		Name:  "session_token",
		Value: "invalid-token",
	}

	resp, _ := makeRequest(t, "GET", "/api/v1/home", nil, invalidCookie)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
