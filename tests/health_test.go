package tests

import (
	"net/http"
	"net/http/httptest"
	"template/internal/handlers"
	"template/internal/repositories"
	"template/internal/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler_HealthCheck(t *testing.T) {
	t.Setenv("PORT", "8080")
	gin.SetMode(gin.TestMode)

	// setup router com dependências reais
	repo := repositories.StaticHealthRepository{}
	svc := services.NewHealthService(repo)
	handler := handlers.NewHealthHandler(svc)

	router := gin.New()
	router.GET("/health", handler.HealthCheck)

	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, `{"status":"healthy"}`, recorder.Body.String())
}
