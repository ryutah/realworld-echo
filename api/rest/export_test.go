package rest

import (
	"context"

	"github.com/labstack/echo/v4"
)

func NewContext(c echo.Context) context.Context {
	return newContext(c)
}

func EchoContextFromContext(ctx context.Context) echo.Context {
	return echoContextFromContext(ctx)
}
