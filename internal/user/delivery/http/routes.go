package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(public, protected, csrfProtected *mux.Router) {
	public.HandleFunc("/register", h.Register).Methods(http.MethodPost, http.MethodOptions)
	public.HandleFunc("/login", h.Login).Methods(http.MethodPost, http.MethodOptions)

	csrfProtected.HandleFunc("/logout", h.Logout).Methods(http.MethodPost, http.MethodOptions)
	protected.HandleFunc("/csrf-token", h.GetCSRFToken).Methods(http.MethodGet, http.MethodOptions)
}
