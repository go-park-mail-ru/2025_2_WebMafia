package middleware

import (
	"net/http"
	"spotify/internal/middleware"
	"spotify/pkg/jwtmanager"
	"spotify/pkg/response"

	pb "spotify/proto/auth"
)

const (
	sessionTokenCookie = "session_token"
	csrfHeader         = "X-CSRF-Token"
)

//go:generate mockgen -destination=../mocks/auth_client_mock.go -package=mocks spotify/proto/auth AuthServiceClient
type AuthGrpcMiddleware struct {
	authClient pb.AuthServiceClient
}

func NewAuthGrpcMiddleware(authClient pb.AuthServiceClient) *AuthGrpcMiddleware {
	return &AuthGrpcMiddleware{
		authClient: authClient,
	}
}

func (m *AuthGrpcMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const op = "middleware.AuthGrpc"

		log := middleware.LoggerFromContext(r.Context())

		cookie, err := r.Cookie(sessionTokenCookie)
		if err != nil {
			log.Warnf("[%s]: no session cookie", op)
			response.UnauthorizedJSON(w)
			return
		}

		validateResp, err := m.authClient.ValidateToken(r.Context(), &pb.ValidateTokenRequest{
			Token: cookie.Value,
		})
		if err != nil {
			log.Errorf("[%s]: grpc call ValidateToken failed: %v", op, err)
			response.InternalErrorJSON(w)
			return
		}
		if !validateResp.IsValid {
			log.Warnf("[%s]: token invalid", op)
			response.UnauthorizedJSON(w)
			return
		}

		claims := &jwtmanager.Claims{
			UserID:    validateResp.GetUserId(),
			SessionID: validateResp.GetSessionId(),
		}

		ctx := middleware.ContextWithClaims(r.Context(), claims)

		if r.Method != http.MethodGet && r.Method != http.MethodHead && r.Method != http.MethodOptions {
			csrfToken := r.Header.Get(csrfHeader)
			if csrfToken == "" {
				log.Warnf("[%s]: missing csrf header", op)
				response.ForbiddenJSON(w)
				return
			}

			csrfResp, err := m.authClient.CheckCSRF(r.Context(), &pb.CheckCSRFRequest{
				UserId:    validateResp.UserId,
				SessionId: validateResp.SessionId,
				CsrfToken: csrfToken,
			})
			if err != nil {
				log.Errorf("[%s]: grpc call CheckCSRF failed: %v", op, err)
				response.InternalErrorJSON(w)
				return
			}
			if !csrfResp.IsValid {
				log.Warnf("[%s]: csrf invalid", op)
				response.ForbiddenJSON(w)
				return
			}
		}

		ctxLogger := log.With("user_id", claims.UserID)
		ctx = middleware.ContextWithLogger(ctx, ctxLogger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
