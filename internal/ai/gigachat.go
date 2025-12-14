package ai

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"spotify/microservices/playlist/dto"
	"strings"
	"time"
)

const (
	baseURL      = "https://gigachat.devices.sberbank.ru/api/v1"
	chatEndpoint = "/chat/completions"
	prompt       = `
	Задача:
	Сгенерируй название и описание плейлиста на основе этих треков.
	Название: до 80 символов.
	Описание: до 200 символов, атмосферное, без перечисления всех треков.
	
	Ответ верни строго в JSON без markdown и комментариев:
	{"title":"...","description":"..."}`
)

type GigaChat struct {
	http      *http.Client
	tm        *tokenManager
	model     string
	maxTracks int
}

type GigaChatConfig struct {
	AuthKey            string
	Model              string
	Timeout            time.Duration
	MaxTracks          int
	InsecureSkipVerify bool
}

func NewGigaChat(cfg GigaChatConfig) *GigaChat {
	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = 20 * time.Second
	}

	maxTracks := cfg.MaxTracks
	if maxTracks <= 0 {
		maxTracks = 10
	}

	transport := http.DefaultTransport.(*http.Transport).Clone()

	if cfg.InsecureSkipVerify {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	client := &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	tm := newTokenManager(cfg.AuthKey, client)

	return &GigaChat{
		http:      client,
		tm:        tm,
		model:     cfg.Model,
		maxTracks: maxTracks,
	}
}

func (g *GigaChat) GeneratePlaylistMeta(ctx context.Context, tracks []dto.Track) (string, string, error) {
	max := g.maxTracks
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

	token, err := g.tm.getToken(ctx)
	if err != nil {
		return "", "", ErrAIAuth
	}

	reqBody := ChatRequest{
		Model: g.model,
		Messages: []ChatMessage{
			{Role: "user", Content: sb.String()},
		},
	}

	body, _ := json.Marshal(reqBody)

	url := baseURL + chatEndpoint
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := g.http.Do(req)
	if err != nil {
		return "", "", ErrAIUnavailable
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusUnauthorized, http.StatusForbidden:
			return "", "", ErrAIAuth
		case http.StatusTooManyRequests:
			return "", "", ErrAIRateLimit
		default:
			return "", "", ErrAIUnavailable
		}
	}

	var out ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", "", err
	}

	if len(out.Choices) == 0 {
		return "", "", fmt.Errorf("empty response")
	}

	raw := out.Choices[0].Message.Content

	clean := strings.TrimSpace(raw)
	clean = strings.TrimPrefix(clean, "```json")
	clean = strings.TrimPrefix(clean, "```")
	clean = strings.TrimSuffix(clean, "```")
	clean = strings.TrimSpace(clean)

	type Meta struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	var meta Meta
	if err := json.Unmarshal([]byte(clean), &meta); err != nil {
		return "", "", ErrAIUnavailable
	}

	return meta.Title, meta.Description, nil
}
