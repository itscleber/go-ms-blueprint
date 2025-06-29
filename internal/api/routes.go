package api

import (
	"template/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, healthHandler *handlers.HealthHandler) {
	registerHealthRoutes(r, healthHandler)
}
