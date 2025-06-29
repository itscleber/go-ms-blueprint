package main

import (
	"context"
	"log"

	"github.com/itscleber/go-ms-blueprint/internal/api"
	"github.com/itscleber/go-ms-blueprint/internal/config"
	"github.com/itscleber/go-ms-blueprint/internal/telemetry"

	"go.opentelemetry.io/otel"
)

func main() {
	ctx := context.Background()

	// inicia tracer
	tp := telemetry.MustInitTracer(ctx)
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Println("tracer shutdown error:", err)
		}
	}()

	// inicia trace principal
	ctx, span := otel.Tracer("main").Start(ctx, "startup")
	defer span.End()

	serviceName := config.LoadEnvAndLogger()

	// inicia o router com trace
	r := api.SetupRouter(serviceName)

	port := config.GetPort()
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
