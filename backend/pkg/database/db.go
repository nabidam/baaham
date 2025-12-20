package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/nabidam/baaham/internal/config"
)

func NewPool(
	cfg *config.Config,
) (*pgxpool.Pool, error) {
	ctx := context.Background()

	dbCfg := cfg.Database
	logger := cfg.Logger
	config, err := pgxpool.ParseConfig(dbCfg.DSN)
	if err != nil {
		logger.Error("failed to parse db dsn", zap.Error(err))
		return nil, err
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.HealthCheckPeriod = time.Duration(30) * time.Second

	logger.Info("connecting to database")

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		logger.Error("failed to create db pool", zap.Error(err))
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		logger.Error("db ping failed", zap.Error(err))
		pool.Close()
		return nil, err
	}

	logger.Info("database connected")

	return pool, nil
}
