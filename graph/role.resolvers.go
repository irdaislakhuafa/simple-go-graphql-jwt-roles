package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph/generated"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph/model"
)

func (r *roleMutationOptionsResolver) Save(ctx context.Context, obj *model.RoleMutationOptions, newRole model.NewRole) (*model.Role, error) {
	// convert to role
	role := r.RoleService.ConvertNewRoleToEntityRole(&newRole)

	// save new role
	role, err := r.RoleService.Save(ctx, role)
	if err != nil {
		return nil, err
	}

	// convert to model
	modelRole := r.RoleService.ConvertEntityRoleToModelRole(role)
	return modelRole, nil
}

func (r *roleQueryOptionsResolver) GetAll(ctx context.Context, obj *model.RoleQueryOptions) ([]*model.Role, error) {
	// get all roles
	roles, err := r.RoleService.GetAll()
	if err != nil {
		return nil, err
	}

	// convert to model role
	var modelRoles []*model.Role
	for _, v := range roles {
		modelRoles = append(modelRoles, r.RoleService.ConvertEntityRoleToModelRole(v))
	}

	return modelRoles, nil
}

// RoleMutationOptions returns generated.RoleMutationOptionsResolver implementation.
func (r *Resolver) RoleMutationOptions() generated.RoleMutationOptionsResolver {
	return &roleMutationOptionsResolver{r}
}

// RoleQueryOptions returns generated.RoleQueryOptionsResolver implementation.
func (r *Resolver) RoleQueryOptions() generated.RoleQueryOptionsResolver {
	return &roleQueryOptionsResolver{r}
}

type roleMutationOptionsResolver struct{ *Resolver }
type roleQueryOptionsResolver struct{ *Resolver }
