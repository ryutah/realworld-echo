// Package auth provides ...
package auth

import "context"

type AuthToken string

func NewAuthToken(s string) AuthToken {
	return AuthToken(s)
}

type tokenKey struct{}

func ContextWithAuthToken(ctx context.Context, token AuthToken) context.Context {
	return context.WithValue(ctx, tokenKey{}, token)
}

func AuthTokenFromContext(ctx context.Context) (AuthToken, bool) {
	val := ctx.Value(tokenKey{})
	if token, ok := val.(AuthToken); ok {
		return token, true
	}
	return "", false
}
