package auth_test

import (
	"testing"

	auth "github.com/mohar9h/go-sanctum"
	"github.com/mohar9h/go-sanctum/storage"
)

func TestCreateRandomToken(t *testing.T) {
	db, err := auth.OpenPostgres(auth.PostgresConfig{
		Host:            "localhost",
		Port:            5432,
		Username:        "postgres",
		Password:        "Mohammad1367",
		Database:        "pecalets",
		SSLMode:         "disable",
		MaxIdleConns:    15,
		MaxOpenConns:    100,
		ConnMaxLifetime: 5,
	})
	if err != nil {
		t.Fatalf("failed to connect to PostgreSQL: %v", err)
	}

	// مهاجرت جدول مورد نیاز برای توکن‌ها
	if err := db.AutoMigrate(&storage.Token{}); err != nil {
		t.Fatalf("failed to migrate token table: %v", err)
	}

	// فراخوانی تابع تولید توکن
	result, err := auth.CreateToken(&auth.TokenOptions{
		UserId:    1,
		Name:      nil,
		Abilities: []string{"*"},
		DB:        db,
	})

	if err != nil {
		t.Fatalf("CreateToken failed: %v", err)
	}

	// خروجی توکن لاگ شود
	t.Logf("Generated Token: %s", result)
}
