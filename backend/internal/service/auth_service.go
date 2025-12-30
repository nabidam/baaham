package service

import (
	"context"
	"errors"
	"time"

	"github.com/nabidam/baaham/internal/domain"
	"github.com/nabidam/baaham/pkg/jwt"
	pass "github.com/nabidam/baaham/pkg/password"
)

type AuthService struct {
	repo      domain.UserRepository
	jwtSecret string
}

func NewAuthService(r domain.UserRepository, jwtSecret string) domain.AuthService {
	return &AuthService{repo: r, jwtSecret: jwtSecret}
}

func (s *AuthService) Login(ctx context.Context, username string, password string) (*domain.LoginResponse, error) {
	// Get user by username
	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	// compare passwords
	isPasswordCorrect := pass.CheckPasswordHash(password, user.PasswordHash)
	if !isPasswordCorrect {
		return nil, errors.New("invalid credentials")
	}

	// generate JWT token
	token, err := jwt.GenerateToken(user.ID, user.Username, user.IsAdmin, []byte(s.jwtSecret), 24*time.Hour)

	return &domain.LoginResponse{
		Token: token,
	}, nil
}
