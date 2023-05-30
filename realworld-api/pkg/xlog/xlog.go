package xlog

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	logger := LoggerFromContext(ctx)
	logger.Info(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	logger := LoggerFromContext(ctx)
	logger.Warn(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	logger := LoggerFromContext(ctx)
	logger.Error(msg, fields...)
}

func Alert(ctx context.Context, msg string, fields ...zap.Field) {
	logger := LoggerFromContext(ctx)
	logger.Error(msg, fields...)
}

type logKey struct{}

func LoggerFromContext(ctx context.Context) *zap.Logger {
	val := ctx.Value(logKey{})
	if logger, ok := val.(*zap.Logger); ok {
		return logger
	}
	logger, err := zap.NewDevelopment(zap.WithCaller(true))
	if err != nil {
		panic(fmt.Sprintf("failed to initialize logger: %v", err))
	}
	return logger
}

func ContextWithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, logKey{}, logger)
}

func ContextWithLogFields(ctx context.Context, fields ...zap.Field) context.Context {
	return ContextWithLogger(ctx, LoggerFromContext(ctx).With(fields...))
}

// see: https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogSeverity
var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel: "DEBUG",
	zapcore.InfoLevel:  "INFO",
	zapcore.WarnLevel:  "WARNING",
	zapcore.ErrorLevel: "ERROR",
}

func encodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(logLevelSeverity[l])
}

func NewLogger() *zap.Logger {
	// see: https://cloud.google.com/logging/docs/structured-logging?hl=ja#special-payload-fields
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.MessageKey = "message"
	encoderConfig.LevelKey = "severity"
	encoderConfig.EncodeLevel = encodeLevel
	encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(zap.DebugLevel)
	cfg.EncoderConfig = encoderConfig
	cfg.DisableStacktrace = true

	logger, err := cfg.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize logger: %v", err))
	}
	return logger
}
