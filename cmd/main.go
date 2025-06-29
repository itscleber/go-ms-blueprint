package main

import (
	"context"
	"log"
	"os"

	"template/internal/api"
	"template/internal/config"
	"template/internal/telemetry"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

var serviceName string

func init() {
	config.SetupLogger()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	serviceName = os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatalf("SERVICE_NAME environment variable is not set")
	}
}

func main() {
	ctx := context.Background()

	tp, err := telemetry.InitTracer(ctx)
	if err != nil {
		log.Fatalf("failed to initialize OpenTelemetry: %v", err)
	}
	defer func() { _ = tp.Shutdown(ctx) }()

	r := gin.Default()
	r.Use(otelgin.Middleware(serviceName))
	api.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
