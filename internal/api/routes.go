package api

import (
	"github.com/itscleber/go-ms-blueprint/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, healthHandler *handlers.HealthHandler, readinessHandler *handlers.ReadinessHandler, livenessHandler *handlers.LivenessHandler) {
	registerHealthRoutes(r, healthHandler)
	registerOpsRoutes(r, readinessHandler, livenessHandler)
}
