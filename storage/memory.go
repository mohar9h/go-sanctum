package storage

import (
	"errors"
	"sync"
	"time"
)

type memoryDriver struct {
	tokens map[int64]*Token
	mu     sync.RWMutex
}

var _ Driver = (*memoryDriver)(nil)

func NewMemoryDriver() Driver {
	return &memoryDriver{
		tokens: make(map[int64]*Token),
	}
}

func (m *memoryDriver) StoreToken(t *Token) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tokens[t.ID] = t
	return nil
}

func (m *memoryDriver) FindToken(id int64) (*Token, error) {
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

func (m *memoryDriver) RevokeToken(id int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.tokens, id)
	return nil
}
