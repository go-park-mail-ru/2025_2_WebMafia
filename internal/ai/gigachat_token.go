package ai

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	gigaChatAuthBaseURL = "https://ngw.devices.sberbank.ru:9443/api/v2"
	gigaChatAuthRoute   = "/oauth"
)

type tokenManager struct {
	authKey string
	client  *http.Client
	token   string
	exp     time.Time
	mu      sync.Mutex
}

func newTokenManager(authKey string, client *http.Client) *tokenManager {
	return &tokenManager{
		authKey: authKey,
		client:  client,
	}
}

func (tm *tokenManager) getToken(ctx context.Context) (string, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if tm.token != "" && time.Now().Before(tm.exp) {
		return tm.token, nil
	}

	values := url.Values{}
	values.Set("scope", "GIGACHAT_API_PERS")

	tokenURL := gigaChatAuthBaseURL + gigaChatAuthRoute

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		tokenURL,
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("RqUID", uuid.NewString())
	req.Header.Set("Authorization", "Basic "+tm.authKey)

	resp, err := tm.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("token request failed: %s", string(body))
	}

	var tr TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return "", err
	}

	if tr.AccessToken == "" {
		return "", errors.New("empty access token")
	}
	tm.token = tr.AccessToken
	tm.exp = time.UnixMilli(tr.ExpiresAt)
	return tm.token, nil
}
