package db

import (
	"time"
)

type TokenModel struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     string    `gorm:"index;not null"`
	TokenHash  []byte    `gorm:"column:token_hash;uniqueIndex;not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	LastUsedAt time.Time `gorm:"not null"`
	ExpiresAt  *time.Time
}
