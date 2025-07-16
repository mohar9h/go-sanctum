package auth

import "github.com/mohar9h/go-sanctum/config"

// Result TokenResult is the result of a successful token creation.
type Result struct {
	PlainText string // what the client receives
	TokenID   string // internal hashed ID for storage
}

type TokenOptions struct {
	UserId    int64
	Name      *string
	Abilities []string
	Config    *config.Config
}
