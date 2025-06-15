package helper

import (
	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword takes a plain text password as input and returns its bcrypt hashed representation as a string.
// It uses the default cost provided by the bcrypt package for hashing.
// Note: The function ignores errors returned by bcrypt.GenerateFromPassword.
func GeneratePassword(password string) string {
	bytePass := []byte(password)
	pass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	return string(pass)
}

// CheckPassword compares a plaintext password with a hashed password and returns true if they match.
// It uses bcrypt to perform the comparison.
// Parameters:
//   - password: the plaintext password to verify.
//   - hash: the bcrypt hashed password to compare against.
//
// Returns:
//   - bool: true if the password matches the hash, false otherwise.
func CheckPassword(password string, hash string) bool {
	bytePass := []byte(password)
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)
	return err == nil
}
