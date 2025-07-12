package token

// Claims defines the payload for JWT token.
type Claims struct {
	UserId    string   `json:"user_id"`
	Name      string   `json:"name"`
	Abilities []string `json:"abilities"`
	ExpireAt  int64    `json:"expire_at,omitempty"`
	IssueAt   int64    `json:"issue_at"`
}
