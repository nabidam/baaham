package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nabidam/baaham/internal/domain"
)

type MainRepository struct {
	HealthRepository domain.HealthRepository
}

func NewMainRepository(db *pgxpool.Pool) *MainRepository {
	healthRepo := NewHealthRepository(db)
	return &MainRepository{HealthRepository: healthRepo}
}
