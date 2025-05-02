package context

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

func GetTraceIDFrom(ctx context.Context) string {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().HasTraceID() {
		return span.SpanContext().TraceID().String()
	}
	return ""
}
