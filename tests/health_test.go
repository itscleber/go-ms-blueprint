package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"template/internal/handlers"
	"template/internal/repositories"
	"template/internal/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type readinessResponse struct {
	Status string `json:"status"`
}

func TestHealthHandler_HealthCheck(t *testing.T) {
	t.Setenv("PORT", "8080")
	gin.SetMode(gin.TestMode)

	repo := repositories.StaticHealthRepository{}
	svc := services.NewHealthService(repo)
	handler := handlers.NewHealthHandler(svc)

	router := gin.New()
	router.GET("/health", handler.HealthCheck)

	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var body services.HealthStatus
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.NoError(t, err)

	assert.Equal(t, "healthy", body.Status)
	assert.NotEmpty(t, body.Uptime)
}

func TestReadinessHandler_ReadinessCheck(t *testing.T) {
	t.Setenv("PORT", "8080")
	gin.SetMode(gin.TestMode)

	repo := repositories.StaticReadinessRepository{}
	svc := services.NewReadinessService(repo)
	handler := handlers.NewReadinessHandler(svc)

	router := gin.New()
	router.GET("/ops/ready", handler.Handle)

	req, _ := http.NewRequest(http.MethodGet, "/ops/ready", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var body readinessResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.NoError(t, err)
	assert.Equal(t, "ready", body.Status)
}
