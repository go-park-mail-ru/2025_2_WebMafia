package middleware

import (
	"bufio"
	"context"
	"errors"
	"net"
	"net/http"
	"spotify/pkg/logger"
	"time"

	"github.com/google/uuid"
)

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

			ctx := ContextWithLogger(context.WithValue(r.Context(), requestIDKey, requestID), contextLogger)

			rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			start := time.Now()

			next.ServeHTTP(rw, r.WithContext(ctx))

			duration := time.Since(start)

			contextLogger.Infof(
				"Request Completed: Method=%s Path=%s StatusCode=%d Duration=%vms RemoteAddr=%s",
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

func (rw *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := rw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("response writer does not support hijacking")
	}
	return hijacker.Hijack()
}

func LoggerFromContext(ctx context.Context) logger.Logger {
	log, ok := ctx.Value(loggerKey).(logger.Logger)
	if !ok {
		l, _ := logger.New("error", logger.ModeDev)
		return l
	}
	return log
}

func ContextWithLogger(ctx context.Context, l logger.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, l)
}
