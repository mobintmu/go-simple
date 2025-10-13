package auth

import (
	"time"

	"go-simple/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(cfg *config.Config, adminID string) (string, error) {
	claims := jwt.MapClaims{
		"admin_id": adminID,
		"exp":      time.Now().Add(time.Hour * time.Duration(cfg.JWTExpiryHours)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}
