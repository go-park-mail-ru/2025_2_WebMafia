package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"spotify/internal/model"
	"spotify/internal/store"
	"spotify/pkg/response"

	"golang.org/x/crypto/bcrypt"
)

type signUpRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpResponse struct {
	ID string `json:"id"`
}

func (i *signUpRequest) validate() error {
	if len(i.Login) < 5 {
		return fmt.Errorf("login is too short (minimum 5 chars)")
	}
	if len(i.Password) < 8 {
		return fmt.Errorf("password is too short (minimum 8 chars)")
	}
	return nil
}

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"access_token"`
}

func (i *loginRequest) validate() error {
	if i.Login == "" {
		return fmt.Errorf("empty login")
	}
	if i.Password == "" {
		return fmt.Errorf("empty password")
	}
	return nil
}

func (h *Handlers) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req signUpRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: "invalid request body"})
		return
	}

	if err := req.validate(); err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return
	}

	_, err := h.store.GetUserByEmail(r.Context(), req.Email)
	if !errors.Is(err, store.ErrUserNotFound) {
		response.JSON(w, http.StatusConflict, response.ErrorResponse{Error: "user with this email already exists"})
		return
	}

	_, err = h.store.GetUserByLogin(r.Context(), req.Login)
	if !errors.Is(err, store.ErrUserNotFound) {
		response.JSON(w, http.StatusConflict, response.ErrorResponse{Error: "user with this login already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorResponse{Error: "failed to hash password"})
		return
	}

	newUser, err := h.store.CreateUser(r.Context(), model.User{
		Login:        req.Login,
		PasswordHash: string(hashedPassword),
		Email:        req.Email,
	})
	if err != nil {
		if errors.Is(err, store.ErrUserAlreadyExists) {
			response.JSON(w, http.StatusConflict, response.ErrorResponse{Error: err.Error()})
			return
		}
		response.JSON(w, http.StatusInternalServerError, response.ErrorResponse{Error: "failed to create user"})
		return
	}

	response.JSON(w, http.StatusCreated, signUpResponse{ID: newUser.ID.String()})
}

func (h *Handlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req loginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: "invalid request body"})
		return
	}
	if err := req.validate(); err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return
	}

	user, err := h.store.GetUserByLogin(r.Context(), req.Login)
	if err != nil {
		if errors.Is(err, store.ErrUserNotFound) {
			response.JSON(w, http.StatusUnauthorized, response.ErrorResponse{Error: "invalid login or password"})
			return
		}
		response.JSON(w, http.StatusInternalServerError, response.ErrorResponse{Error: "internal server error"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		response.JSON(w, http.StatusUnauthorized, response.ErrorResponse{Error: "invalid login or password"})
		return
	}

	token, err := h.generateJWT(user.ID.String())
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorResponse{Error: "failed to generate token"})
		return
	}

	response.JSON(w, http.StatusOK, loginResponse{Token: token})
}

type claims struct {
	UserID string `json:"sub"`
	Exp    int64  `json:"exp"`
	Iat    int64  `json:"iat"`
}

func (h *Handlers) generateJWT(userID string) (string, error) {
	now := time.Now()
	expiresAt := now.Add(h.accessTokenTTL)

	claims := claims{
		UserID: userID,
		Exp:    expiresAt.Unix(),
		Iat:    now.Unix(),
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
	signature := h.createSignature(signatureInput)

	return signatureInput + "." + signature, nil
}

func (h *Handlers) createSignature(data string) string {
	hasher := hmac.New(sha256.New, []byte(h.jwtSecretKey))
	hasher.Write([]byte(data))
	return base64.RawURLEncoding.EncodeToString(hasher.Sum(nil))
}

func (h *Handlers) validateToken(tokenString string) (string, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return "", errors.New("invalid token format")
	}

	signatureInput := parts[0] + "." + parts[1]
	expectedSignature := h.createSignature(signatureInput)
	if !hmac.Equal([]byte(parts[2]), []byte(expectedSignature)) {
		return "", errors.New("invalid signature")
	}

	payloadJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("failed to decode payload: %w", err)
	}

	var claims claims
	if err := json.Unmarshal(payloadJSON, &claims); err != nil {
		return "", fmt.Errorf("failed to parse claims: %w", err)
	}

	if time.Now().Unix() > claims.Exp {
		return "", errors.New("token expired")
	}

	return claims.UserID, nil
}
