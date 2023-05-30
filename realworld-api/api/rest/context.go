package rest

import (
	"context"

	"github.com/labstack/echo/v4"
)

type echoCtxKey struct{}

func newContext(c echo.Context) context.Context {
	ctx := c.Request().Context()
	return context.WithValue(ctx, echoCtxKey{}, c)
}

func echoContextFromContext(ctx context.Context) echo.Context {
	val := ctx.Value(echoCtxKey{})
	if ec, ok := val.(echo.Context); ok {
		return ec
	}
	panic("context has no echo.Context")
}
