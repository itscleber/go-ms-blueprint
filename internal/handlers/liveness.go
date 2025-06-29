package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LivenessService interface {
	IsAlive() bool
}

type LivenessHandler struct {
	svc LivenessService
}

func NewLivenessHandler(svc LivenessService) *LivenessHandler {
	return &LivenessHandler{svc: svc}
}

func (h *LivenessHandler) Handle(c *gin.Context) {
	status := "not alive"
	code := http.StatusServiceUnavailable

	if h.svc.IsAlive() {
		status = "alive"
		code = http.StatusOK
	}

	c.JSON(code, gin.H{"status": status})
}
