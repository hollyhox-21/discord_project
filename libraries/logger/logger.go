package logger

import (
	"context"
)

func Debugw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	FromContext(ctx).Debugw(msg, keysAndValues...)
}

func Infow(ctx context.Context, msg string, keysAndValues ...interface{}) {
	FromContext(ctx).Infow(msg, keysAndValues...)
}

func Warnw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	FromContext(ctx).Warnw(msg, keysAndValues...)
}

func Errorw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	FromContext(ctx).Errorw(msg, keysAndValues...)
}

func Fatalw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	FromContext(ctx).Fatalw(msg, keysAndValues...)
}

func Debugf(ctx context.Context, msg string, args ...interface{}) {
	FromContext(ctx).Debugf(msg, args...)
}

func Infof(ctx context.Context, msg string, args ...interface{}) {
	FromContext(ctx).Infof(msg, args...)
}

func Warnf(ctx context.Context, msg string, args ...interface{}) {
	FromContext(ctx).Warnf(msg, args...)
}

func Errorf(ctx context.Context, msg string, args ...interface{}) {
	FromContext(ctx).Errorf(msg, args...)
}

func Fatalf(ctx context.Context, msg string, args ...interface{}) {
	FromContext(ctx).Fatalf(msg, args...)
}
