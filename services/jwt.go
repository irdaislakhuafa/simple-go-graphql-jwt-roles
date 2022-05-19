package services

import (
	"context"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/entities"
)

var secretKey []byte = []byte(func() string {
	envSecretKey := os.Getenv("APP_SECRET_KEY")
	if envSecretKey == "" {
		log.Println("APP_SECRET_KEY is empty or not valid, using default secret key!")
		return "default_secret"
	}
	return envSecretKey
}())

type TokenClaims struct {
	UserId string   `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.StandardClaims
}

// mwthod to generate jwt.Token
func generateJwtTokenWithClaims(ctx context.Context, user *entities.User) *jwt.Token {
	log.Println("entering method to generate jwt.Token with claims")
	jwtToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&TokenClaims{
			UserId: user.ID,
			Roles: func() (roles []string) {
				for _, v := range user.Roles {
					roles = append(roles, v.Name)
				}
				return roles
			}(),
		},
	)
	log.Println("success generate jwt.Token with claims")
	return jwtToken
}

// to generate jwt token string
func GenerateTokenString(ctx context.Context, user *entities.User) (*string, error) {
	log.Println("entering method to generate token string")

	jwtToken := generateJwtTokenWithClaims(ctx, user)
	tokenStringm, err := jwtToken.SignedString(secretKey)
	if err != nil {
		log.Println("failed to generate token string:", err)
		return nil, err
	}

	log.Println("success to generate token string")
	return &tokenStringm, nil
}
