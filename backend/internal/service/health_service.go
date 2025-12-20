package service

import (
	"context"

	"github.com/nabidam/baaham/internal/repository"
)

type IHealthService interface {
	HealthCheck(ctx context.Context) (bool, error)
}

type HealthService struct {
	repo repository.IHealthRepository
}

func NewHealthService(r repository.IHealthRepository) *HealthService {
	return &HealthService{repo: r}
}

func (s *HealthService) HealthCheck(ctx context.Context) (bool, error) {
	return s.repo.Check(ctx)
}
