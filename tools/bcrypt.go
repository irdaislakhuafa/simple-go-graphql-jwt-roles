package tools

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// for hashing password before save password to database
func HashPassword(password *string) (*string, error) {
	log.Println("entering method to hashing password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error while hashing password")
		return password, err
	}

	log.Println("success hashed password")
	stringHash := string(hashedPassword)
	return &stringHash, nil
}
