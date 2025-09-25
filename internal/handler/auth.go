package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"spotify/internal/service"
	"spotify/pkg/response"
)

type signUpRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpResponse struct {
	ID string `json:"id"`
}

func (i *signUpRequest) validateSignUpRequest() error {
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
	Token string `json:"token"`
}

func (i *loginRequest) validateLoginRequest() error {
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

	if err := req.validateSignUpRequest(); err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return
	}

	newUser, err := h.authService.Register(r.Context(), req.Email, req.Login, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrEmailExists) || errors.Is(err, service.ErrLoginExists) {
			response.JSON(w, http.StatusConflict, response.ErrorResponse{Error: err.Error()})
			return
		}
		response.JSON(w, http.StatusInternalServerError, response.ErrorResponse{Error: "internal server error"})
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
	if err := req.validateLoginRequest(); err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return
	}

	token, err := h.authService.Login(r.Context(), req.Login, req.Password)

	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			response.JSON(w, http.StatusUnauthorized, response.ErrorResponse{Error: err.Error()})
			return
		}
		response.JSON(w, http.StatusInternalServerError, response.ErrorResponse{Error: "internal server error"})
		return
	}

	response.JSON(w, http.StatusOK, loginResponse{Token: token})

}
