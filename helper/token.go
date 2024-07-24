package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUseCase interface {
	GenerateAccessToken(claims JwtCustomClaims) (string, error)
}

type TokenUseCaseImpl struct{}

type JwtCustomClaims struct {
	ID    string `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
	Phone string `json:"no_phone"`
	jwt.RegisteredClaims
}

func NewTokenUseCase() *TokenUseCaseImpl {
	return &TokenUseCaseImpl{}
}

func (t *TokenUseCaseImpl) GenerateAccessToken(claims JwtCustomClaims) (string, error) {

	expirationTime := time.Now().Add(1 * time.Hour)
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encodedToken, err := plainToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return encodedToken, nil
}