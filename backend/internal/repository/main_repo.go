package repository

import "github.com/jackc/pgx/v5/pgxpool"

type MainRepository struct {
	HealthRepository *HealthRepository
}

func NewMainRepository(db *pgxpool.Pool) *MainRepository {
	healthRepo := NewHealthRepository(db)
	return &MainRepository{HealthRepository: healthRepo}
}
