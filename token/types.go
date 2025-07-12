package token

import "github.com/mohar9h/go-sanctum/config"

var cfg *config.Config

func Setup(c *config.Config) {
	cfg = c
}

// Result TokenResult is the result of a successful token creation.
type Result struct {
	PlainText string // what the client receives
	TokenID   string // internal hashed ID for storage
}
