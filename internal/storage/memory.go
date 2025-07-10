package storage

import (
	"context"
	"github.com/mohar9h/go-sanctum/internal/auth"
	"sync"
	"time"
)

type MemoryStorage struct {
	mu     sync.RWMutex
	tokens map[string]auth.Token
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tokens: make(map[string]auth.Token),
	}
}

func (s *MemoryStorage) RevokeToken(ctx context.Context, token string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for key, t := range s.tokens {
		if valid, _ := auth.VerifyToken(token, t.Hash); valid {
			delete(s.tokens, key)
			return nil
		}
	}

	return auth.ErrTokenInvalid
}

func (s *MemoryStorage) StoreToken(ctx context.Context, token auth.Token) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tokens[string(token.Hash)] = token
	return nil
}

func (s *MemoryStorage) ValidateToken(ctx context.Context, plaintext string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, token := range s.tokens {
		if valid, _ := auth.VerifyToken(plaintext, token.Hash); valid {
			// اگر توکن تاریخ انقضا دارد و منقضی شده
			if token.ExpiresAt != nil && time.Now().After(*token.ExpiresAt) {
				return "", auth.ErrTokenExpired
			}
			return token.UserID, nil
		}
	}

	return "", auth.ErrTokenInvalid
}
