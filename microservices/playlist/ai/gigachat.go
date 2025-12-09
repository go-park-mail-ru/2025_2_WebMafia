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
	sb.WriteString("Вот список треков:\n")
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

	sb.WriteString("\nСгенерируй короткое название (до 80 символов) и короткое описание (до 200 символов). Формат:\n")
	sb.WriteString("title: ...\ndescription: ...")

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
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", "", fmt.Errorf("gigachat error: %s", resp.Status)
	}

	var out ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", "", err
	}

	if len(out.Choices) == 0 {
		return "", "", fmt.Errorf("empty response")
	}

	raw := out.Choices[0].Message.Content

	var title, desc string
	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "title:") {
			title = strings.TrimSpace(strings.TrimPrefix(line, "title:"))
		} else if strings.HasPrefix(line, "description:") {
			desc = strings.TrimSpace(strings.TrimPrefix(line, "description:"))
		}
	}

	return title, desc, nil
}
