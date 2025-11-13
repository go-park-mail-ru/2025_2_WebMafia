package router

import (
	albumDelivery "spotify/internal/album/delivery/http"
	artistDelivery "spotify/internal/artist/delivery/http"
	"spotify/internal/middleware"
	playlistDelivery "spotify/internal/playlist/delivery/http"
	trackDelivery "spotify/internal/track/delivery/http"
	userDelivery "spotify/internal/user/delivery/http"
	"spotify/pkg/logger"

	"github.com/gorilla/mux"
)

type AppHandlers struct {
	UserHandler     *userDelivery.Handler
	ArtistHandler   *artistDelivery.Handler
	AlbumHandler    *albumDelivery.Handler
	TrackHandler    *trackDelivery.Handler
	PlaylistHandler *playlistDelivery.Handler
}

func NewRouter(logger logger.Logger,
	handlers AppHandlers,
	auth *middleware.Auth,
	csrf *middleware.CSRF,
	cfg middleware.CORSConfig) *mux.Router {

	r := mux.NewRouter()

	r.Use(middleware.RequestLoggerMiddleware(logger))
	r.Use(middleware.CORS(cfg))

	api := r.PathPrefix("/api/v1").Subrouter()

	protected := api.PathPrefix("").Subrouter()
	protected.Use(auth.AuthMiddleware)

	public := api.PathPrefix("").Subrouter()

	csrfProtected := protected.PathPrefix("").Subrouter()
	csrfProtected.Use(csrf.CSRFMiddleware)

	handlers.UserHandler.RegisterRoutes(public, protected, csrfProtected)

	handlers.ArtistHandler.RegisterRoutes(public)
	handlers.AlbumHandler.RegisterRoutes(public)
	handlers.TrackHandler.RegisterRoutes(public, protected, csrfProtected)
	handlers.PlaylistHandler.RegisterRoutes(public, protected)

	return r
}
