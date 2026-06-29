package middleware

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tsusowake/go.server/pkg/context"
)

func NewLoggerMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		BeforeNextFunc: func(c echo.Context) {
			slog.LogAttrs(
				c.Request().Context(),
				slog.LevelInfo,
				"request",
				slog.String("method", c.Request().Method),
				slog.String("uri", c.Request().RequestURI),
				slog.String("ip_address", c.RealIP()),
				slog.String("user_agent", c.Request().UserAgent()),
				slog.String("trace_id", context.GetTraceIDFrom(c.Request().Context())),
			)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				slog.LogAttrs(
					c.Request().Context(),
					slog.LevelInfo,
					"response",
					slog.String("method", v.Method),
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("latency", v.Latency.String()),
					slog.String("trace_id", context.GetTraceIDFrom(c.Request().Context())),
				)
			} else {
				slog.LogAttrs(
					c.Request().Context(),
					slog.LevelError,
					"response error",
					slog.String("method", v.Method),
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("latency", v.Latency.String()),
					slog.String("trace_id", context.GetTraceIDFrom(c.Request().Context())),
				)
			}
			return nil
		},
	})
}
