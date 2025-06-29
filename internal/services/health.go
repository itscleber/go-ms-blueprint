package services

import "template/internal/repositories"

type HealthServiceInterface interface {
	Check() string
}

type HealthService struct {
	repo repositories.HealthRepository
}

func NewHealthService(repo repositories.HealthRepository) *HealthService {
	return &HealthService{repo: repo}
}

func (s *HealthService) Check() string {
	return s.repo.Status()
}

