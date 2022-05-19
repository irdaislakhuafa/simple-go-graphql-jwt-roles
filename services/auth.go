package services

import (
	"context"
	"fmt"
	"log"
	"strings"

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
	if err != nil {
		return nil, err
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

// method to register new user
func (as *AuthService) Register(ctx context.Context, newUser *model.NewUser) (*model.ResponseToken, error) {
	log.Println("entering method to register new user")

	// check is user already exists
	us := GetUserService()
	rs := GetRoleService()
	user, err := us.GetByEmail(ctx, &newUser.Email)
	if err != nil || user == nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}

	// convert to user entities
	user = us.ConvertNewUserToEntityUserWithoutRoles(newUser)

	// get roles by names
	for _, v := range newUser.Roles {
		role, err := rs.GetByName(context.Background(), &v)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("role \"%s\" not found!", strings.ToUpper(v))
			}
			return nil, err
		}
		user.Roles = append(user.Roles, *role) // FIXME: why this error
	}

	// save user and roles in bridge table
	// TODO: how save user_roles

	// save new user
	user, err = us.Save(ctx, user)
	if err != nil || user == nil {
		return nil, err
	}

	// generate token
	tokenString, err := GenerateTokenString(ctx, user)
	if err != nil {
		return nil, err
	}

	// return token
	log.Println("success register")
	return &model.ResponseToken{Token: *tokenString}, nil
}
