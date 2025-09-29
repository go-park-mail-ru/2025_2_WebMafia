package router

import (
	"net/http"
	"spotify/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter(h *handler.Handlers, corsConfig handler.CORSConfig) *mux.Router {
	r := mux.NewRouter()

	r.Use(handler.CORS(corsConfig))

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/register", h.RegisterHandler).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/login", h.LoginHandler).Methods(http.MethodPost, http.MethodOptions)

	protected := api.PathPrefix("").Subrouter()
	protected.Use(h.AuthMiddleware)

	protected.HandleFunc("/logout", h.LogoutHandler).Methods(http.MethodPost, http.MethodOptions)
	protected.HandleFunc("/home", h.HomeHandler).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/tracks", h.GetAllTracksHandler).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/tracks/{id}", h.GetTrackByIDHandler).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/artists", h.GetAllArtistsHandler).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/artists/{id}", h.GetArtistByIDHandler).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/albums", h.GetAllAlbumsHandler).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/albums/{id}", h.GetAlbumByIDHandler).Methods(http.MethodGet, http.MethodOptions)
	return r
}
