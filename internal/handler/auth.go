package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"spotify/internal/model"
	"spotify/internal/store"
	"spotify/pkg/response"

	"golang.org/x/crypto/bcrypt"
)

type registerRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerResponse struct {
	ID string `json:"id"`
}

func (i *registerRequest) validate() error {
	if len(i.Login) < 5 {
		return fmt.Errorf("login is too short (minimum 5 chars)")
	}
	if len(i.Password) < 8 {
		return fmt.Errorf("password is too short (minimum 8 chars)")
	}
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(i.Email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type loginResponse struct {
	ID string `json:"id"`
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

type logoutResponse struct {
	Status string `json:"status"`
}

func (h *Handlers) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req registerRequest

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

	token, err := h.jwtManager.Generate(newUser.ID.String())
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorResponse{Error: "failed to generate token"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(h.jwtManager.GetTTL()),
		HttpOnly: true,
		Path:     "/",
	})

	response.JSON(w, http.StatusCreated, registerResponse{ID: newUser.ID.String()})
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

	token, err := h.jwtManager.Generate(user.ID.String())
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorResponse{Error: "failed to generate token"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(h.jwtManager.GetTTL()),
		HttpOnly: true,
		Path:     "/",
	})

	response.JSON(w, http.StatusOK, loginResponse{ID: user.ID.String()})
}

func (h *Handlers) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	response.JSON(w, http.StatusOK, logoutResponse{Status: "ok"})
}
