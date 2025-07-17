package storage

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type gormDriver struct {
	db *gorm.DB
}

func NewGormDriver(db *gorm.DB) Driver {
	return &gormDriver{db: db}
}

func (g *gormDriver) StoreToken(t *Token) error {
	return g.db.Create(t).Error
}

func (g *gormDriver) FindByID(id int64) (*Token, error) {
	var token Token
	if err := g.db.First(&token, "id = ?", id).Error; err != nil {
		return nil, err
	}
	if token.ExpiresAt != nil && time.Now().After(*token.ExpiresAt) {
		return nil, errors.New("token expired")
	}
	return &token, nil
}

func (g *gormDriver) FindByHash(hash string) (*Token, error) {
	var token Token
	if err := g.db.First(&token, "token = ?", hash).Error; err != nil {
		return nil, err
	}
	if token.ExpiresAt != nil && time.Now().After(*token.ExpiresAt) {
		return nil, errors.New("token expired")
	}
	return &token, nil
}

func (g *gormDriver) RevokeToken(id string) error {
	return g.db.Delete(&Token{}, "id = ?", id).Error
}

func (g *gormDriver) TouchLastUsed(hash string) error {
	return g.db.Model(&Token{}).
		Where("token = ?", hash).
		Update("last_used", time.Now()).
		Error
}
