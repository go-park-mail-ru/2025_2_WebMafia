package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	user, err := h.svc.Register(r.Context(), req.Login, req.Email, req.Password)

	if err != nil {
		log.Printf("%s service error: %v", op, err)
		handleServiceError(w, err)
		return
	}

	token, err := h.jwtManager.Generate(user)
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

	response.JSON(w, http.StatusCreated, registerResponse{ID: user.ID.String()})
}
