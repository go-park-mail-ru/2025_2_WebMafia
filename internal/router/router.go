package router

import (
	"spotify/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter(h *handler.Handlers) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/registration", h.RegisterHandler).Methods("POST")
	api.HandleFunc("/login", h.LoginHandler).Methods("POST")

	protected := api.PathPrefix("").Subrouter()
	protected.Use(h.AuthMiddleware)

	return r
}
