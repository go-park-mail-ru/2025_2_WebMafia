package middleware

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"spotify/config"
	"spotify/pkg/jwtmanager"
	"spotify/pkg/logger"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthMiddleware(t *testing.T) {
	jwtManager := jwtmanager.NewManager("secret", time.Minute)
	authMiddleware := NewAuthMiddleware(jwtManager)

	mockLogger, _ := logger.New("error", "dev")
	ctxWithLogger := context.WithValue(context.Background(), loggerKey, mockLogger)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := ClaimsFromContext(r.Context())
		require.NoError(t, err)
		assert.NotEmpty(t, claims.UserID)
		w.WriteHeader(http.StatusOK)
	})

	t.Run("success", func(t *testing.T) {
		token, _ := jwtManager.Generate("user123")
		req := httptest.NewRequest(http.MethodGet, "/", nil).WithContext(ctxWithLogger)
		req.AddCookie(&http.Cookie{Name: sessionTokenCookie, Value: token})
		rr := httptest.NewRecorder()

		authMiddleware.AuthMiddleware(handler).ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("no cookie", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil).WithContext(ctxWithLogger)
		rr := httptest.NewRecorder()
		authMiddleware.AuthMiddleware(handler).ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("invalid token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil).WithContext(ctxWithLogger)
		req.AddCookie(&http.Cookie{Name: sessionTokenCookie, Value: "invalid.token.string"})
		rr := httptest.NewRecorder()
		authMiddleware.AuthMiddleware(handler).ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("malformed cookie header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil).WithContext(ctxWithLogger)
		req.Header.Set("Cookie", "invalid-cookie-format")
		rr := httptest.NewRecorder()
		authMiddleware.AuthMiddleware(handler).ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})
}

func TestContextHelpers(t *testing.T) {
	t.Run("GetUserID from context", func(t *testing.T) {
		claims := &jwtmanager.Claims{UserID: "user123"}
		ctx := context.WithValue(context.Background(), claimsKey, claims)

		userID, ok := GetUserID(ctx)
		assert.True(t, ok)
		assert.Equal(t, "user123", userID)
	})

	t.Run("GetUserID from empty context", func(t *testing.T) {
		_, ok := GetUserID(context.Background())
		assert.False(t, ok)
	})

	t.Run("ClaimsFromContext from empty context", func(t *testing.T) {
		_, err := ClaimsFromContext(context.Background())
		assert.Error(t, err)
	})
}

type mockCSRFManager struct {
	checkFunc func(userID, sessionID, clientToken string) (bool, error)
}

func (m *mockCSRFManager) Check(userID, sessionID, clientToken string) (bool, error) {
	if m.checkFunc != nil {
		return m.checkFunc(userID, sessionID, clientToken)
	}
	return false, errors.New("checkFunc not implemented")
}

func TestCSRFMiddleware(t *testing.T) {
	mockLogger, _ := logger.New("error", "dev")
	ctxWithLogger := context.WithValue(context.Background(), loggerKey, mockLogger)
	claims := &jwtmanager.Claims{UserID: "user123", SessionID: "session123"}
	ctxWithClaims := context.WithValue(ctxWithLogger, claimsKey, claims)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	t.Run("success", func(t *testing.T) {
		mockMgr := &mockCSRFManager{
			checkFunc: func(userID, sessionID, clientToken string) (bool, error) {
				assert.Equal(t, claims.UserID, userID)
				assert.Equal(t, claims.SessionID, sessionID)
				assert.Equal(t, "valid-token", clientToken)
				return true, nil
			},
		}
		csrfMiddleware := NewCSRFMiddleware(mockMgr)

		req := httptest.NewRequest(http.MethodPost, "/", nil).WithContext(ctxWithClaims)
		req.Header.Set(CSRFHeader, "valid-token")
		rr := httptest.NewRecorder()

		csrfMiddleware.CSRFMiddleware(handler).ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("fail - no claims in context", func(t *testing.T) {
		csrfMiddleware := NewCSRFMiddleware(&mockCSRFManager{})
		req := httptest.NewRequest(http.MethodPost, "/", nil).WithContext(ctxWithLogger)
		rr := httptest.NewRecorder()
		csrfMiddleware.CSRFMiddleware(handler).ServeHTTP(rr, req)
		assert.Equal(t, http.StatusForbidden, rr.Code)
	})
}

func TestCORSMiddleware(t *testing.T) {
	config := config.CORSConfig{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}
	corsMiddleware := CORS(config)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	t.Run("allowed origin", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Origin", "http://localhost:8080")
		rr := httptest.NewRecorder()
		corsMiddleware(handler).ServeHTTP(rr, req)
		assert.Equal(t, "http://localhost:8080", rr.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "true", rr.Header().Get("Access-Control-Allow-Credentials"))
	})

	t.Run("disallowed origin", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Origin", "http://disallowed.com")
		rr := httptest.NewRecorder()
		corsMiddleware(handler).ServeHTTP(rr, req)
		assert.Empty(t, rr.Header().Get("Access-Control-Allow-Origin"))
	})

	t.Run("options request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodOptions, "/", nil)
		req.Header.Set("Origin", "http://localhost:8080")
		rr := httptest.NewRecorder()
		corsMiddleware(handler).ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "GET, POST", rr.Header().Get("Access-Control-Allow-Methods"))
	})
}

func TestRequestLoggerMiddleware(t *testing.T) {
	mockLogger, _ := logger.New("info", "dev")
	loggerMiddleware := RequestLoggerMiddleware(mockLogger)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxLogger := LoggerFromContext(r.Context())
		require.NotNil(t, ctxLogger)

		requestID, ok := r.Context().Value(requestIDKey).(string)
		require.True(t, ok)
		assert.NotEmpty(t, requestID)

		w.WriteHeader(http.StatusAccepted)
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()
	loggerMiddleware(handler).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusAccepted, rr.Code)
}
