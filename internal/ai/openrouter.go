package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"spotify/microservices/playlist/dto"
	"strings"
	"time"
)

const (
	openRouterBaseURL      = "https://openrouter.ai/api/v1"
	openRouterChatEndpoint = "/chat/completions"
)

type OpenRouter struct {
	http      *http.Client
	authKey   string
	model     string
	maxTracks int
}

type OpenRouterConfig struct {
	AuthKey   string
	Model     string
	Timeout   time.Duration
	MaxTracks int
}

func NewOpenRouter(cfg OpenRouterConfig) *OpenRouter {
	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = 20 * time.Second
	}

	maxTracks := cfg.MaxTracks
	if maxTracks <= 0 {
		maxTracks = 10
	}

	client := &http.Client{
		Timeout: timeout,
	}

	return &OpenRouter{
		http:      client,
		authKey:   cfg.AuthKey,
		model:     cfg.Model,
		maxTracks: maxTracks,
	}
}

func (o *OpenRouter) GeneratePlaylistMeta(ctx context.Context, tracks []dto.Track) ([]Meta, error) {
	max := o.maxTracks
	if len(tracks) < max {
		max = len(tracks)
	}

	var sb strings.Builder
	sb.WriteString("Список треков:\n")

	for i := 0; i < max; i++ {
		t := tracks[i]
		sb.WriteString(fmt.Sprintf("%d. %s — ", i+1, t.Title))

		names := make([]string, len(t.Artists))
		for j, a := range t.Artists {
			names[j] = a.Name
		}
		sb.WriteString(strings.Join(names, ", "))
		sb.WriteString("\n")
	}

	sb.WriteString("\n")
	sb.WriteString(prompt)

	reqBody := ChatRequest{
		Model: o.model,
		Messages: []ChatMessage{
			{Role: "user", Content: sb.String()},
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	chatURL, err := url.JoinPath(openRouterBaseURL, openRouterChatEndpoint)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		chatURL,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+o.authKey)
	req.Header.Set("HTTP-Referer", "wave-music.ru")
	req.Header.Set("X-Title", "WaveMusic")
	req.Header.Set("User-Agent", "WaveMusic/1.0")

	resp, err := o.http.Do(req)
	if err != nil {
		return nil, ErrAIUnavailable
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusUnauthorized, http.StatusForbidden:
			return nil, ErrAIAuth
		case http.StatusTooManyRequests:
			return nil, ErrAIRateLimit
		default:
			return nil, ErrAIUnavailable
		}
	}

	var out ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, ErrAIUnavailable
	}

	if len(out.Choices) == 0 {
		return nil, ErrAINoChoices
	}

	metas := make([]Meta, 0, len(out.Choices))

	for _, choice := range out.Choices {
		raw := strings.TrimSpace(choice.Message.Content)

		clean := strings.TrimPrefix(raw, "```json")
		clean = strings.TrimPrefix(clean, "```")
		clean = strings.TrimSuffix(clean, "```")
		clean = strings.TrimSpace(clean)

		var meta Meta
		if err := json.Unmarshal([]byte(clean), &meta); err != nil {
			continue
		}

		metas = append(metas, meta)
	}

	if len(metas) == 0 {
		return nil, ErrAINoChoices
	}

	return metas, nil
}
