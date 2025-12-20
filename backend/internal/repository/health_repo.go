package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IHealthRepository interface {
	Check(ctx context.Context) (bool, error)
}

type HealthRepository struct {
	db *pgxpool.Pool
}

func NewHealthRepository(db *pgxpool.Pool) *HealthRepository {
	return &HealthRepository{db: db}
}

func (r *HealthRepository) Check(ctx context.Context) (bool, error) {
	return true, nil
}
