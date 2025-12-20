package service

import "context"

type HealthRepository interface {
	Check(ctx context.Context) (bool, error)
}

type HealthService struct {
	repo HealthRepository
}

func NewHealthService(r HealthRepository) *HealthService {
	return &HealthService{repo: r}
}

func (s *HealthService) HealthCheck(ctx context.Context) (bool, error) {
	return s.repo.Check(ctx)
}
