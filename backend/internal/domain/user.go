package domain

import (
	"context"
	"time"
)

type User struct {
	ID           string    `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	IsAdmin      bool      `db:"is_admin"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type UserRepository interface {
	Create(ctx context.Context, username string, passwordHash string, isAdmin bool) (*User, error)
	List(ctx context.Context) ([]User, error)
	UpdatePassword(ctx context.Context, username string, passwordHash string) error
	Delete(ctx context.Context, username string) error
	GetByUsername(ctx context.Context, username string) (*User, error)
}

type UserService interface {
	Create(ctx context.Context, username string, password string, isAdmin bool) (User, error)
}
