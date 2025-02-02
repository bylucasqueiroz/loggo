package loggo

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// getLogLevel retrieves the log level from environment variables.
func getLogLevel() zapcore.Level {
	level := strings.ToLower(getEnv("LOG_LEVEL", "info"))
	switch level {
	case "debug":
		return zap.DebugLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

// getLogEncoder selects JSON or console encoding based on environment.
func getLogEncoder() zapcore.Encoder {
	format := strings.ToLower(os.Getenv("LOG_FORMAT"))
	if format == "console" {
		return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	}
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

// getLogOutput sets the log output destination.
func getLogOutput() zapcore.WriteSyncer {
	output := getEnv("LOG_OUTPUT", "stdout")
	if output == "stdout" {
		return zapcore.Lock(os.Stdout)
	}
	file, err := os.Create(output)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}
	return zapcore.AddSync(file)
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
