package services

import (
	"context"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/entities"
)

type TokenClaims struct {
	UserId string   `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.StandardClaims
}

// mwthod to generate jwt.Token
func GenerateJwtTokenWithClaims(ctx context.Context, user *entities.User) *jwt.Token {
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
