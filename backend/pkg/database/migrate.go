package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/nabidam/baaham/migrations"
	"github.com/pressly/goose/v3"
)

func RunMigrations(dbPool *pgxpool.Pool) error {
	goose.SetBaseFS(migrations.EmbedFS)
	goose.SetDialect("postgres")

	db := stdlib.OpenDBFromPool(dbPool)

	return goose.Up(db, ".")
}
