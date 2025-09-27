package handler

import (
	"context"
	"net/http"
	"spotify/pkg/response"
)

type ctxKeyUserID string

const userIDKey ctxKeyUserID = "userID"

func (h *Handlers) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response.JSON(w, http.StatusUnauthorized, response.ErrorResponse{Error: "unauthorized: no token provided"})
				return
			}
			response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: "bad request"})
			return
		}

		tokenString := cookie.Value

		userID, err := h.jwtManager.Validate(tokenString)
		if err != nil {
			response.JSON(w, http.StatusUnauthorized, response.ErrorResponse{Error: "invalid token"})
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
