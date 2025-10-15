package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"spotify/internal/middleware"
)

func (h *Handler) RegisterRouter(r *mux.Router, auth *middleware.Auth) {

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/register", h.Register).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/login", h.Login).Methods(http.MethodPost, http.MethodOptions)

	protected := api.PathPrefix("").Subrouter()
	protected.Use(auth.AuthMiddleware)

	protected.HandleFunc("/logout", h.Logout).Methods(http.MethodPost, http.MethodOptions)

}
