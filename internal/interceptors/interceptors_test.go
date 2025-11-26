package interceptors

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"spotify/internal/middleware"
	"spotify/pkg/logger"
)

type mockLogger struct {
	logger.Logger
	debugFunc func(string, ...interface{})
	infoFunc  func(string, ...interface{})
	errorFunc func(string, ...interface{})
	withFunc  func(...interface{}) logger.Logger
}

func (m *mockLogger) Debugf(t string, args ...interface{}) {
	if m.debugFunc != nil {
		m.debugFunc(t, args...)
	}
}
func (m *mockLogger) Infof(t string, args ...interface{}) {
	if m.infoFunc != nil {
		m.infoFunc(t, args...)
	}
}
func (m *mockLogger) Errorf(t string, args ...interface{}) {
	if m.errorFunc != nil {
		m.errorFunc(t, args...)
	}
}
func (m *mockLogger) With(args ...interface{}) logger.Logger {
	if m.withFunc != nil {
		return m.withFunc(args...)
	}
	return m
}

func TestPanicRecovery(t *testing.T) {
	t.Run("panic recovered", func(t *testing.T) {
		ml := &mockLogger{
			errorFunc: func(s string, i ...interface{}) {
				assert.Contains(t, s, "panic recovered")
			},
		}
		ctx := middleware.ContextWithLogger(context.Background(), ml)

		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			panic("oops")
		}

		_, err := PanicRecovery(ctx, nil, &grpc.UnaryServerInfo{}, handler)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected error")
	})

	t.Run("no panic", func(t *testing.T) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return "ok", nil
		}
		resp, err := PanicRecovery(context.Background(), nil, &grpc.UnaryServerInfo{}, handler)
		assert.NoError(t, err)
		assert.Equal(t, "ok", resp)
	})
}

func TestRequestLogger(t *testing.T) {
	t.Run("success logging", func(t *testing.T) {
		ml := &mockLogger{
			withFunc: func(i ...interface{}) logger.Logger {
				return &mockLogger{
					infoFunc: func(s string, i ...interface{}) {
						assert.Contains(t, s, "gRPC request")
					},
				}
			},
		}

		interceptor := RequestLogger(ml)
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return "ok", nil
		}

		resp, err := interceptor(context.Background(), nil, &grpc.UnaryServerInfo{FullMethod: "/test"}, handler)
		assert.NoError(t, err)
		assert.Equal(t, "ok", resp)
	})

	t.Run("error logging", func(t *testing.T) {
		ml := &mockLogger{
			withFunc: func(i ...interface{}) logger.Logger {
				return &mockLogger{
					infoFunc: func(s string, i ...interface{}) {},
					errorFunc: func(s string, i ...interface{}) {
						assert.Contains(t, s, "failed")
					},
				}
			},
		}

		interceptor := RequestLogger(ml)
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return nil, errors.New("err")
		}

		_, err := interceptor(context.Background(), nil, &grpc.UnaryServerInfo{FullMethod: "/test"}, handler)
		assert.Error(t, err)
	})
}
