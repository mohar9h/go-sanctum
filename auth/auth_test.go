package auth_test

import (
	auth "github.com/mohar9h/go-sanctum"
	"testing"
)

func TestCreateRandomToken(t *testing.T) {

	result, err := auth.CreateToken(&auth.TokenOptions{
		UserId:    1,
		Name:      nil,
		Abilities: []string{"*"},
	})

	if err != nil {
		t.Fatalf("CreateToken failed: %v", err)
	}

	t.Logf("Generated Token: %+v", result)
}
