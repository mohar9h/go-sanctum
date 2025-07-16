package storage

import "time"

// Token APIToken defines the persistent structure stored in SQL/Redis.
type Token struct {
	ID        int64      // ID
	UserId    int64      // Owner
	Token     string     // Token before hash
	Name      *string    // Device or client name
	Abilities string     // Comma-separated abilities
	CreatedAt time.Time  // Issued at
	ExpiresAt *time.Time // Optional expiration
}
