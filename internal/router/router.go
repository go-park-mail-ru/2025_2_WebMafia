package router

import (
	albumHandler "spotify/internal/album/delivery/http"
	artistHandler "spotify/internal/artist/delivery/http"
	"spotify/internal/middleware"
	trackHandler "spotify/internal/track/delivery/http"

	"github.com/gorilla/mux"
)

func NewRouter(
	trackHandlers *trackHandler.Handler,
	artistHandlers *artistHandler.Handler,
	albumHandlers *albumHandler.Handler,
	corsConfig middleware.CORSConfig,
) *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.CORS(corsConfig))

	api := r.PathPrefix("/api/v1").Subrouter()

	protected := api.PathPrefix("").Subrouter()

	trackHandlers.RegisterRoutes(protected)
	artistHandlers.RegisterRoutes(protected)
	albumHandlers.RegisterRoutes(protected)

	return r
}
