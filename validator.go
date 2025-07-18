package auth

import (
	"errors"
	"github.com/mohar9h/go-sanctum/config"
	"strings"
	"time"

	"github.com/mohar9h/go-sanctum/storage"
	"github.com/mohar9h/go-sanctum/utils"
)

var ErrTokenInvalid = errors.New("token is invalid or expired")

func ValidateToken(raw string, cfg *config.Config) (*storage.Token, error) {

	if cfg == nil {
		cfg = config.DefaultConfig()
	}
	cfg.ApplyDefaults()

	parts := strings.Split(strings.ReplaceAll(raw, " ", ""), "|")
	if len(parts) != 2 {
		return nil, ErrTokenInvalid
	}

	hashed := utils.HashToken(parts[1])

	tok, err := cfg.Storage.FindByHash(hashed)
	if err != nil {
		return nil, err
	}

	if tok.Token != hashed {
		return nil, ErrTokenInvalid
	}

	if tok.ExpiresAt != nil && time.Now().After(*tok.ExpiresAt) {
		return nil, errors.New("token expired")
	}

	go func() {
		err = cfg.Storage.TouchLastUsed(tok.Token)
		if err != nil {

		}
	}()

	return tok, nil
}
