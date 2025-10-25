package csrfmanager

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type payload struct {
	UserID    string `json:"sub"`
	SessionID string `json:"jti"`
	ExpiresAt int64  `json:"exp"`
}

type Manager struct {
	secret []byte
	ttl    time.Duration
}

func NewManager(secret string, ttl time.Duration) *Manager {
	return &Manager{secret: []byte(secret), ttl: ttl}
}

func (m *Manager) Generate(userID, sessionID string) (string, error) {
	p := payload{
		UserID:    userID,
		SessionID: sessionID,
		ExpiresAt: time.Now().Add(m.ttl).Unix(),
	}
	data, err := json.Marshal(p)
	if err != nil {
		return "", fmt.Errorf("failed to marshal csrf payload: %w", err)
	}

	h := hmac.New(sha256.New, m.secret)
	_, err = h.Write(data)
	if err != nil {
		return "", fmt.Errorf("failed to write data to hmac: %w", err)
	}
	signature := hex.EncodeToString(h.Sum(nil))

	token := fmt.Sprintf("%s:%d", signature, p.ExpiresAt)

	return token, nil
}

func (m *Manager) Check(userID, sessionID, clientToken string) (bool, error) {
	parts := strings.Split(clientToken, ":")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid csrf token format")
	}
	signatureFromClient := parts[0]
	expiresAtStr := parts[1]

	expiresAt, err := strconv.ParseInt(expiresAtStr, 10, 64)
	if err != nil {
		return false, fmt.Errorf("invalid expiration time in csrf token: %w", err)
	}

	if time.Now().Unix() > expiresAt {
		return false, fmt.Errorf("csrf token expired")
	}

	expectedPayload := payload{
		UserID:    userID,
		SessionID: sessionID,
		ExpiresAt: expiresAt,
	}
	expectedData, err := json.Marshal(expectedPayload)
	if err != nil {
		return false, fmt.Errorf("failed to marshal expected payload for check: %w", err)
	}

	h := hmac.New(sha256.New, m.secret)
	_, err = h.Write(expectedData)
	if err != nil {
		return false, fmt.Errorf("failed to write data to hmac for check: %w", err)
	}
	expectedSignature := []byte(hex.EncodeToString(h.Sum(nil)))

	isValid := hmac.Equal([]byte(signatureFromClient), expectedSignature)
	if !isValid {
		return false, fmt.Errorf("invalid csrf token signature")
	}

	return true, nil
}
