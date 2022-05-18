package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph/generated"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph/model"
)

func (r *roleMutationOptionsResolver) Save(ctx context.Context, obj *model.RoleMutationOptions, newRole model.NewRole) (*model.Role, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *roleQueryOptionsResolver) GetAll(ctx context.Context, obj *model.RoleQueryOptions) ([]*model.Role, error) {
	panic(fmt.Errorf("not implemented"))
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
