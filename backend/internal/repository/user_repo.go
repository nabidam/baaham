package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nabidam/baaham/internal/domain"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Create(ctx context.Context, username string, passwordHash string, isAdmin bool) (*domain.User, error) {
	var u domain.User
	err := repo.db.QueryRow(ctx, `
		INSERT INTO users (username, password_hash, is_admin)
		VALUES ($1, $2, $3)
		RETURNING id, username, password_hash, is_admin, created_at, updated_at
	`, username, passwordHash, isAdmin).Scan(
		&u.ID,
		&u.Username,
		&u.PasswordHash,
		&u.IsAdmin,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	return &u, err
}

func (repo *UserRepository) UpdatePassword(ctx context.Context, username string, passwordHash string) error {
	cmd, err := repo.db.Exec(ctx, `
		UPDATE users
		SET password_hash = $1, updated_at = now()
		WHERE username = $2
	`, passwordHash, username)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (repo *UserRepository) List(ctx context.Context) ([]domain.User, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT id, username, password_hash, is_admin, created_at, updated_at
		FROM users
		ORDER BY created_at ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.PasswordHash,
			&u.IsAdmin,
			&u.CreatedAt,
			&u.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (repo *UserRepository) Delete(ctx context.Context, username string) error {
	_, err := repo.db.Exec(ctx, `
		DELETE FROM users WHERE username = $1
	`, username)
	return err
}
