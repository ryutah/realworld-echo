package auth_test

import (
	"context"
	"testing"

	. "github.com/ryutah/realworld-echo/realworld-api/pkg/auth"
	"github.com/stretchr/testify/assert"
)

func TestAuthToken(t *testing.T) {
	ctx := context.Background()

	token := NewAuthToken("token")
	newCtx := ContextWithAuthToken(ctx, token)
	got, ok := AuthTokenFromContext(newCtx)

	if !ok {
		t.Error("should be return true")
	}
	assert.Equal(t, token, got)
}

func TestAuthTokenWihtoutSetToken(t *testing.T) {
	ctx := context.Background()

	got, ok := AuthTokenFromContext(ctx)

	if ok {
		t.Error("should be return false")
	}
	assert.Equal(t, NewAuthToken("") , got)
}
