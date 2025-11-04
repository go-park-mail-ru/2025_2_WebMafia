package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(public, protected, csrfProtected *mux.Router) {
	public.HandleFunc("/register", h.Register).Methods(http.MethodPost, http.MethodOptions)
	public.HandleFunc("/login", h.Login).Methods(http.MethodPost, http.MethodOptions)

	protected.HandleFunc("/csrf-token", h.GetCSRFToken).Methods(http.MethodGet, http.MethodOptions)

	csrfProtected.HandleFunc("/logout", h.Logout).Methods(http.MethodPost, http.MethodOptions)
	csrfProtected.HandleFunc("/avatar", h.UploadAvatar).Methods(http.MethodPost, http.MethodOptions)
	csrfProtected.HandleFunc("/avatar", h.DeleteAvatar).Methods(http.MethodDelete, http.MethodOptions)

	protected.HandleFunc("/profile", h.UpdateProfile).Methods(http.MethodPut, http.MethodOptions)
	protected.HandleFunc("/me", h.GetProfile).Methods(http.MethodGet, http.MethodOptions)
}
