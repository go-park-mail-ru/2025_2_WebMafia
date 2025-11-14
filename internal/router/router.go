package router

import (
	albumDelivery "spotify/internal/album/delivery/http"
	artistDelivery "spotify/internal/artist/delivery/http"
	"spotify/internal/metrics"
	"spotify/internal/middleware"
	trackDelivery "spotify/internal/track/delivery/http"
	userDelivery "spotify/internal/user/delivery/http"
	"spotify/pkg/logger"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type AppHandlers struct {
	UserHandler   *userDelivery.Handler
	ArtistHandler *artistDelivery.Handler
	AlbumHandler  *albumDelivery.Handler
	TrackHandler  *trackDelivery.Handler
}

func NewRouter(logger logger.Logger,
	handlers AppHandlers,
	auth *middleware.Auth,
	csrf *middleware.CSRF,
	cfg middleware.CORSConfig,
	appMetrics *metrics.Metrics) *mux.Router {

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.RequestLoggerMiddleware(logger))
	api.Use(middleware.CORS(cfg))
	api.Use(middleware.MetricsMiddleware(appMetrics))

	r.Handle("/metrics", promhttp.Handler()).Methods("GET")

	public := api.PathPrefix("").Subrouter()

	protected := api.PathPrefix("").Subrouter()
	protected.Use(auth.AuthMiddleware)

	csrfProtected := protected.PathPrefix("").Subrouter()
	csrfProtected.Use(csrf.CSRFMiddleware)

	handlers.UserHandler.RegisterRoutes(public, protected, csrfProtected)

	handlers.ArtistHandler.RegisterRoutes(public)
	handlers.AlbumHandler.RegisterRoutes(public)
	handlers.TrackHandler.RegisterRoutes(public, protected, csrfProtected)

	return r
}
