package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"spotify/internal/model"
	"spotify/internal/repository"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo       repository.UserRepository
	jwtSecretKey   string
	accessTokenTTL time.Duration
}

func NewAuthService(userRepo repository.UserRepository, secret string, ttl time.Duration) AuthService {
	return &authService{
		userRepo:       userRepo,
		jwtSecretKey:   secret,
		accessTokenTTL: ttl,
	}
}

func (s *authService) Register(ctx context.Context, email, login, password string) (*model.User, error) {
	_, err := s.userRepo.GetUserByEmail(ctx, email)
	if err == nil {
		return &model.User{}, ErrEmailExists
	} else if !errors.Is(err, repository.ErrUserNotFound) {
		return &model.User{}, fmt.Errorf("internal error checking email: %w", err)
	}

	_, err = s.userRepo.GetUserByLogin(ctx, login)
	if err == nil {
		return &model.User{}, ErrLoginExists
	} else if !errors.Is(err, repository.ErrUserNotFound) {
		return &model.User{}, fmt.Errorf("internal error checking login: %w", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return &model.User{}, fmt.Errorf("failed to hash password: %w", err)
	}

	newUser, err := s.userRepo.CreateUser(ctx, model.User{
		Login:        login,
		PasswordHash: string(hashedPassword),
		Email:        email,
	})
	if err != nil {
		return &model.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return newUser, nil
}

func (s *authService) Login(ctx context.Context, login, password string) (string, error) {
	user, err := s.userRepo.GetUserByLogin(ctx, login)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return "", ErrInvalidCredentials
		}
		return "", fmt.Errorf("failed to get user by login: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := s.generateJWT(user.ID.String())
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

type Claims struct {
	UserID string `json:"sub"`
	Exp    int64  `json:"exp"`
	Iat    int64  `json:"iat"`
}

func (s *authService) generateJWT(userID string) (string, error) {
	now := time.Now()
	expiresAt := now.Add(s.accessTokenTTL)

	claims := Claims{
		UserID: userID,
		Exp:    expiresAt.Unix(),
		Iat:    now.Unix(),
	}

	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}
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
	signature := s.createSignature(signatureInput)

	return signatureInput + "." + signature, nil
}

func (s *authService) createSignature(data string) string {
	h := hmac.New(sha256.New, []byte(s.jwtSecretKey))
	h.Write([]byte(data))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}
func (s *authService) ValidateToken(tokenString string) (string, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return "", errors.New("invalid token format")
	}
	headerEncoded, payloadEncoded, signatureEncoded := parts[0], parts[1], parts[2]
	signatureInput := headerEncoded + "." + payloadEncoded
	expectedSignature := s.createSignature(signatureInput)
	if !hmac.Equal([]byte(signatureEncoded), []byte(expectedSignature)) {
		return "", errors.New("invalid signature")
	}
	payloadJSON, err := base64.RawURLEncoding.DecodeString(payloadEncoded)
	if err != nil {
		return "", fmt.Errorf("failed to decode payload: %w", err)
	}
	var claims Claims
	if err := json.Unmarshal(payloadJSON, &claims); err != nil {
		return "", fmt.Errorf("failed to parse claims: %w", err)
	}
	if time.Now().Unix() > claims.Exp {
		return "", errors.New("token expired")
	}
	return claims.UserID, nil
}
