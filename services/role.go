package services

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/config"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/entities"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph/model"
)

type RoleService struct{}

var roleService *RoleService = &RoleService{}

func GetRoleService() *RoleService {
	return roleService
}

// to save new role
func (rs *RoleService) Save(ctx context.Context, role *entities.Role) (*entities.Role, error) {
	log.Println("entering method to save new role")

	if err := config.GetDB().Save(role).Error; err != nil {
		log.Println("failed to save new role:", err)
		return nil, err
	}

	log.Println("success save new role")
	return role, nil
}

// method to get role by name with ignore case
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

// to get all roles
func (rs *RoleService) GetAll() ([]*entities.Role, error) {
	log.Println("entering method to get all roles")

	var roles []*entities.Role
	if err := config.GetDB().Find(&roles).Error; err != nil {
		log.Println("failed to get all roles:", err)
		return nil, err
	}

	log.Println("success get all roles")
	return roles, nil
}

// method to convert Entity Role to Model Role
func (rs *RoleService) ConvertEntityRoleToModelRole(role *entities.Role) *model.Role {
	log.Println("entering method to convert Entity Role to Model Role")
	modelRole := &model.Role{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
	}
	log.Println("success convert")
	return modelRole
}

// method to convert NewRole to Entity Role
func (rs *RoleService) ConvertNewRoleToEntityRole(newRole *model.NewRole) *entities.Role {
	log.Println("entering method to convert NewRole to Entity Role")
	role := &entities.Role{
		ID:          uuid.NewString(),
		Name:        newRole.Name,
		Description: newRole.Description,
	}
	log.Println("success convert")
	return role
}
