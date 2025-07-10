package auth

import "time"

type Token struct {
	ID         string     `json:"id"`
	Plaintext  string     `json:"plaintext"`
	Hash       []byte     `json:"-"`
	UserID     string     `json:"user_id"`
	IP         string     `json:"ip"`
	UserAgent  string     `json:"user_agent"`
	ExpiresAt  *time.Time `json:"expires_at"`
	CreatedAt  time.Time  `json:"created_at"`
	LastUsedAt time.Time  `json:"last_used_at"`
}
