package middleware

import (
	"context"
	"net/http"
	"time"

	"spotify/pkg/logger"

	"github.com/google/uuid"
)

type ctxKey string

const (
	RequestIDKey ctxKey = "requestID"
	Logger       ctxKey = "logger"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func RequestLogger(log logger.ILogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			requestID := uuid.New().String()
			contextLogger := log.With("request_id", requestID)

			ctx := r.Context()
			ctx = context.WithValue(ctx, RequestIDKey, requestID)
			ctx = context.WithValue(ctx, Logger, contextLogger)

			rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			start := time.Now()

			next.ServeHTTP(rw, r.WithContext(ctx))

			duration := time.Since(start)

			contextLogger.Infow("request completed",
				"method", r.Method,
				"path", r.URL.Path,
				"status_code", rw.statusCode,
				"duration_ms", duration.Milliseconds(),
				"remote_addr", r.RemoteAddr,
			)
		})
	}
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func LoggerFromContext(ctx context.Context) logger.ILogger {
	log, ok := ctx.Value(Logger).(logger.ILogger)
	if !ok {
		l, _ := logger.New("error", logger.ModeDev)
		return l
	}
	return log
}
