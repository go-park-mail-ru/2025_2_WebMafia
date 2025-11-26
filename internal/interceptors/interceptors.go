package interceptors

import (
	"context"
	"spotify/internal/middleware"
	"spotify/pkg/logger"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func PanicRecovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log := middleware.LoggerFromContext(ctx)
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("gRPC panic recovered: %v", r)
			err = status.Errorf(codes.Internal, "unexpected error")
		}
	}()

	return handler(ctx, req)
}

func RequestLogger(log logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()

		requestID := uuid.New().String()

		contextLogger := log.With("request_id", requestID, "method", info.FullMethod, "protocol", "grpc")

		ctx = middleware.ContextWithLogger(ctx, contextLogger)

		contextLogger.Infof("gRPC request started: %s", info.FullMethod)

		resp, err = handler(ctx, req)

		duration := time.Since(start)

		if err != nil {
			contextLogger.Errorf("gRPC request failed: method=%s duration=%v error=%v", info.FullMethod, duration, err)
		} else {
			contextLogger.Infof("gRPC request completed: method=%s duration=%v", info.FullMethod, duration)
		}

		return resp, err
	}
}
