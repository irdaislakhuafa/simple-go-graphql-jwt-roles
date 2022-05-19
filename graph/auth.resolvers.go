package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph/generated"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph/model"
)

func (r *authOptionsResolver) Register(ctx context.Context, obj *model.AuthOptions, newUser model.NewUser) (*model.ResponseToken, error) {
	return r.AuthService.Register(ctx, &newUser)
}

func (r *authOptionsResolver) Login(ctx context.Context, obj *model.AuthOptions, user *model.LoginUser) (*model.ResponseToken, error) {
	return r.AuthService.Login(ctx, user)
}

func (r *mutationResolver) Auth(ctx context.Context) (*model.AuthOptions, error) {
	return &model.AuthOptions{}, nil
}

func (r *mutationResolver) Role(ctx context.Context) (*model.RoleMutationOptions, error) {
	return &model.RoleMutationOptions{}, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.UserOptions, error) {
	return &model.UserOptions{}, nil
}

func (r *queryResolver) Role(ctx context.Context) (*model.RoleQueryOptions, error) {
	return &model.RoleQueryOptions{}, nil
}

// AuthOptions returns generated.AuthOptionsResolver implementation.
func (r *Resolver) AuthOptions() generated.AuthOptionsResolver { return &authOptionsResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type authOptionsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
