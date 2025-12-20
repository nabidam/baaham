package service

import (
	"context"

	"github.com/nabidam/baaham/internal/repository"
)

type HealthService struct {
	repo *repository.HealthRepository
}

func NewHealthService(r *repository.HealthRepository) *HealthService {
	return &HealthService{repo: r}
}

func (s *HealthService) HealthCheck(ctx context.Context) (bool, error) {
	return s.repo.Check(ctx)
}
