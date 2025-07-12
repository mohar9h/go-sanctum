package token

import (
	"crypto/sha256"
	"encoding/hex"
)

// hashToken returns SHA256 hash of token (for storage).
func hashToken(raw string) string {
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}
