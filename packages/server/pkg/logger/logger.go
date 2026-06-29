package logger

import (
	"context"
	"io"
	"log/slog"
)

type LogMetadata struct {
	Service string
	Env     string
}

func NewLogger(
	w io.Writer,
	level slog.Level,
	metadata LogMetadata,
) *slog.Logger {
	h := newLogHandler(w, level, metadata)
	return slog.New(h)
}

type LogHandler struct {
	h        slog.Handler
	metadata slog.Attr
}

func newLogHandler(
	w io.Writer,
	level slog.Level,
	metadata LogMetadata,
) slog.Handler {
	jsonHandler := slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
	})
	h := LogHandler{
		h: jsonHandler,
		metadata: slog.Group("metadata",
			slog.String("service", metadata.Service),
			slog.String("env", metadata.Env),
		),
	}
	return &h
}

func (h *LogHandler) Handle(ctx context.Context, r slog.Record) error {
	r.AddAttrs(h.metadata)
	return h.h.Handle(ctx, r)
}

func (h *LogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.h.Enabled(ctx, level)
}
func (h *LogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h.h.WithAttrs(attrs)
}
func (h *LogHandler) WithGroup(g string) slog.Handler {
	return h.h.WithGroup(g)
}
