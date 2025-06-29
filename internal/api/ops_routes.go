package api

import (
	"template/internal/handlers"

	"github.com/gin-gonic/gin"
)

func registerOpsRoutes(r *gin.Engine, readiness *handlers.ReadinessHandler, liveness *handlers.LivenessHandler) {
	r.GET("/ops/ready", readiness.Handle)
	r.GET("/ops/live", liveness.Handle)
}
