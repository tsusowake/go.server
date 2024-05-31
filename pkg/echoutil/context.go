package echoutil

import (
	"context"

	"github.com/labstack/echo/v4"
)

func FromEchoContext(ctx echo.Context) context.Context {
	return ctx.Request().Context()
}
