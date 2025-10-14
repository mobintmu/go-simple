package test

import (
	"fmt"
	"go-simple/internal/auth"
	"go-simple/internal/config"
	"testing"
)

func TestJWTTokenCreator(t *testing.T) {
	// t.Parallel()
	WithHttpTestServer(t, func() {
		cfg, err := config.NewConfig()
		if err != nil {
			t.Fatalf("Failed to load config: %v", err)
		}
		adminID := "admin-generated-id"
		token, err := auth.GenerateToken(cfg, adminID)
		if err != nil {
			t.Fatalf("Failed to generate token: %v", err)
		}
		if token == "" {
			t.Fatal("Generated token is empty")
		}
		fmt.Printf("✅ Generated JWT Token: \n\n%s\n\n✅ Generated JWT Token", token)
	})
}
