package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"time"
)

var (
	ErrTokenExpired    = errors.New("token expired")
	ErrTokenInvalid    = errors.New("invalid token")
	ErrTokenGeneration = errors.New("token generation failed")
)

// GenerateToken ایجاد توکن جدید
func GenerateToken(userID string, expiresIn *time.Duration) (*Token, error) {
	token := &Token{
		UserID:     userID,
		CreatedAt:  time.Now(),
		LastUsedAt: time.Now(),
	}

	if expiresIn != nil {
		expiresAt := time.Now().Add(*expiresIn)
		token.ExpiresAt = &expiresAt
	}

	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		return nil, ErrTokenGeneration
	}

	token.Plaintext = hex.EncodeToString(randomBytes)
	hashed, err := hashToken(token.Plaintext)
	if err != nil {
		return nil, err
	}
	token.Hash = hashed

	return token, nil
}

// VerifyToken بررسی تطابق توکن
func VerifyToken(plaintext string, hashed []byte) (bool, error) {
	hashedInput, err := hashToken(plaintext)
	if err != nil {
		return false, err
	}

	return subtle.ConstantTimeCompare(hashedInput, hashed) == 1, nil
}

func hashToken(token string) ([]byte, error) {
	// در عمل از الگوریتمی مثل bcrypt استفاده کنید
	return []byte(token), nil
}
