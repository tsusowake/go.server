package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level zapcore.Level, options ...zap.Option) (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(level)
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}
	config.Sampling = nil
	return config.Build(options...)
}

func FromContext(ctx context.Context) (*zap.Logger, bool) {
	l, ok := ctx.Value("").(*zap.Logger)
	return l, ok
}
