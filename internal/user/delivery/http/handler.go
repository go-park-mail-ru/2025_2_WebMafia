package http

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"spotify/internal/user/model"
	"spotify/internal/user/service"
	"spotify/pkg/jwtmanager"
	"spotify/pkg/response"
	"time"
)

const registerLogPref = "[Register] "
const sessionTokenCookie = "session_token"

type IService interface {
	Register(ctx context.Context, login, email, password string) (*model.User, error)
}
type Handler struct {
	svc        IService
	jwtManager *jwtmanager.Manager
}

func NewHandler(svc IService) *Handler {
	return &Handler{svc: svc}
}

type registerRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerResponse struct {
	ID string `json:"id"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req registerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf(registerLogPref+"invalid body: %v", err)
		response.BadRequestJSON(w)
		return
	}

	if len(req.Login) < 5 {
		log.Printf(registerLogPref + "validation error: login too short")
		response.BadRequestJSON(w)
		return
	}
	if len(req.Password) < 8 {
		log.Printf(registerLogPref + "validation error: password too short")
		response.BadRequestJSON(w)
		return
	}
	if !containsAt(req.Email) {
		log.Printf(registerLogPref + "validation error: invalid email format")
		response.BadRequestJSON(w)
		return
	}

	user, err := h.svc.Register(r.Context(), req.Login, req.Email, req.Password)

	if err != nil {
		log.Printf(registerLogPref+"service error: %v", err)
		var svcErr *service.IsServiceError
		if errors.As(err, &svcErr) {
			switch {
			case errors.Is(err, service.ErrValidation):
				response.BadRequestJSON(w)
			case errors.Is(err, service.ErrConflict):
				response.ConflictJSON(w)
			default:
				response.InternalErrorJSON(w)
			}
		} else {
			response.InternalErrorJSON(w)
		}
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

	response.JSON(w, http.StatusCreated, registerResponse{ID: user.ID.String()})
}

func containsAt(email string) bool {
	for _, c := range email {
		if c == '@' {
			return true
		}
	}
	return false
}
