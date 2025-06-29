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
	if h.svc.IsReady() {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusServiceUnavailable)
	}
}
