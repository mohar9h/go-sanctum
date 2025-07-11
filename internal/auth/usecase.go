package auth

import (
	"context"
	"time"
)

type Usecase struct {
	storage TokenStorage
}

func (uc *Usecase) Login(userID string, expiresIn *time.Duration) (*Token, error) {
	return GenerateToken(userID, expiresIn)
}

func (uc *Usecase) Logout(token string) error {
	return uc.storage.RevokeToken(context.Background(), token)
}
