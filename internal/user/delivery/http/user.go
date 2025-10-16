package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"spotify/internal/user/dto"
	"spotify/pkg/response"
	"strings"
	"time"
)

const sessionTokenCookie = "session_token"

type registerRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
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

type registerResponse struct {
	ID string `json:"id"`
}
type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
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

type loginResponse struct {
	ID string `json:"id"`
}

type logoutResponse struct {
	Status string `json:"status"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	const op = "[Register] "
	defer r.Body.Close()

	var req registerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("%s invalid body: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	if err := req.validate(); err != nil {
		log.Printf("%s validation error: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	user, err := h.svc.Register(r.Context(), dto.RegisterRequest{
		Login:    req.Login,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		log.Printf("%s service error: %v", op, err)
		handleServiceError(w, err)
		return
	}

	token, err := h.jwtManager.Generate(user.ID)
	if err != nil {
		log.Printf("%s ERROR: failed to generate token", op)
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

	response.JSON(w, http.StatusCreated, registerResponse{ID: user.ID})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	const op = "[Login] "
	defer r.Body.Close()

	var req registerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("%s invalid body: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	if err := req.validate(); err != nil {
		log.Printf("%s validation error: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	user, err := h.svc.Login(r.Context(), dto.LoginRequest{
		Login:    req.Login,
		Password: req.Password,
	})

	if err != nil {
		log.Printf("%s service error:: %v", op, err)
		handleServiceError(w, err)
		return
	}

	token, err := h.jwtManager.Generate(user.ID)
	if err != nil {
		log.Printf("%s ERROR: failed to generate token", op)
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

	response.JSON(w, http.StatusOK, loginResponse{ID: user.ID})
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     sessionTokenCookie,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	response.JSON(w, http.StatusOK, logoutResponse{Status: "ok"})
}
