package services

import (
	"context"
	"log"

	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph/model"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/tools"
	"gorm.io/gorm"
)

type AuthService struct{}
type AuthServiceInterface interface {
	Login(ctx context.Context, loginUser *model.LoginUser) (*model.ResponseToken, error)
	Register(ctx context.Context, newUser *model.NewUser) (*model.ResponseToken, error)
}

var authService *AuthService = &AuthService{}

func GetAuthService() AuthServiceInterface {
	return authService
}

// method to login
func (as *AuthService) Login(ctx context.Context, loginUser *model.LoginUser) (*model.ResponseToken, error) {
	log.Println("entering method to login user")

	// check is user exists?
	us := GetUserService()
	user, err := us.GetByEmail(ctx, &loginUser.Email)
	if err != nil || user == nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}

	// check password
	if isOk, err := tools.CompareHashAndReal(&user.Password, &loginUser.Password); !isOk || err != nil {
		return nil, err
	}

	// generate token string
	tokenString, err := GenerateTokenString(ctx, user)
	if err != nil {
		return nil, err
	}

	// return token
	log.Println("success login")
	return &model.ResponseToken{Token: *tokenString}, nil
}
