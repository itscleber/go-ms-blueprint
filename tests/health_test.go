package tests

import (
	"net/http"
	"net/http/httptest"
	"template/api"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	api.RegisterRoutes(r)

	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"status":"healthy"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
