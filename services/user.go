package services

import (
	"context"
	"log"

	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/config"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/entities"
)

type UserService struct{}

type UserServiceInterface interface {
	GetByEmail(ctx context.Context, email *string) (*entities.User, error)
	GetAll() ([]*entities.User, error)
	Save(ctx context.Context, user *entities.User) (*entities.User, error)
}

var userService *UserService = &UserService{}

func GetUserService() UserServiceInterface {
	return userService
}

func (us *UserService) Save(ctx context.Context, user *entities.User) (*entities.User, error) {
	log.Println("entering method to save new user")

	if err := config.GetDB().Save(user).Error; err != nil {
		log.Println("failed to save new user:", err)
		return nil, err
	}

	log.Println("success save new user")
	return user, nil
}
