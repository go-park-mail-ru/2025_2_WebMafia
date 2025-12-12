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
	sb.WriteString("У меня есть плейлист. Ниже список треков:\\n")
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

	sb.WriteString("\nНа основе этих треков придумай:\n")
	sb.WriteString("1) красивое, короткое и цепляющее название плейлиста\n")
	sb.WriteString("2) атмосферное, чуть более подробное описание\n\n")

	sb.WriteString("Верни результат строго в виде JSON без комментариев, без markdown и без лишних пояснений.\n")
	sb.WriteString("Структура должна быть такой:\n")
	sb.WriteString("{\"title\": \"...\"," +
		"\"description\": \"...\"}\n")

	token, err := g.tm.getToken(ctx)
	if err != nil {
		return "", "", err
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

	fmt.Println("RESPONSE START:")
	fmt.Println(raw)
	fmt.Println("RESPONSE END:")

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
