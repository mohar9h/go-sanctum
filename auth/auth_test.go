package auth_test

import (
	auth "github.com/mohar9h/go-sanctum"
	"testing"
	"time"

	"github.com/mohar9h/go-sanctum/config"
	"github.com/mohar9h/go-sanctum/storage"
)

func TestCreateRandomToken(t *testing.T) {
	cfg := &config.Config{
		TokenLength:      32,
		TokenPrefix:      "kir_",
		ExpireAt:         2 * time.Hour,
		SigningMethod:    "HS256",
		SigningKey:       "test-key",
		AbilityDelimiter: ":",
		Storage:          storage.NewMemoryDriver(),
	}

	err := auth.Init(cfg)
	if err != nil {
		t.Fatalf("Init failed: %v", err)
	}

	result, err := auth.CreateToken("user123", "test-device", []string{"read:posts", "write:comments"})
	if err != nil {
		t.Fatalf("CreateToken failed: %v", err)
	}

	t.Logf("Generated Token: %+v", result)
}
