package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (h Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.Register).Methods(http.MethodPost, http.MethodOptions)
}
