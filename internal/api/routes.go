package api

import (
	"template/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, healthHandler *handlers.HealthHandler, readinessHandler *handlers.ReadinessHandler, livenessHandler *handlers.LivenessHandler) {
	registerHealthRoutes(r, healthHandler)
	registerOpsRoutes(r, readinessHandler, livenessHandler)
}
