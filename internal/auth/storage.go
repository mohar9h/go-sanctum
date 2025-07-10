package auth

import "context"

// TokenStorage رابط ذخیره‌سازی و مدیریت توکن‌ها
type TokenStorage interface {
	StoreToken(ctx context.Context, token Token) error
	ValidateToken(ctx context.Context, plaintext string) (string, error)
	RevokeToken(ctx context.Context, token string) error
}
