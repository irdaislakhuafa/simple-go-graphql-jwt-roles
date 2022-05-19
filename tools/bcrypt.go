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

// to compare between real password and hashed password
func CompareHashAndReal(hashedPassword, realPassword *string) bool {
	log.Println("entering method to compare hashed password and real password")
	err := bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(*realPassword))
	if err != nil {
		log.Println("failed to compare password:", err)
		return false
	}

	log.Println("success compare password")
	return true
}
