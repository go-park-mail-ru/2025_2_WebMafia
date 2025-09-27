package router

import (
	"net/http"
	"spotify/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter(h *handler.Handlers) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/register", h.RegisterHandler).Methods(http.MethodPost)
	api.HandleFunc("/login", h.LoginHandler).Methods(http.MethodPost)
	api.HandleFunc("/logout", h.LogoutHandler).Methods(http.MethodPost)

	protected := api.PathPrefix("").Subrouter()
	protected.Use(h.AuthMiddleware)

	return r
}
