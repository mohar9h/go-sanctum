package auth_test

import (
	"github.com/mohar9h/go-sanctum/config"
	"github.com/mohar9h/go-sanctum/storage"
	"testing"

	auth "github.com/mohar9h/go-sanctum"
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

func TestValidateToken(t *testing.T) {

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

	tokenStr, err := auth.CreateToken(&auth.TokenOptions{
		UserId:    1,
		Abilities: []string{"read"},
		DB:        db,
	})
	if err != nil {
		t.Fatal(err)
	}

	cfg := config.DefaultConfig()
	cfg.Storage = storage.NewGormDriver(db)

	token, err := auth.ValidateToken(tokenStr, cfg)
	if err != nil {
		t.Fatalf("Validation failed: %v", err)
	}

	t.Logf("Valid Token: %+v", token)
}

func TestRevokeToken(t *testing.T) {
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

	tokenStr, err := auth.CreateToken(&auth.TokenOptions{
		UserId:    1,
		Abilities: []string{"*"},
		DB:        db, // فرض: اینجا DB از قبل آماده است
	})
	if err != nil {
		t.Fatalf("CreateToken failed: %v", err)
	}

	cfg := config.DefaultConfig()
	cfg.Storage = storage.NewGormDriver(db)
	// باطل‌کردن توکن
	err = auth.RevokeToken(tokenStr, cfg)
	if err != nil {
		t.Fatalf("RevokeToken failed: %v", err)
	}

	// بررسی اعتبارسنجی بعد از ابطال
	_, err = auth.ValidateToken(tokenStr, nil)
	if err == nil {
		t.Fatalf("Token should be revoked but is still valid")
	}

	t.Logf("Token successfully revoked: %v", tokenStr)
}
