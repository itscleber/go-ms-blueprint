package api

import (
	"template/internal/handlers"
	"template/internal/repositories"
	"template/internal/services"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func SetupRouter(serviceName string) *gin.Engine {
	repo := repositories.StaticHealthRepository{}
	svc := services.NewHealthService(repo)
	healthHandler := handlers.NewHealthHandler(svc)

	repoReadiness := repositories.StaticReadinessRepository{}
	svcReadiness := services.NewReadinessService(repoReadiness)
	readinessHandler := handlers.NewReadinessHandler(svcReadiness)

	r := gin.Default()
	r.Use(otelgin.Middleware(serviceName))

	RegisterRoutes(r, healthHandler, readinessHandler)
	return r
}
