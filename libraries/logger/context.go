package logger

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type ctxLogger struct{}

// ToContext создаёт контекст с переданным логгером
func ToContext(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// FromContext возвращает логгер из контекста
func FromContext(ctx context.Context) *zap.SugaredLogger {
	l := getLogger(ctx)

	//span := trace.SpanContextFromContext(ctx)
	//
	//return loggerWithSpanContext(l, span)
	return l
}

func getLogger(ctx context.Context) *zap.SugaredLogger {
	if l, ok := ctx.Value(ctxLogger{}).(*zap.SugaredLogger); ok {
		return l
	}
	return zap.S()
}

func loggerWithSpanContext(l *zap.SugaredLogger, sc trace.SpanContext) *zap.SugaredLogger {
	return l.Desugar().With(
		zap.String("trace_id", sc.TraceID().String()),
		zap.String("span_id", sc.SpanID().String()),
	).Sugar()
}
