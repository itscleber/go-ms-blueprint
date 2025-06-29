package api

import (
	"github.com/itscleber/go-ms-blueprint/internal/handlers"

	"github.com/gin-gonic/gin"
)

func registerHealthRoutes(r *gin.Engine, h *handlers.HealthHandler) {
	r.GET("/v1/health", h.HealthCheck)
}
