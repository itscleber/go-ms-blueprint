package services

import "github.com/itscleber/go-ms-blueprint/internal/repositories"

type LivenessService struct {
	repo repositories.LivenessRepository
}

func NewLivenessService(repo repositories.LivenessRepository) *LivenessService {
	return &LivenessService{repo: repo}
}

func (s *LivenessService) IsAlive() bool {
	return s.repo.IsAlive()
}
