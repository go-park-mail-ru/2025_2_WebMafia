package middleware

import (
	"context"
	"errors"
	"log"
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
		const op = "[AuthMiddleware] "
		cookie, err := r.Cookie(sessionTokenCookie)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				log.Printf("%s no token provided", op)
				response.UnauthorizedJSON(w)
				return
			}
			log.Printf("%s ERROR: bad request", op)
			response.BadRequestJSON(w)
			return
		}

		claims, err := a.jwt.Validate(cookie.Value)
		if err != nil {
			log.Printf("%s invalid token", op)
			response.UnauthorizedJSON(w)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
