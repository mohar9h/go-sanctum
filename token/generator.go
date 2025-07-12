package token

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mohar9h/go-sanctum/storage"
	"github.com/mohar9h/go-sanctum/utils"
	"strings"
	"time"
)

func Create(userID, name string, abilities []string) (string, error) {
	if cfg.UseJWT {
		return createJWT(userID, name, abilities)
	}

	tok, err := CreateRandomToken(userID, name, abilities)

	if err != nil {
		return "", err
	}
	return userID + " | " + tok.PlainText, nil
}

func createJWT(userID, name string, abilities []string) (string, error) {
	header := map[string]string{
		"alg":  cfg.SigningMethod,
		"type": "JWT",
	}

	now := time.Now().Unix()
	expireAt := int64(0)
	if cfg.ExpireAt > 0 {
		expireAt = now + int64(cfg.ExpireAt)
	}

	payload := Claims{
		UserId:    userID,
		Name:      name,
		Abilities: abilities,
		IssueAt:   now,
		ExpireAt:  expireAt,
	}

	// Encode header and payload
	headerJson, _ := json.Marshal(header)
	payloadJson, _ := json.Marshal(payload)

	headerB64 := base64.RawURLEncoding.EncodeToString(headerJson)
	payloadB64 := base64.RawURLEncoding.EncodeToString(payloadJson)

	message := headerB64 + "." + payloadB64

	var signature string
	var err error

	switch cfg.SigningMethod {
	case "HS256":
		signature, err = utils.SignHMAC(message, cfg.SigningKey)
	case "RS256":
		signature, err = utils.SignRSA(message, cfg.PrivateKey)
	default:
		return "", errors.New("unsupported signing method")
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s.%s", message, signature), nil

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
