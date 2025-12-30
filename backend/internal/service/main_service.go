package service

import (
	"github.com/nabidam/baaham/internal/config"
	"github.com/nabidam/baaham/internal/domain"
	"github.com/nabidam/baaham/internal/repository"
)

type MainService struct {
	HealthService domain.HealthService
	AuthService   domain.AuthService
}

func NewMainService(repo *repository.MainRepository, cfg *config.Config) *MainService {
	healthSvc := NewHealthService(repo.HealthRepository)
	authSvc := NewAuthService(repo.UserRepository, cfg.JWTSecret)

	return &MainService{HealthService: healthSvc, AuthService: authSvc}
}
