package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReadinessService interface {
	IsReady() bool
}

type ReadinessHandler struct {
	svc ReadinessService
}

func NewReadinessHandler(svc ReadinessService) *ReadinessHandler {
	return &ReadinessHandler{svc: svc}
}

func (h *ReadinessHandler) Handle(c *gin.Context) {
	status := "not ready"
	code := http.StatusServiceUnavailable

	if h.svc.IsReady() {
		status = "ready"
		code = http.StatusOK
	}

	c.JSON(code, gin.H{"status": status})
}
