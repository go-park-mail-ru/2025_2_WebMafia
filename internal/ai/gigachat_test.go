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

func TestGigaChat_GeneratePlaylistMeta_EmptyChoices(t *testing.T) {
	client := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			body := io.NopCloser(strings.NewReader(`{
				"choices": []
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

	_, err := g.GeneratePlaylistMeta(
		context.Background(),
		[]dto.Track{{Title: "Song"}},
	)

	require.Error(t, err)
}

func TestGigaChat_GeneratePlaylistMeta_AllChoicesInvalidJSON(t *testing.T) {
	client := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			body := io.NopCloser(strings.NewReader(`{
				"choices": [{
					"message": { "content": "not a json" }
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

	_, err := g.GeneratePlaylistMeta(
		context.Background(),
		[]dto.Track{{Title: "Song"}},
	)

	require.Error(t, err)
}

func TestGigaChat_GeneratePlaylistMeta_RateLimit(t *testing.T) {
	client := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Body:       io.NopCloser(strings.NewReader("")),
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

	_, err := g.GeneratePlaylistMeta(
		context.Background(),
		[]dto.Track{{Title: "Song"}},
	)

	require.ErrorIs(t, err, ErrAIRateLimit)
}
