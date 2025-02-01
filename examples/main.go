package main

import (
	"context"
	"log"

	"github.com/bylucasqueiroz/loglib"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	logger := loglib.GetLogger()
	ctx := loglib.WithCorrelationID(context.Background(), "12345-ABCDE")

	logger.Info(ctx, "Application started")
	logger.Debug(ctx, "Debugging some logic")
	logger.Error(ctx, "An error occurred")
}
