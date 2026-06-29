package echoutil

import (
	"context"

	"github.com/labstack/echo/v5"
)

func FromEchoContext(ctx *echo.Context) context.Context {
	return ctx.Request().Context()
}
