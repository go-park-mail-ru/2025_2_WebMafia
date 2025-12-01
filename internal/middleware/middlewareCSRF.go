package middleware

import (
	"net/http"
	"spotify/pkg/response"
)

const CSRFHeader = "X-CSRF-Token"

type CSRFManager interface {
	Check(userID, sessionID, clientToken string) (bool, error)
}

type CSRF struct {
	csrfManager CSRFManager
}

func NewCSRFMiddleware(csrfManager CSRFManager) *CSRF {
	return &CSRF{csrfManager: csrfManager}
}

func (m *CSRF) CSRFMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const op = "middleware.CSRF"
		log := LoggerFromContext(r.Context())

		claims, err := ClaimsFromContext(r.Context())
		if err != nil {
			log.Errorf("[%s]: %v", op, err)
			response.ForbiddenJSON(w)
			return
		}

		csrfToken := r.Header.Get(CSRFHeader)
		if csrfToken == "" {
			log.Warnf("[%s]: missing CSRF token in header", op)
			response.ForbiddenJSON(w)
			return
		}

		isValid, err := m.csrfManager.Check(claims.UserID, claims.SessionID, csrfToken)
		if err != nil || !isValid {
			log.Warnf("[%s]: invalid CSRF token: %v", op, err)
			response.ForbiddenJSON(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
