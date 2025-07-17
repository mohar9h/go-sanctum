package storage

import (
	"errors"
	"sync"
	"time"
)

type memoryDriver struct {
	tokens map[string]*Token // key is hashed token string
	mu     sync.RWMutex
}

var _ Driver = (*memoryDriver)(nil)

func NewMemoryDriver() Driver {
	return &memoryDriver{
		tokens: make(map[string]*Token),
	}
}

// StoreToken stores the token using its hashed value as key
func (m *memoryDriver) StoreToken(t *Token) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tokens[t.Token] = t
	return nil
}

// FindByID looks up token by its internal ID (numeric)
func (m *memoryDriver) FindByID(id int64) (*Token, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, token := range m.tokens {
		if token.ID == id {
			if token.ExpiresAt != nil && time.Now().After(*token.ExpiresAt) {
				return nil, errors.New("token expired")
			}
			return token, nil
		}
	}
	return nil, errors.New("token not found")
}

// FindByHash looks up token by its hashed token string
func (m *memoryDriver) FindByHash(hash string) (*Token, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	tok, ok := m.tokens[hash]
	if !ok {
		return nil, errors.New("token not found")
	}
	if tok.ExpiresAt != nil && time.Now().After(*tok.ExpiresAt) {
		return nil, errors.New("token expired")
	}
	return tok, nil
}

// RevokeToken removes a token by its hashed value
func (m *memoryDriver) RevokeToken(hash string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.tokens, hash)
	return nil
}

// TouchLastUsed updates the last used time for analytics or session freshness
func (m *memoryDriver) TouchLastUsed(hash string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	tok, ok := m.tokens[hash]
	if !ok {
		return errors.New("token not found")
	}
	now := time.Now()
	tok.LastUsedAt = &now
	return nil
}
