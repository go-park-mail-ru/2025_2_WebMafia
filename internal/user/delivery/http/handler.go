package http

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"spotify/internal/user/model"
	"spotify/internal/user/service"
)

type IService interface {
	Register(ctx context.Context, login, email, password string) (model.User, error)
}
type Handler struct {
	svc IService
}

type registerRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerResponse struct {
	ID string `json:"id"`
}

func NewHandler(svc IService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req registerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[Register] invalid body: %v", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.svc.Register(r.Context(), req.Login, req.Email, req.Password)

	if err != nil {
		var svcErr *service.IsServiceError
		if errors.As(err, &svcErr) {
			log.Printf("[Register] %s: %s", svcErr.Type, svcErr.Message)
			switch svcErr.Type {
			case service.ErrValidation:
				http.Error(w, svcErr.Message, http.StatusBadRequest)
			case service.ErrConflict:
				http.Error(w, svcErr.Message, http.StatusConflict)
			default:
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		} else {
			log.Printf("[Register] internal: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	resp := registerResponse{ID: user.ID.String()}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
