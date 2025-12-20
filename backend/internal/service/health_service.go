package service

import (
	"context"

	"github.com/nabidam/baaham/internal/domain"
)

type HealthService struct {
	repo domain.HealthRepository
}

func NewHealthService(r domain.HealthRepository) domain.HealthService {
	return &HealthService{repo: r}
}

func (s *HealthService) HealthCheck(ctx context.Context) (bool, error) {
	return s.repo.Check(ctx)
}
