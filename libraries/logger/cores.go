package logger

import (
	"context"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const defaultLevel = zap.InfoLevel

func init() {
	zap.ReplaceGlobals(initLogger())
}

func Logger() *zap.SugaredLogger {
	return zap.S()
}

func WithKV(ctx context.Context, keysAndValues ...interface{}) context.Context {
	l := FromContext(ctx).With(keysAndValues...)
	return ToContext(ctx, l.Desugar())
}

func initLogger() *zap.Logger {
	level := getLevel()

	config := zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			zapcore.RFC3339NanoTimeEncoder(time.UTC(), encoder)
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(os.Stderr), level)

	return zap.New(core)
}

func getLevel() zap.AtomicLevel {
	var level zapcore.Level
	switch strings.ToLower(os.Getenv("LOGGER_LEVEL")) {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = defaultLevel
	}

	return zap.NewAtomicLevelAt(level)
}
