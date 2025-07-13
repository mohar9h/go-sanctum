package token

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/mohar9h/go-sanctum/storage"
	"strings"
	"time"
)

func Create(userID, name string, abilities []string) (string, error) {

	tok, err := CreateRandomToken(userID, name, abilities)

	if err != nil {
		return "", err
	}
	return userID + " | " + tok.PlainText, nil
}

func CreateRandomToken(userID, name string, abilities []string) (*Result, error) {
	if cfg.Storage == nil {
		return nil, errors.New("no storage backend configured")
	}

	// Generate random token
	buf := make([]byte, cfg.TokenLength)
	if _, err := rand.Read(buf); err != nil {
		return nil, err
	}

	raw := base64.RawURLEncoding.EncodeToString(buf)
	plainText := cfg.TokenPrefix + raw
	hashed := hashToken(plainText)

	// Handle expiration logic
	var expireAt *time.Time
	if cfg.ExpireAt > 0 {
		t := time.Now().Add(cfg.ExpireAt)
		expireAt = &t
	}

	token := &storage.Token{
		ID:        hashed,
		UserID:    userID,
		Name:      name,
		Abilities: strings.Join(abilities, ","),
		CreatedAt: time.Now(),
		ExpiresAt: expireAt,
	}

	if err := cfg.Storage.StoreToken(token); err != nil {
		return nil, err
	}

	return &Result{
		PlainText: plainText,
		TokenID:   hashed,
	}, nil
}
