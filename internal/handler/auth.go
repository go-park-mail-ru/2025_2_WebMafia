package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"spotify/internal/store"
	"spotify/internal/user/model"
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
	if !strings.Contains(i.Email, "@") {
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
		log.Printf("ERROR: invalid request body")
		response.BadRequestJSON(w)
		return
	}

	if err := req.validate(); err != nil {
		log.Printf("ERROR: %v", err)
		response.BadRequestJSON(w)
		return
	}

	_, err := h.store.GetUserByEmail(r.Context(), req.Email)
	if !errors.Is(err, store.ErrUserNotFound) {
		log.Printf("ERROR: user with this email already exists")
		response.ConflictJSON(w)
		return
	}

	_, err = h.store.GetUserByLogin(r.Context(), req.Login)
	if !errors.Is(err, store.ErrUserNotFound) {
		log.Printf("ERROR: user with this login already exists")
		response.ConflictJSON(w)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("ERROR: failed to hash password")
		response.InternalErrorJSON(w)
		return
	}

	newUser, err := h.store.CreateUser(r.Context(), model.User{
		Login:        req.Login,
		PasswordHash: string(hashedPassword),
		Email:        req.Email,
	})
	if err != nil {
		if errors.Is(err, store.ErrUserAlreadyExists) {
			log.Printf("ERROR: %v", err)
			response.ConflictJSON(w)
			return
		}
		log.Printf("ERROR: failed to create user")
		response.InternalErrorJSON(w)
		return
	}

	token, err := h.jwtManager.Generate(newUser)
	if err != nil {
		log.Printf("ERROR: failed to generate token")
		response.InternalErrorJSON(w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     sessionTokenCookie,
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
		log.Printf("ERROR: invalid request body")
		response.BadRequestJSON(w)
		return
	}
	if err := req.validate(); err != nil {
		log.Printf("ERROR: %v", err)
		response.BadRequestJSON(w)
		return
	}

	user, err := h.store.GetUserByLogin(r.Context(), req.Login)
	if err != nil {
		if errors.Is(err, store.ErrUserNotFound) {
			log.Printf("ERROR: invalid login or password")
			response.UnauthorizedJSON(w)
			return
		}
		log.Printf("ERROR: internal server error")
		response.InternalErrorJSON(w)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		log.Printf("ERROR: invalid login or password")
		response.UnauthorizedJSON(w)
		return
	}

	token, err := h.jwtManager.Generate(user)
	if err != nil {
		log.Printf("ERROR: failed to generate token")
		response.InternalErrorJSON(w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     sessionTokenCookie,
		Value:    token,
		Expires:  time.Now().Add(h.jwtManager.GetTTL()),
		HttpOnly: true,
		Path:     "/",
	})

	response.JSON(w, http.StatusOK, loginResponse{ID: user.ID.String()})
}

func (h *Handlers) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     sessionTokenCookie,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	response.JSON(w, http.StatusOK, logoutResponse{Status: "ok"})
}
