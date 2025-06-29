package services

import (
	"github.com/itscleber/go-ms-blueprint/internal/repositories"
	"time"
)

type HealthStatus struct {
	Status string `json:"status"`
	Uptime string `json:"uptime,omitempty"`
}

type HealthService struct {
	repo      repositories.HealthRepository
	startTime time.Time
}

func NewHealthService(repo repositories.HealthRepository) *HealthService {
	return &HealthService{
		repo:      repo,
		startTime: time.Now(),
	}
}

func (s *HealthService) Check() HealthStatus {
	uptime := time.Since(s.startTime).Truncate(time.Second).String()
	return HealthStatus{
		Status: s.repo.Status(),
		Uptime: uptime,
	}
}
