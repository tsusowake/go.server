package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/morikuni/failure/v2"
)

func RecoverMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					// ErrAbortHandler の場合はそのまま panic させる
					if r == http.ErrAbortHandler {
						panic(r)
					}

					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}

					wrapped := failure.Wrap(err, failure.Context{
						"recovered": "panic occurred",
					})
					ctx.Echo().HTTPErrorHandler(wrapped, ctx)
				}
			}()
			return next(ctx)
		}
	}
}
