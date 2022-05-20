package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/entities"
)

// var secretKey []byte = []byte(*getSecretKey())

func getSecretKey() []byte {
	envSecretKey := os.Getenv("APP_SECRET_KEY")
	if envSecretKey == "" {
		log.Println("APP_SECRET_KEY is empty or not valid, using default secret key!")
		return []byte("default_secret")
	}
	return []byte(envSecretKey)
}

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

			StandardClaims: jwt.StandardClaims{
				// TODO: add IssueAt and ExpiredAt
				IssuedAt: time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Minute * (func() time.Duration {
					expiredInMinute, err := strconv.Atoi(os.Getenv("APP_TOKEN_EXPIRED_IN_MINUTE"))
					if err != nil {
						log.Println("APP_TOKEN_EXPIRED_IN_MINUTE is not valid, using default value 5 minute")
						return time.Duration(5)
					}
					return time.Duration(expiredInMinute)
				}())).Unix(),
			},
		},
	)
	log.Println("success generate jwt.Token with claims")
	return jwtToken
}

// to generate jwt token string
func GenerateTokenString(ctx context.Context, user *entities.User) (*string, error) {
	log.Println("entering method to generate token string")

	jwtToken := generateJwtTokenWithClaims(ctx, user)
	tokenStringm, err := jwtToken.SignedString(getSecretKey())
	if err != nil {
		log.Println("failed to generate token string:", err)
		return nil, err
	}

	log.Println("success to generate token string")
	return &tokenStringm, nil
}

// method to validate token
func ValidateTokenString(ctx context.Context, tokenString *string) (*jwt.Token, error) {
	log.Println("entering method to validate token string")

	jwtToken, err := jwt.ParseWithClaims(*tokenString, &TokenClaims{}, keyFunc)
	if err != nil {
		log.Println("failed to validate token:", err)
		return nil, err
	}

	log.Println("success validate token string")
	return jwtToken, nil
}

func keyFunc(jwtToken *jwt.Token) (any, error) {
	if _, isOk := jwtToken.Method.(*jwt.SigningMethodHMAC); !isOk {
		return nil, fmt.Errorf("signing method is not valid")
	}
	return getSecretKey(), nil
}

// to get claims from jwt.Token
func GetAllClaimsFromJwtToken(ctx context.Context, jwtToken *jwt.Token) (*TokenClaims, error) {
	log.Println("entering method to get all claims from jwt.Token")

	claims, isOk := jwtToken.Claims.(*TokenClaims)
	if !isOk {
		errMessage := "failed to get all claims from jwt.Token, token is invalid"
		log.Println(errMessage)
		return nil, fmt.Errorf(errMessage)
	}

	log.Println("success get all claims from jwt.Token")
	return claims, nil
}
