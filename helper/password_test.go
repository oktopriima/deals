package helper

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestGeneratePassword(t *testing.T) {
	password := "mySecret123"
	hashed := GeneratePassword(password)

	if hashed == "" {
		t.Fatal("expected hashed password, got empty string")
	}

	// Check that the hash matches the original password
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		t.Errorf("hashed password does not match original password: %v", err)
	}
}

func TestGeneratePassword_DifferentHashes(t *testing.T) {
	password := "samePassword"
	hash1 := GeneratePassword(password)
	hash2 := GeneratePassword(password)

	if hash1 == hash2 {
		t.Error("expected different hashes for the same password due to salting")
	}
}
func TestCheckPassword_CorrectPassword(t *testing.T) {
	password := "correctPassword"
	hashed := GeneratePassword(password)

	if !CheckPassword(password, hashed) {
		t.Error("expected CheckPassword to return true for correct password")
	}
}

func TestCheckPassword_WrongPassword(t *testing.T) {
	password := "correctPassword"
	wrongPassword := "wrongPassword"
	hashed := GeneratePassword(password)

	if CheckPassword(wrongPassword, hashed) {
		t.Error("expected CheckPassword to return false for wrong password")
	}
}

func TestCheckPassword_InvalidHash(t *testing.T) {
	password := "somePassword"
	invalidHash := "notAValidHash"

	if CheckPassword(password, invalidHash) {
		t.Error("expected CheckPassword to return false for invalid hash")
	}
}
