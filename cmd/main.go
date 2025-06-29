package main

import (
	"context"
	"log"

	"github.com/itscleber/go-ms-blueprint/internal/api"
	"github.com/itscleber/go-ms-blueprint/internal/config"
	"github.com/itscleber/go-ms-blueprint/internal/telemetry"
)

func main() {
	ctx := context.Background()
	serviceName := config.LoadEnvAndLogger()

	tp := telemetry.MustInitTracer(ctx)
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Println("tracer shutdown error:", err)
		}
	}()


	r := api.SetupRouter(serviceName)

	port := config.GetPort()
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
