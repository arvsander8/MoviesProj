package main

import (
	"testing"
	"time"
)

// Mock data for jwtUser
var testUser = &jwtUser{
	ID:        1,
	FirstName: "John",
	LastName:  "Doe",
}

// TestGenerateTokenPair tests the GenerateTokenPair method for success
func TestGenerateTokenPair(t *testing.T) {
	auth := Auth{
		Issuer:        "testIssuer",
		Audience:      "testAudience",
		Secret:        "testSecret",
		TokenExpiry:   time.Hour,
		RefreshExpiry: time.Hour * 24,
	}

	tokenPairs, err := auth.GenerateTokenPair(testUser)
	if err != nil {
		t.Fatalf("GenerateTokenPair() error = %v", err)
	}
	if tokenPairs.Token == "" || tokenPairs.RefreshToken == "" {
		t.Errorf("Expected non-empty tokens, got %v", tokenPairs)
	}

	// Additional checks can be made here, such as verifying token structure or claims
}
