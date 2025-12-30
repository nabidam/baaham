package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nabidam/baaham/internal/domain"
)

func GenerateToken(userID string, username string, isAdmin bool, secret []byte, expiration time.Duration) (string, error) {
	claims := domain.UserClaims{
		UserID:   userID,
		Username: username,
		IsAdmin:  isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
