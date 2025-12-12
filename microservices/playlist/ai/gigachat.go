package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"spotify/microservices/playlist/dto"
)

const chatURL = "https://gigachat.devices.sberbank.ru/api/v1/chat/completions"

type GigaChat struct {
	http  *http.Client
	tm    *tokenManager
	model string
}

type GigaChatConfig struct {
	AuthKey string
	Model   string
}

func NewGigaChat(cfg GigaChatConfig) *GigaChat {
	return &GigaChat{
		http:  &http.Client{Timeout: 20 * time.Second},
		tm:    newTokenManager(cfg.AuthKey),
		model: cfg.Model,
	}
}

func (g *GigaChat) GeneratePlaylistMeta(ctx context.Context, tracks []dto.Track) (string, string, error) {
	max := 10
	if len(tracks) < 10 {
		max = len(tracks)
	}

	var sb strings.Builder
	sb.WriteString("Список треков:\\n")
	for i := 0; i < max; i++ {
		t := tracks[i]
		sb.WriteString(fmt.Sprintf("%d. %s — ", i+1, t.Title))
		for j, a := range t.Artists {
			if j > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(a.Name)
		}
		sb.WriteString("\n")
	}

	sb.WriteString("\nЗадача:\n")
	sb.WriteString("Сгенерируй название и описание плейлиста на основе этих треков.\n")
	sb.WriteString("Название: до 80 символов.\n")
	sb.WriteString("Описание: до 200 символов, атмосферное, без перечисления всех треков.\n")

	sb.WriteString("Ответ верни строго в JSON без markdown и комментариев:\n")
	sb.WriteString("{\"title\":\"...\",\"description\":\"...\"}")

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

	req, err := http.NewRequestWithContext(ctx, "POST", chatURL, bytes.NewBuffer(body))
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
			if resp.StatusCode >= 500 {
				return "", "", ErrAIUnavailable
			}
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
