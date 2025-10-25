package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"spotify/pkg/jwtmanager"
	"spotify/pkg/response"
)

type ctxKeyClaims string

const (
	sessionTokenCookie string       = "session_token"
	claimsKey          ctxKeyClaims = "claims"
)

type Auth struct {
	jwt *jwtmanager.Manager
}

func NewAuthMiddleware(jwt *jwtmanager.Manager) *Auth {
	return &Auth{jwt: jwt}
}

func (a *Auth) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const op = "AuthMiddleware"
		log := LoggerFromContext(r.Context())

		cookie, err := r.Cookie(sessionTokenCookie)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				log.Warnf("[%s]: No token provided", op)
				response.UnauthorizedJSON(w)
				return
			}
			log.Errorf("[%s]: Error getting cookie: %v", op, err)
			response.BadRequestJSON(w)
			return
		}

		claims, err := a.jwt.Validate(cookie.Value)
		if err != nil {
			log.Warnf("[%s]: Invalid token: %v", op, err)
			response.UnauthorizedJSON(w)
			return
		}

		ctx := context.WithValue(r.Context(), claimsKey, claims)
		ctxLogger := log.With("user_id", claims.UserID)
		ctx = context.WithValue(ctx, loggerKey, ctxLogger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ClaimsFromContext(ctx context.Context) (*jwtmanager.Claims, error) {
	claims, ok := ctx.Value(claimsKey).(*jwtmanager.Claims)
	if !ok {
		return nil, fmt.Errorf("no claims found in context")
	}
	return claims, nil
}

func GetUserID(ctx context.Context) (string, bool) {
	claims, err := ClaimsFromContext(ctx)
	if err != nil {
		return "", false
	}
	return claims.UserID, true
}
