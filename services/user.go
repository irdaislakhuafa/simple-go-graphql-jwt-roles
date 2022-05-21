package services

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/config"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/entities"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph/model"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/tools"
)

type UserService struct{}

type UserServiceInterface interface {
	GetByEmail(ctx context.Context, email *string) (*entities.User, error)
	GetAll() ([]*entities.User, error)
	Save(ctx context.Context, user *entities.User) (*entities.User, error)
	ConvertNewUserToEntityUserWithoutRoles(newUser *model.NewUser) *entities.User
	ConvertEntityUserToModelUser(user *entities.User) *model.User
}

var userService *UserService = &UserService{}

func GetUserService() UserServiceInterface {
	return userService
}

// to save new user
func (us *UserService) Save(ctx context.Context, user *entities.User) (*entities.User, error) {
	log.Println("entering method to save new user")

	if err := config.GetDB().Create(user).Error; err != nil {
		log.Println("failed to save new user:", err)
		return nil, err
	}

	log.Println("success save new user")
	return user, nil
}

// to get user by email with ignore case
func (us *UserService) GetByEmail(ctx context.Context, email *string) (*entities.User, error) {
	log.Println("entering method to get users by email")

	user := &entities.User{}
	if err := config.GetDB().Preload("Roles").Where("LOWER(email) = LOWER(?)", *email).Take(user).Error; err != nil {
		log.Println("failed to get user by email:", err)
		return nil, err
	}

	log.Println("success get user by email")
	return user, nil
}

// to get all user
func (us *UserService) GetAll() ([]*entities.User, error) {
	log.Println("entering method to get all user")

	var users []*entities.User
	if err := config.GetDB().Preload("Roles").Find(&users).Error; err != nil {
		log.Println("failed to get all user:", err)
		return nil, err
	}

	log.Println("success get all user")
	return users, nil
}

// method to convert NewUser to Entity User
func (us *UserService) ConvertNewUserToEntityUserWithoutRoles(newUser *model.NewUser) *entities.User {
	log.Println("entering method to convert NewUser to Entity User")
	user := &entities.User{
		ID:    uuid.NewString(),
		Name:  newUser.Name,
		Email: newUser.Email,
		Password: func() string {
			hashedPassword, err := tools.HashPassword(&newUser.Password)
			if err != nil {
				log.Println(err)
				return newUser.Password
			}
			return *hashedPassword
		}(),
	}
	log.Println("success convert")
	return user
}

// method to convert Entity User to Model User
func (us *UserService) ConvertEntityUserToModelUser(user *entities.User) *model.User {
	log.Println("entering method to convert Entity User to Model User")
	modelUser := &model.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Roles: func() []string {
			var roles []string
			for _, v := range user.Roles {
				roles = append(roles, v.Name)
			}
			return roles
		}(),
	}
	log.Println("success convert")
	return modelUser
}

// method to get user by id
func (us *UserService) GetUserByID(ctx context.Context, userId *string) (*entities.User, error) {
	log.Println("entering method to get user by id")

	user := &entities.User{}
	if err := config.GetDB().Where("id = ?", *userId).Take(user).Error; err != nil {
		log.Println("failed to get user by id:", err)
		return nil, err
	}

	log.Println("success get user by id")
	return user, nil
}
