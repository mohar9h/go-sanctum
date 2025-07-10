package config

import "time"

type SanctumConfig struct {
	TokenExpiration *time.Duration `mapstructure:"token_expiration"`
	TokenPrefix     string         `mapstructure:"token_prefix"`
	CookieName      string         `mapstructure:"cookie_name"`
	CookieSecure    bool           `mapstructure:"cookie_secure"`
	CookieHTTPOnly  bool           `mapstructure:"cookie_http_only"`
	CookieSameSite  string         `mapstructure:"cookie_same_site"`
	TokenLength     int            `mapstructure:"token_length"`
	CSRFEnabled     bool           `mapstructure:"csrf_enabled"`
	CSRFHeader      string         `mapstructure:"csrf_header"`
	StorageDriver   string         `mapstructure:"storage_driver"`
}

func DefaultConfig() *SanctumConfig {
	return &SanctumConfig{
		TokenExpiration: nil,
		TokenPrefix:     "Bearer",
		CookieName:      "sanctum_token",
		CookieSecure:    true,
		CookieHTTPOnly:  true,
		CookieSameSite:  "Lax",
		TokenLength:     64,
		CSRFEnabled:     true,
		CSRFHeader:      "X-CSRF-TOKEN",
		StorageDriver:   "memory",
	}
}
