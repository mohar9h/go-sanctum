package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/mohar9h/go-sanctum/internal/auth"
)

type contextKey string

const (
	UserIDKey contextKey = "userID"
)

func Authenticate(tokenStore auth.TokenStorage) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			// 1. بررسی توکن API از هدر Authorization
			if authHeader := r.Header.Get("Authorization"); authHeader != "" {
				if strings.HasPrefix(authHeader, "Bearer ") {
					token := strings.TrimPrefix(authHeader, "Bearer ")
					if userID, err := tokenStore.ValidateToken(r.Context(), token); err == nil {
						ctx = context.WithValue(ctx, UserIDKey, userID)
						next.ServeHTTP(w, r.WithContext(ctx))
						return
					}
				}
			}

			// 2. بررسی توکن از کوکی (برای برنامه‌های وب)
			if cookie, err := r.Cookie("sanctum_token"); err == nil {
				if userID, err := tokenStore.ValidateToken(r.Context(), cookie.Value); err == nil {
					ctx = context.WithValue(ctx, UserIDKey, userID)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
			}

			// 3. عدم احراز هویت
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		})
	}
}
