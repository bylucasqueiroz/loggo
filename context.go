package loggo

import (
	"context"
	"os"
)

// correlationKey is a custom type to avoid collisions in context keys
type correlationKey string

// getCorrelationKey retrieves the correlation ID key from environment variables
func getCorrelationKey() string {
	key := os.Getenv("LOG_CORRELATION_ID_KEY")
	if key == "" {
		key = "correlation_id"
	}
	return key
}

// GetCorrelationID retrieves the Correlation ID from context or returns "unknown".
func GetCorrelationID(ctx context.Context) string {
	if ctx == nil {
		return "unknown"
	}
	if id, ok := ctx.Value(correlationKey(getCorrelationKey())).(string); ok {
		return id
	}
	return "unknown"
}

// WithCorrelationID injects a Correlation ID into the context.
func WithCorrelationID(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, correlationKey(getCorrelationKey()), correlationID)
}
