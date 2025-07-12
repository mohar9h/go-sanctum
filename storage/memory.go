package storage

import (
	"errors"
	"sync"
	"time"
)

type memoryDriver struct {
	tokens map[string]*Token
	mu     sync.RWMutex
}

var _ Driver = (*memoryDriver)(nil)

func NewMemoryDriver() Driver {
	return &memoryDriver{
		tokens: make(map[string]*Token),
	}
}

func (m *memoryDriver) StoreToken(t *Token) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tokens[t.ID] = t
	return nil
}

func (m *memoryDriver) FindToken(id string) (*Token, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	tok, ok := m.tokens[id]
	if !ok {
		return nil, errors.New("token not found")
	}
	if tok.ExpiresAt != nil && time.Now().After(*tok.ExpiresAt) {
		return nil, errors.New("token expired")
	}
	return tok, nil
}

func (m *memoryDriver) RevokeToken(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.tokens, id)
	return nil
}
