package router

import (
	"spotify/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter(h *handler.Handlers) *mux.Router {
	r := mux.NewRouter()
	r.Use(handler.CORS)
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/home", h.HomeHandler).Methods("GET")
	api.HandleFunc("/tracks", h.GetAllTracksHandler).Methods("GET")
	api.HandleFunc("/tracks/{id}", h.GetTrackByIDHandler).Methods("GET")
	api.HandleFunc("/artists", h.GetAllArtistsHandler).Methods("GET")
	api.HandleFunc("/artists/{id}", h.GetArtistByIDHandler).Methods("GET")
	api.HandleFunc("/albums", h.GetAllAlbumsHandler).Methods("GET")
	api.HandleFunc("/albums/{id}", h.GetAlbumByIDHandler).Methods("GET")
	api.HandleFunc("/registration", h.RegisterHandler).Methods("POST")
	api.HandleFunc("/autorization", h.AutorizationHandler).Methods("POST")
	return r
}
