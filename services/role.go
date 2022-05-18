package services

import (
	"context"
	"log"

	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/config"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/entities"
)

type RoleService struct{}

var roleService *RoleService = &RoleService{}

func GetRoleService() *RoleService {
	return roleService
}

func (rs *RoleService) Save(ctx context.Context, role *entities.Role) (*entities.Role, error) {
	log.Println("entering method to save new role")

	if err := config.GetDB().Save(role).Error; err != nil {
		log.Println("failed to save new role:", err)
		return nil, err
	}

	log.Println("success save new role")
	return role, nil
}

func (rs *RoleService) GetByName(ctx context.Context, name *string) (*entities.Role, error) {
	log.Println("entering method to get role by name")

	role := &entities.Role{}
	if err := config.GetDB().Where("LOWER(name) = LOWER(?)", *name).First(role).Error; err != nil {
		log.Println("failed to get role by name:", err)
		return nil, err
	}

	log.Println("success get role by name")
	return role, nil
}
