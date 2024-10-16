package helper

import (
	"errors"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUseCase interface {
	GenerateAccessToken(claims JwtCustomClaims) (string, error)
	BlacklistAccessToken(token string) error
	IsTokenBlacklisted(token string) bool
}

type TokenUseCaseImpl struct {
	blacklistedTokens map[string]struct{}
	mu                sync.RWMutex
}

type JwtCustomClaims struct {
	ID      string `json:"id"`
	Name    string `json:"username"`
	Fullname string `json:"fullname"`
	Email   string `json:"email"`
	Phone   string `json:"no_phone"`
	Picture string `json:"picture"`
	Role string `json:"role"`
	Google string `json:"google_id"`
	NoKk string   `json:"no_kk"`
	jwt.RegisteredClaims
}

func NewTokenUseCase() *TokenUseCaseImpl {
	return &TokenUseCaseImpl{
		blacklistedTokens: make(map[string]struct{}),
	}
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

func (t *TokenUseCaseImpl) BlacklistAccessToken(token string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if _, exists := t.blacklistedTokens[token]; exists {
		return errors.New("token already blacklisted")
	}

	t.blacklistedTokens[token] = struct{}{}
	return nil
}

func (t *TokenUseCaseImpl) IsTokenBlacklisted(token string) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()

	_, exists := t.blacklistedTokens[token]
	return exists
}
