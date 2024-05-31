package context

import "context"

type contextKey string

const (
	appLoggerKey contextKey = "app_logger"
)

func ContextWithAppLogger(ctx context.Context)
