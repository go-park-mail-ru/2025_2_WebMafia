package handler

import (
	"context"
	"net/http"
	"spotify/pkg/response"
	"strings"
)

func (h *Handlers) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.JSON(w, http.StatusUnauthorized, response.ErrorResponse{Error: "header required"})
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.JSON(w, http.StatusUnauthorized, response.ErrorResponse{Error: "invalid header "})
			return
		}
		tokenString := parts[1]
		userID, err := h.validateToken(tokenString)
		if err != nil {
			response.JSON(w, http.StatusUnauthorized, response.ErrorResponse{Error: "invalid token"})
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
