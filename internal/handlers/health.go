package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"template/internal/services"
)

type HealthHandler struct {
	svc services.HealthServiceInterface
}

func NewHealthHandler(svc services.HealthServiceInterface) *HealthHandler {
	return &HealthHandler{svc: svc}
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	status := h.svc.Check()
	c.JSON(http.StatusOK, gin.H{"status": status})
}
