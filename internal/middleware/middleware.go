package middleware

// import (
// 	"context"
// 	"errors"
// 	"log"
// 	"net/http"
// 	"spotify/pkg/response"
// )

// type ctxKeyUserID string

// const userIDKey ctxKeyUserID = "userID"

// func (h *Handlers) AuthMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cookie, err := r.Cookie(sessionTokenCookie)
// 		if err != nil {
// 			if errors.Is(err, http.ErrNoCookie) {
// 				log.Printf("ERROR: no token provided")
// 				response.UnauthorizedJSON(w)
// 				return
// 			}
// 			log.Printf("ERROR: bad request")
// 			response.BadRequestJSON(w)
// 			return
// 		}

// 		tokenString := cookie.Value

// 		claims, err := h.jwtManager.Validate(tokenString)
// 		if err != nil {
// 			log.Printf("ERROR: invalid token")
// 			response.UnauthorizedJSON(w)
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), userIDKey, claims)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
