package http

import (
	"net/http"
	"spotify/internal/middleware"
	"spotify/pkg/jwtmanager"
	"spotify/pkg/response"
)

func (h *Handler) GetCSRFToken(w http.ResponseWriter, r *http.Request) {
	const op = "csrf.handler.GetCSRFToken"
	log := middleware.LoggerFromContext(r.Context())

	claims, ok := r.Context().Value(middleware.ClaimsKey).(*jwtmanager.Claims)
	if !ok {
		log.Errorf("[%s] failed to get claims from context", op)
		response.InternalErrorJSON(w)
		return
	}

	token, err := h.csrfManager.Generate(claims.UserID, claims.SessionID)
	if err != nil {
		log.Errorf("[%s] failed to generate csrf token: %v", op, err)
		response.InternalErrorJSON(w)
		return
	}

	log.Debugf("[%s] successfully generated csrf token for user %s", op, claims.UserID)

	response.JSON(w, http.StatusOK, map[string]string{"csrf_token": token})
}
