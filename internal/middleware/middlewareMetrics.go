package middleware

import (
	"net/http"
	"spotify/internal/metrics"
	"time"

	"github.com/gorilla/mux"
)

func MetricsMiddleware(m *metrics.Metrics) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			start := time.Now()

			next.ServeHTTP(rw, r)

			duration := time.Since(start)
			var path string
			route := mux.CurrentRoute(r)
			if route != nil {
				var err error
				path, err = route.GetPathTemplate()
				if err != nil {
					LoggerFromContext(r.Context()).Warnf("Failed to get path template: %v", err)
					path = "unknown"
				}
			} else {
				path = "unknown"
			}

			m.IncHttpRequestsTotal(rw.statusCode, r.Method, path)
			m.ObserveHttpRequestDuration(r.Method, path, duration)
		})
	}
}
