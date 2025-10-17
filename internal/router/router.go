package router

import (
	albumDelivery "spotify/internal/album/delivery/http"
	artistDelivery "spotify/internal/artist/delivery/http"
	"spotify/internal/middleware"
	trackDelivery "spotify/internal/track/delivery/http"

	"github.com/gorilla/mux"
)

type AppHandlers struct {
	ArtistHandler *artistDelivery.Handler
	AlbumHandler  *albumDelivery.Handler
	TrackHandler  *trackDelivery.Handler
}

func NewRouter(handlers AppHandlers, cfg middleware.CORSConfig) *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.CORS(cfg))

	api := r.PathPrefix("/api/v1").Subrouter()

	handlers.ArtistHandler.RegisterRoutes(api)
	handlers.AlbumHandler.RegisterRoutes(api)
	handlers.TrackHandler.RegisterRoutes(api)

	return r
}
