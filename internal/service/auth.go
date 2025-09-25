package service

import (
	"context"
	"errors"
	"fmt"
	"spotify/internal/model"
	"spotify/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func (s *authService) Register(ctx context.Context, email, login, password string) (model.User, error) {
	_, err := s.userRepo.GetUserByEmail(ctx, email)
	if err == nil {
		return model.User{}, ErrEmailExists
	} else if !errors.Is(err, repository.ErrUserNotFound) {
		return model.User{}, fmt.Errorf("internal error checking email: %w", err)
	}

	_, err = s.userRepo.GetUserByLogin(ctx, login)
	if err == nil {
		return model.User{}, ErrLoginExists
	} else if !errors.Is(err, repository.ErrUserNotFound) {
		return model.User{}, fmt.Errorf("internal error checking login: %w", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to hash password: %w", err)
	}

	newUser, err := s.userRepo.CreateUser(ctx, model.User{
		Login:        login,
		PasswordHash: string(hashedPassword),
		Email:        email,
	})
	if err != nil {
		return model.User{}, fmt.Errorf("failed to create user: %w", err)
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

	claims := jwt.RegisteredClaims{
		Subject:   user.ID.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.accessTokenTTL)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.jwtSecretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}
