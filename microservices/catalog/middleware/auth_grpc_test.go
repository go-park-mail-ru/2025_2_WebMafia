package middleware

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc"

	"spotify/internal/middleware"
	"spotify/microservices/catalog/mocks"
	pb "spotify/proto/auth"
)

func TestAuthGrpcMiddleware_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthClient := mocks.NewMockAuthServiceClient(ctrl)
	mw := NewAuthGrpcMiddleware(mockAuthClient)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, ok := middleware.GetUserID(r.Context())
		if ok && userID == "user123" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusTeapot)
		}
	})

	t.Run("no cookie", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()

		mw.Handle(nextHandler).ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("grpc validate error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.AddCookie(&http.Cookie{Name: sessionTokenCookie, Value: "token"})
		rr := httptest.NewRecorder()

		mockAuthClient.EXPECT().ValidateToken(gomock.Any(), &pb.ValidateTokenRequest{Token: "token"}).
			Return(nil, errors.New("rpc error"))

		mw.Handle(nextHandler).ServeHTTP(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("token invalid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.AddCookie(&http.Cookie{Name: sessionTokenCookie, Value: "token"})
		rr := httptest.NewRecorder()

		mockAuthClient.EXPECT().ValidateToken(gomock.Any(), &pb.ValidateTokenRequest{Token: "token"}).
			Return(&pb.ValidateTokenResponse{IsValid: false}, nil)

		mw.Handle(nextHandler).ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("success GET request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.AddCookie(&http.Cookie{Name: sessionTokenCookie, Value: "token"})
		rr := httptest.NewRecorder()

		mockAuthClient.EXPECT().ValidateToken(gomock.Any(), &pb.ValidateTokenRequest{Token: "token"}).
			Return(&pb.ValidateTokenResponse{IsValid: true, UserId: "user123", SessionId: "sess1"}, nil)

		mw.Handle(nextHandler).ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("POST missing CSRF header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.AddCookie(&http.Cookie{Name: sessionTokenCookie, Value: "token"})
		rr := httptest.NewRecorder()

		mockAuthClient.EXPECT().ValidateToken(gomock.Any(), gomock.Any()).
			Return(&pb.ValidateTokenResponse{IsValid: true, UserId: "user123", SessionId: "sess1"}, nil)

		mw.Handle(nextHandler).ServeHTTP(rr, req)
		assert.Equal(t, http.StatusForbidden, rr.Code)
	})

	t.Run("POST csrf check error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.AddCookie(&http.Cookie{Name: sessionTokenCookie, Value: "token"})
		req.Header.Set(csrfHeader, "csrf_token")
		rr := httptest.NewRecorder()

		mockAuthClient.EXPECT().ValidateToken(gomock.Any(), gomock.Any()).
			Return(&pb.ValidateTokenResponse{IsValid: true, UserId: "user123", SessionId: "sess1"}, nil)

		mockAuthClient.EXPECT().CheckCSRF(gomock.Any(), &pb.CheckCSRFRequest{
			UserId:    "user123",
			SessionId: "sess1",
			CsrfToken: "csrf_token",
		}, gomock.Any()).Return(nil, errors.New("rpc error"))

		mw.Handle(nextHandler).ServeHTTP(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("POST csrf invalid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.AddCookie(&http.Cookie{Name: sessionTokenCookie, Value: "token"})
		req.Header.Set(csrfHeader, "csrf_token")
		rr := httptest.NewRecorder()

		mockAuthClient.EXPECT().ValidateToken(gomock.Any(), gomock.Any()).
			Return(&pb.ValidateTokenResponse{IsValid: true, UserId: "user123", SessionId: "sess1"}, nil)

		mockAuthClient.EXPECT().CheckCSRF(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&pb.CheckCSRFResponse{IsValid: false}, nil)

		mw.Handle(nextHandler).ServeHTTP(rr, req)
		assert.Equal(t, http.StatusForbidden, rr.Code)
	})

	t.Run("POST success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.AddCookie(&http.Cookie{Name: sessionTokenCookie, Value: "token"})
		req.Header.Set(csrfHeader, "csrf_token")
		rr := httptest.NewRecorder()

		mockAuthClient.EXPECT().ValidateToken(gomock.Any(), gomock.Any()).
			Return(&pb.ValidateTokenResponse{IsValid: true, UserId: "user123", SessionId: "sess1"}, nil)

		mockAuthClient.EXPECT().CheckCSRF(gomock.Any(), &pb.CheckCSRFRequest{
			UserId:    "user123",
			SessionId: "sess1",
			CsrfToken: "csrf_token",
		}, gomock.Any()).Return(&pb.CheckCSRFResponse{IsValid: true}, nil)

		mw.Handle(nextHandler).ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

type MockAuthClient struct {
	pb.AuthServiceClient
}

func (m *MockAuthClient) ValidateToken(ctx context.Context, in *pb.ValidateTokenRequest, opts ...grpc.CallOption) (*pb.ValidateTokenResponse, error) {
	return &pb.ValidateTokenResponse{IsValid: true}, nil
}
