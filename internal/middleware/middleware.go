package middleware

import (
	"errors"
	"net/http"

	"spotify/pkg/jwtmanager"
	"spotify/pkg/response"
)

const (
	sessionTokenCookie string = "session_token"
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

		ctx := ContextWithClaims(r.Context(), claims)
		ctxLogger := log.With("user_id", claims.UserID)
		ctx = ContextWithLogger(ctx, ctxLogger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
