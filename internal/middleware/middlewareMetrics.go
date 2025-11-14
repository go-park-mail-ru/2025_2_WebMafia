package middleware

import (
	"net/http"
	"strconv"
	"time"

	"spotify/internal/metrics"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

func MetricsMiddleware(m *metrics.Metrics) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			start := time.Now()

			next.ServeHTTP(rw, r)

			duration := time.Since(start)
			route := mux.CurrentRoute(r)
			path, _ := route.GetPathTemplate()

			m.RequestsTotal.With(prometheus.Labels{
				"code":   strconv.Itoa(rw.statusCode),
				"method": r.Method,
				"path":   path,
			}).Inc()

			m.RequestDuration.With(prometheus.Labels{
				"method": r.Method,
				"path":   path,
			}).Observe(duration.Seconds())
		})
	}
}
