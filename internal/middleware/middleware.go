package middleware

import (
	"context"
	"errors"
	"net/http"
	"spotify/pkg/jwtmanager"
	"spotify/pkg/response"
)

type ctxKeyUserID string

const (
	userIDKey          ctxKeyUserID = "userID"
	sessionTokenCookie string       = "session_token"
)

type Auth struct {
	jwt *jwtmanager.Manager
}

func NewAuthMiddleware(jwt *jwtmanager.Manager) *Auth {
	return &Auth{jwt: jwt}
}

func (a *Auth) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := LoggerFromContext(r.Context()).With("op", "AuthMiddleware")

		cookie, err := r.Cookie(sessionTokenCookie)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				log.Warnw("No token provided")
				response.UnauthorizedJSON(w)
				return
			}
			log.Errorw("Error getting cookie", "error", err)
			response.BadRequestJSON(w)
			return
		}

		claims, err := a.jwt.Validate(cookie.Value)
		if err != nil {
			log.Warnw("Invalid token", "error", err)
			response.UnauthorizedJSON(w)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)
		ctxLogger := log.With("user_id", claims.UserID)
		ctx = context.WithValue(ctx, loggerKey, ctxLogger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
