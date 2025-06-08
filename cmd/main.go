package main

import (
	"context"
	"log"
	"os"

	"template/api"
	"template/config"
	"template/telemetry"

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
	tp, err := telemetry.InitTracer()
	if err != nil {
		log.Fatalf("failed to initialize OpenTelemetry: %v", err)
	}
	defer func() { _ = tp.Shutdown(context.Background()) }()

	r := gin.Default()
	r.Use(otelgin.Middleware(serviceName))
	api.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
