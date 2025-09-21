package router

import (
	"spotify/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter(h *handler.Handlers) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", h.HomeHandler).Methods("GET")
	r.HandleFunc("/registration/", h.RegisterHandler).Methods("POST")
	r.HandleFunc("/autorization/", h.AutorizationHandler).Methods("POST")
	return r
}
