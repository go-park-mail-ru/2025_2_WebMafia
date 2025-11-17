package metrics

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	RequestsTotal   *prometheus.CounterVec
	RequestDuration *prometheus.HistogramVec
}

func New(serviceName string) *Metrics {
	constLabels := prometheus.Labels{"service": serviceName}

	requestsTotal := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "wavemusic",
			Name:        "http_requests_total",
			Help:        "Total number of HTTP requests.",
			ConstLabels: constLabels,
		},
		[]string{"code", "method", "path"},
	)

	requestDuration := promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace:   "wavemusic",
			Name:        "http_request_duration_seconds",
			Help:        "Duration of HTTP requests in seconds.",
			Buckets:     prometheus.DefBuckets,
			ConstLabels: constLabels,
		},
		[]string{"method", "path"},
	)

	return &Metrics{
		RequestsTotal:   requestsTotal,
		RequestDuration: requestDuration,
	}
}

func (m *Metrics) IncHttpRequestsTotal(statusCode int, method, path string) {
	m.RequestsTotal.With(prometheus.Labels{
		"code":   strconv.Itoa(statusCode),
		"method": method,
		"path":   path,
	}).Inc()
}

func (m *Metrics) ObserveHttpRequestDuration(method, path string, duration time.Duration) {
	m.RequestDuration.With(prometheus.Labels{
		"method": method,
		"path":   path,
	}).Observe(duration.Seconds())
}
