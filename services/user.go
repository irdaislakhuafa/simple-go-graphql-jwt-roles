package services

import (
	"context"
	"log"

	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/config"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/entities"
)

type UserService struct{}

type UserServiceInterface[T any] interface {
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

func (us *UserService) GetByEmail(ctx context.Context, email *string) (*entities.User, error) {
	log.Println("entering method to get users by email")

	user := &entities.User{}
	if err := config.GetDB().Where("LOWER(email) = LOWER(?)", *email).Take(user).Error; err != nil {
		log.Println("failed to get user by email:", err)
		return nil, err
	}

	log.Println("success get user by email")
	return user, nil
}

func (us *UserService) GetAll() ([]*entities.User, error) {
	log.Println("entering method to get all user")

	var users []*entities.User
	if err := config.GetDB().Find(&users).Error; err != nil {
		log.Println("failed to get all user:", err)
		return nil, err
	}

	log.Println("success get all user")
	return users, nil
}