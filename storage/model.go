package storage

import "time"

// Token APIToken defines the persistent structure stored in SQL/Redis.
type Token struct {
	ID        string     // Hashed ID
	UserID    string     // Owner
	Name      string     // Device or client name
	Abilities string     // Comma-separated abilities
	CreatedAt time.Time  // Issued at
	ExpiresAt *time.Time // Optional expiration
}
