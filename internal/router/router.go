package router

import (
	albumDelivery "spotify/internal/album/delivery/http"
	artistDelivery "spotify/internal/artist/delivery/http"
	"spotify/internal/middleware"
	trackDelivery "spotify/internal/track/delivery/http"
	userDelivery "spotify/internal/user/delivery/http"

	"github.com/gorilla/mux"
)

type AppHandlers struct {
	UserHandler   *userDelivery.Handler
	ArtistHandler *artistDelivery.Handler
	AlbumHandler  *albumDelivery.Handler
	TrackHandler  *trackDelivery.Handler
}

func NewRouter(handlers AppHandlers, auth *middleware.Auth, cfg middleware.CORSConfig) *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.CORS(cfg))

	api := r.PathPrefix("/api/v1").Subrouter()

	public := api.PathPrefix("").Subrouter()

	protected := api.PathPrefix("").Subrouter()
	protected.Use(auth.AuthMiddleware)

	handlers.UserHandler.RegisterRoutes(public, protected)

	handlers.ArtistHandler.RegisterRoutes(public)
	handlers.AlbumHandler.RegisterRoutes(public)
	handlers.TrackHandler.RegisterRoutes(public)

	return r
}
