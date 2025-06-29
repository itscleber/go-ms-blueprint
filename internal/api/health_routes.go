package api

import (
	"template/internal/handlers"

	"github.com/gin-gonic/gin"
)

func registerHealthRoutes(r *gin.Engine, h *handlers.HealthHandler) {
	r.GET("/v1/health", h.HealthCheck)
}
