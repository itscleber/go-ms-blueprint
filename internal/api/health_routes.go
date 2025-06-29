package api

import (
	"template/internal/handlers"
	"template/internal/repositories"
	"template/internal/services"

	"github.com/gin-gonic/gin"
)

func registerHealthRoutes(r *gin.Engine) {
	repo := repositories.StaticHealthRepository{}
	svc := services.NewHealthService(repo)
	h := handlers.NewHealthHandler(svc)

	r.GET("/health", h.HealthCheck)
}
