package loggo

import (
	"context"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zapLogger *zap.Logger
}

var (
	instance *Logger
	once     sync.Once
)

// NewLogger initializes the logger with environment-based config.
func NewLogger() *Logger {
	once.Do(func() {
		core := zapcore.NewCore(getLogEncoder(), getLogOutput(), getLogLevel())

		logger := zap.New(core)
		instance = &Logger{zapLogger: logger}
	})
	return instance
}

// GetLogger returns the singleton instance.
func GetLogger() *Logger {
	if instance == nil {
		return NewLogger()
	}
	return instance
}

// Logging functions with Correlation ID
func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("correlation_id", GetCorrelationID(ctx)))
	l.zapLogger.Info(msg, fields...)
}

func (l *Logger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("correlation_id", GetCorrelationID(ctx)))
	l.zapLogger.Debug(msg, fields...)
}

func (l *Logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("correlation_id", GetCorrelationID(ctx)))
	l.zapLogger.Error(msg, fields...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("correlation_id", GetCorrelationID(ctx)))
	l.zapLogger.Fatal(msg, fields...)
}
