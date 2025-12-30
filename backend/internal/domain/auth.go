package domain

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

// type AuthRepository interface {
// 	CheckCredentials(ctx context.Context, username string, passwordHash string) (*User, error)
// }

type AuthService interface {
	Login(ctx context.Context, username string, password string) (*LoginResponse, error)
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
