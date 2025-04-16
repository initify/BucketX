package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(token string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	return string(hash)
}

func VerifyToken(token string, authToken string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(token), []byte(authToken))
	return err == nil
}
