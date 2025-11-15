package jwtmanager

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Claims struct {
	UserID    string `json:"sub"`
	SessionID string `json:"jti"`
	Exp       int64  `json:"exp"`
	Iat       int64  `json:"iat"`
	Role      string `json:"role"`
}

type Manager struct {
	secretKey      string
	accessTokenTTL time.Duration
}

func NewManager(secretKey string, ttl time.Duration) *Manager {
	return &Manager{secretKey: secretKey, accessTokenTTL: ttl}
}

func (m *Manager) GetTTL() time.Duration {
	return m.accessTokenTTL
}

func (m *Manager) Generate(userID string) (string, error) {
	now := time.Now()
	expiresAt := now.Add(m.accessTokenTTL)

	claims := Claims{
		UserID:    userID,
		SessionID: uuid.New().String(),
		Exp:       expiresAt.Unix(),
		Iat:       now.Unix(),
	}

	header := map[string]string{"alg": "HS256", "typ": "JWT"}
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerEncoded := base64.RawURLEncoding.EncodeToString(headerJSON)

	payloadJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	payloadEncoded := base64.RawURLEncoding.EncodeToString(payloadJSON)

	signatureInput := headerEncoded + "." + payloadEncoded
	signature := m.createSignature(signatureInput)

	return signatureInput + "." + signature, nil
}

func (m *Manager) createSignature(data string) string {
	hasher := hmac.New(sha256.New, []byte(m.secretKey))
	hasher.Write([]byte(data))
	return base64.RawURLEncoding.EncodeToString(hasher.Sum(nil))
}

func (m *Manager) Validate(tokenString string) (*Claims, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format")
	}

	signatureInput := parts[0] + "." + parts[1]
	expectedSignature := m.createSignature(signatureInput)
	if !hmac.Equal([]byte(parts[2]), []byte(expectedSignature)) {
		return nil, errors.New("invalid signature")
	}

	payloadJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	var claims Claims
	if err := json.Unmarshal(payloadJSON, &claims); err != nil {
		return nil, fmt.Errorf("failed to parse claims: %w", err)
	}

	if time.Now().Unix() > claims.Exp {
		return nil, errors.New("token expired")
	}

	return &claims, nil
}
