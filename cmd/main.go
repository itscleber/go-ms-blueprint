package main

import (
	"context"
	"log"

	"template/internal/api"
	"template/internal/config"
	"template/internal/telemetry"
)

func main() {
	ctx := context.Background()
	serviceName := config.LoadEnvAndLogger()

	tp := telemetry.MustInitTracer(ctx)
	defer tp.Shutdown(ctx)

	r := api.SetupRouter(serviceName)

	port := config.GetPort()
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
