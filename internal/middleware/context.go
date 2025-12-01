package middleware

import (
	"context"
	"fmt"

	"spotify/pkg/jwtmanager"
)

type ctxKey string

const (
	claimsKey ctxKey = "claims"
)

func ContextWithClaims(ctx context.Context, claims *jwtmanager.Claims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
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
