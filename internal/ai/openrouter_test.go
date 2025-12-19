package ai

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"spotify/microservices/playlist/dto"

	"github.com/stretchr/testify/require"
)

func TestNewOpenRouter_Defaults(t *testing.T) {
	or := NewOpenRouter(OpenRouterConfig{})

	require.Equal(t, 10, or.maxTracks)
	require.NotNil(t, or.http)
	require.Equal(t, 20*time.Second, or.http.Timeout)
}

func TestOpenRouter_GeneratePlaylistMeta_Success(t *testing.T) {
	client := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			require.Equal(t, "/api/v1/chat/completions", r.URL.Path)
			require.Equal(t, "Bearer test-key", r.Header.Get("Authorization"))

			body := `{
				"choices": [{
					"message": {
						"content": "{\"title\":\"Chill\",\"description\":\"Relax vibes\"}"
					}
				}]
			}`

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(body)),
				Header:     make(http.Header),
			}, nil
		}),
	}

	or := NewOpenRouter(OpenRouterConfig{
		AuthKey: "test-key",
		Model:   "test-model",
	})
	or.http = client

	tracks := []dto.Track{
		{
			Title: "Track 1",
			Artists: []dto.Artist{
				{Name: "Artist 1"},
			},
		},
	}

	meta, err := or.GeneratePlaylistMeta(context.Background(), tracks)
	require.NoError(t, err)
	require.Len(t, meta, 1)
	require.Equal(t, "Chill", meta[0].Title)
	require.Equal(t, "Relax vibes", meta[0].Description)
}

func testStatus(t *testing.T, status int, expected error) {
	client := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: status,
				Body:       io.NopCloser(strings.NewReader("{}")),
				Header:     make(http.Header),
			}, nil
		}),
	}

	or := NewOpenRouter(OpenRouterConfig{
		AuthKey: "key",
		Model:   "model",
	})
	or.http = client

	_, err := or.GeneratePlaylistMeta(context.Background(), []dto.Track{
		{Title: "t"},
	})

	require.ErrorIs(t, err, expected)
}

func TestOpenRouter_GeneratePlaylistMeta_AuthError(t *testing.T) {
	testStatus(t, http.StatusUnauthorized, ErrAIAuth)
}

func TestOpenRouter_GeneratePlaylistMeta_RateLimit(t *testing.T) {
	testStatus(t, http.StatusTooManyRequests, ErrAIRateLimit)
}

func TestOpenRouter_GeneratePlaylistMeta_Unavailable(t *testing.T) {
	testStatus(t, http.StatusInternalServerError, ErrAIUnavailable)
}

func TestOpenRouter_GeneratePlaylistMeta_NoChoices(t *testing.T) {
	client := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"choices":[]}`)),
				Header:     make(http.Header),
			}, nil
		}),
	}

	or := NewOpenRouter(OpenRouterConfig{
		AuthKey: "key",
		Model:   "model",
	})
	or.http = client

	_, err := or.GeneratePlaylistMeta(context.Background(), []dto.Track{
		{Title: "t"},
	})

	require.ErrorIs(t, err, ErrAINoChoices)
}
