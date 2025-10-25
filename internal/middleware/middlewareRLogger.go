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
	requestIDKey ctxKey = "requestID"
	loggerKey    ctxKey = "logger"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func RequestLoggerMiddleware(log logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			requestID := uuid.New().String()
			contextLogger := log.With("request_id", requestID)

			ctx := r.Context()
			ctx = context.WithValue(ctx, requestIDKey, requestID)
			ctx = context.WithValue(ctx, loggerKey, contextLogger)

			rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			start := time.Now()

			next.ServeHTTP(rw, r.WithContext(ctx))

			duration := time.Since(start)

			contextLogger.Infof(
				"Request Completed: Method=%s Path=%s StatusCode=%d Duration=%v RemoteAddr=%s",
				r.Method,
				r.URL.Path,
				rw.statusCode,
				duration.Milliseconds(),
				r.RemoteAddr,
			)
		})
	}
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func LoggerFromContext(ctx context.Context) logger.Logger {
	log, ok := ctx.Value(loggerKey).(logger.Logger)
	if !ok {
		l, _ := logger.New("error", logger.ModeDev)
		return l
	}
	return log
}
