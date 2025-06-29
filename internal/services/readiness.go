package services

import "github.com/itscleber/go-ms-blueprint/internal/repositories"

type ReadinessService struct {
	repo repositories.ReadinessRepository
}

func NewReadinessService(repo repositories.ReadinessRepository) *ReadinessService {
	return &ReadinessService{repo: repo}
}

func (s *ReadinessService) IsReady() bool {
	return s.repo.IsReady()
}
