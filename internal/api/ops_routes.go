package api

import (
	"template/internal/handlers"

	"github.com/gin-gonic/gin"
)

func registerOpsRoutes(r *gin.Engine, h *handlers.ReadinessHandler) {
	r.GET("/ops/ready", h.Handle)
}
