package ai

import (
	"context"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"spotify/microservices/playlist/dto"
	"strings"
	"testing"
	"time"
)

func TestTokenManager_GetToken_Success(t *testing.T) {
	oldClient := http.DefaultClient
	defer func() { http.DefaultClient = oldClient }()

	http.DefaultClient = &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			body := io.NopCloser(strings.NewReader(`{
				"access_token": "token123",
				"expires_at": 9999999999999
			}`))

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       body,
				Header:     make(http.Header),
			}, nil
		}),
	}

	tm := newTokenManager("key")

	token, err := tm.getToken(context.Background())
	require.NoError(t, err)
	require.Equal(t, "token123", token)
}

func TestGigaChat_GeneratePlaylistMeta_Success(t *testing.T) {
	client := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			body := io.NopCloser(strings.NewReader(`{
				"choices": [{
					"message": {
						"role": "assistant",
						"content": "{\"title\":\"My playlist\",\"description\":\"Nice vibes\"}"
					}
				}]
			}`))

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       body,
				Header:     make(http.Header),
			}, nil
		}),
	}

	g := &GigaChat{
		http: client,
		tm: &tokenManager{
			token: "token",
			exp:   time.Now().Add(time.Hour),
		},
		model: "test",
	}

	title, desc, err := g.GeneratePlaylistMeta(
		context.Background(),
		[]dto.Track{{Title: "Song"}},
	)

	require.NoError(t, err)
	require.Equal(t, "My playlist", title)
	require.Equal(t, "Nice vibes", desc)
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}
