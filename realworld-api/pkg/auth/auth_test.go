package auth_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	. "github.com/ryutah/realworld-echo/realworld-api/pkg/auth"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtesting"
)

func TestAuthToken(t *testing.T) {
	ctx := context.Background()

	token := NewAuthToken("token")
	newCtx := ContextWithAuthToken(ctx, token)
	got, ok := AuthTokenFromContext(newCtx)

	if !ok {
		t.Error("should be return true")
	}
	if diff := cmp.Diff(token, got); diff != "" {
		xtesting.PrintDiff(t, "AuthTokenFromContext", diff)
	}
}

func TestAuthTokenWihtoutSetToken(t *testing.T) {
	ctx := context.Background()

	got, ok := AuthTokenFromContext(ctx)

	if ok {
		t.Error("should be return false")
	}
	if diff := cmp.Diff(NewAuthToken(""), got); diff != "" {
		xtesting.PrintDiff(t, "AuthTokenFromContext", diff)
	}
}
