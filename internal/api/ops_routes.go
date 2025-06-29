package api

import (
	"github.com/itscleber/go-ms-blueprint/internal/handlers"

	"github.com/gin-gonic/gin"
)

func registerOpsRoutes(r *gin.Engine, readiness *handlers.ReadinessHandler, liveness *handlers.LivenessHandler) {
	r.GET("/ops/ready", readiness.Handle)
	r.GET("/ops/live", liveness.Handle)
}
