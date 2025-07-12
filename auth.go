package auth

import (
	"errors"
	"fmt"
	"github.com/mohar9h/go-sanctum/token"

	"github.com/mohar9h/go-sanctum/config"
	"github.com/mohar9h/go-sanctum/storage"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	cfg             *config.Config
)

// Init initializes the auth system with the given configuration.
// This should be called once at startup.
func Init(c *config.Config) error {
	if err := c.Validate(); err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}
	cfg = c
	token.Setup(cfg)
	if cfg.Storage != nil {
		storage.Setup(cfg.Storage)
	}
	return nil
}

// CreateToken creates a new auth for the given user, name, and abilities.
// It supports both JWT and random-auth modes.
func CreateToken(userID, name string, abilities []string) (string, error) {
	if cfg == nil {
		return "", errors.New("auth not initialized")
	}

	return token.Create(userID, name, abilities)
}
