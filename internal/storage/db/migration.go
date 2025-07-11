package db

import (
	"fmt"
	"gorm.io/gorm"
)

func MigrateSanctumTables(db *gorm.DB) error {
	err := db.AutoMigrate(&TokenModel{})
	if err != nil {
		return fmt.Errorf("failed to migrate token table: %w", err)
	}
	return nil
}
