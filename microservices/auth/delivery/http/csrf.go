package http

//go:generate easyjson $GOFILE

import (
	"net/http"
	"spotify/internal/middleware"
	"spotify/pkg/response"
)

//easyjson:json
type csrfResponse struct {
	Token string `json:"csrf_token"`
}

func (h *Handler) GetCSRFToken(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.GetCSRFToken"
	log := middleware.LoggerFromContext(r.Context())

	claims, err := middleware.ClaimsFromContext(r.Context())
	if err != nil {
		log.Errorf("[%s]: %v", op, err)
		response.UnauthorizedJSON(w)
		return
	}

	token, err := h.csrfManager.Generate(claims.UserID, claims.SessionID)
	if err != nil {
		log.Errorf("[%s]: failed to generate csrf token: %v", op, err)
		response.InternalErrorJSON(w)
		return
	}

	log.Debugf("[%s]: successfully generated csrf token for user %s", op, claims.UserID)

	response.JSON(w, http.StatusOK, csrfResponse{Token: token})
}
