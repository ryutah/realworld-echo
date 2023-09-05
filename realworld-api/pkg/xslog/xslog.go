package xslog

import (
	"context"
	"log/slog"
	"os"

	"github.com/samber/lo"
)

const (
	levelAlert = slog.LevelError + 1
)

func Debug(ctx context.Context, msg string, attrs ...slog.Attr) {
	logger := LoggerFromContext(ctx)
	logger.DebugContext(ctx, msg, lo.ToAnySlice(attrs)...)
}

func Info(ctx context.Context, msg string, attrs ...slog.Attr) {
	logger := LoggerFromContext(ctx)
	logger.InfoContext(ctx, msg, lo.ToAnySlice(attrs)...)
}

func Warn(ctx context.Context, msg string, attrs ...slog.Attr) {
	logger := LoggerFromContext(ctx)
	logger.WarnContext(ctx, msg, lo.ToAnySlice(attrs)...)
}

func Error(ctx context.Context, msg string, attrs ...slog.Attr) {
	logger := LoggerFromContext(ctx)
	logger.ErrorContext(ctx, msg, lo.ToAnySlice(attrs)...)
}

func Alert(ctx context.Context, msg string, attrs ...slog.Attr) {
	logger := LoggerFromContext(ctx)
	logger.Log(ctx, levelAlert, msg, lo.ToAnySlice(attrs)...)
}

type logKey struct{}

func LoggerFromContext(ctx context.Context) *slog.Logger {
	val := ctx.Value(logKey{})
	if logger, ok := val.(*slog.Logger); ok {
		return logger
	}

	return NewLogger()
}

func ContextWithAttrs(ctx context.Context, attrs ...slog.Attr) context.Context {
	logger := LoggerFromContext(ctx)
	return ContextWithLogger(ctx, logger.With(lo.ToAnySlice(attrs)...))
}

func ContextWithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, logKey{}, logger)
}

func NewLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key != slog.LevelKey {
				return a
			}
			if a.Value.Any().(slog.Level) == levelAlert {
				a.Value = slog.StringValue("ALERT")
			}
			return a
		},
	}))
}
