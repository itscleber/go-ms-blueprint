package handlers

import (
	"net/http"
	"github.com/itscleber/go-ms-blueprint/internal/services"

	"github.com/gin-gonic/gin"
)

type HealthService interface {
	Check() services.HealthStatus
}

type HealthHandler struct {
	svc HealthService
}

func NewHealthHandler(svc HealthService) *HealthHandler {
	return &HealthHandler{svc: svc}
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	status := h.svc.Check()
	c.JSON(http.StatusOK, status)
}
