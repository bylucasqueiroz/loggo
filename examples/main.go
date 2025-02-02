package main

import (
	"context"
	"log"

	"github.com/bylucasqueiroz/loggo"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	logger := loggo.GetLogger()
	ctx := loggo.WithCorrelationID(context.Background(), "12345-ABCDE")

	logger.Info(ctx, "Application started")
	logger.Debug(ctx, "Debugging some logic")
	logger.Error(ctx, "An error occurred")
}
