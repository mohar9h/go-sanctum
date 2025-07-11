package db

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/mohar9h/go-sanctum/internal/auth"
)

type Storage struct {
	db *gorm.DB
}

func NewDBStorage(db *gorm.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) StoreToken(ctx context.Context, token auth.Token) error {
	model := TokenModel{
		UserID:     token.UserID,
		TokenHash:  token.Hash,
		CreatedAt:  token.CreatedAt,
		LastUsedAt: token.LastUsedAt,
		ExpiresAt:  token.ExpiresAt,
	}
	return s.db.WithContext(ctx).Create(&model).Error
}

func (s *Storage) ValidateToken(ctx context.Context, plaintext string) (string, error) {
	var tokens []TokenModel
	if err := s.db.WithContext(ctx).Find(&tokens).Error; err != nil {
		return "", err
	}

	for _, t := range tokens {
		valid, _ := auth.VerifyToken(plaintext, t.TokenHash)
		if valid {
			if t.ExpiresAt != nil && time.Now().After(*t.ExpiresAt) {
				return "", auth.ErrTokenExpired
			}
			// Update last_used_at
			t.LastUsedAt = time.Now()
			_ = s.db.WithContext(ctx).Save(&t).Error
			return t.UserID, nil
		}
	}

	return "", auth.ErrTokenInvalid
}

func (s *Storage) RevokeToken(ctx context.Context, plaintext string) error {
	var tokens []TokenModel
	if err := s.db.WithContext(ctx).Find(&tokens).Error; err != nil {
		return err
	}

	for _, t := range tokens {
		valid, _ := auth.VerifyToken(plaintext, t.TokenHash)
		if valid {
			return s.db.WithContext(ctx).Delete(&t).Error
		}
	}

	return errors.New("token not found")
}
