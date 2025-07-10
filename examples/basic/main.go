package main

import (
	"fmt"
	"github.com/mohar9h/go-sanctum/internal/auth"
	"github.com/mohar9h/go-sanctum/internal/storage"
	"github.com/mohar9h/go-sanctum/pkg/http/middleware"
	"net/http"
)

func main() {
	// 1. ایجاد ذخیره‌سازی
	store := storage.NewMemoryStorage()

	// 2. ایجاد روتر
	mux := http.NewServeMux()

	// 3. افزودن میدلور احراز هویت
	authMiddleware := middleware.Authenticate(store)

	// 4. تعریف مسیرها
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// در واقعیت باید کاربر را از دیتابیس بیابید
		//expiration := 24 * time.Hour
		token, err := auth.GenerateToken("user123", nil)
		if err != nil {
			http.Error(w, "Token generation failed", http.StatusInternalServerError)
			return
		}

		// ذخیره توکن
		if err := store.StoreToken(r.Context(), *token); err != nil {
			http.Error(w, "Token storage failed", http.StatusInternalServerError)
			return
		}

		// برگرداندن توکن به کاربر
		_, err = fmt.Fprintf(w, "Your token: %s", token.Plaintext)
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	})

	mux.Handle("/protected", authMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			userID := r.Context().Value(middleware.UserIDKey).(string)
			_, err := fmt.Fprintf(w, "Hello %s! This is protected content.", userID)
			if err != nil {
				http.Error(w, "Failed to write response", http.StatusInternalServerError)
				return
			}
		},
	)))

	// 5. راه‌اندازی سرور
	fmt.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
