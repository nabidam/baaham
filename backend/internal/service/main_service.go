package service

import "github.com/nabidam/baaham/internal/repository"

type MainService struct {
	HealthService IHealthService
}

func NewMainService(repo *repository.MainRepository) *MainService {
	healthSvc := NewHealthService(repo.HealthRepository)

	return &MainService{HealthService: healthSvc}
}
