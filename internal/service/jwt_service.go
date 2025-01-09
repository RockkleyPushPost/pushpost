package service

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

type JWTService struct {
	secretKey   string
	tokenExpiry time.Duration
}

func NewJWTService(tokenExpiry time.Duration, secretKey string) *JWTService {
	return &JWTService{tokenExpiry: tokenExpiry, secretKey: secretKey}
}

func (j *JWTService) GenerateToken(userUUID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_uuid": userUUID,
		"exp":       time.Now().Add(j.tokenExpiry).Unix(),
		"iat":       time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}
