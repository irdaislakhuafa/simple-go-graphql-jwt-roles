package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph/generated"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph/model"
)

func (r *userOptionsResolver) GetAll(ctx context.Context, obj *model.UserOptions) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserOptions returns generated.UserOptionsResolver implementation.
func (r *Resolver) UserOptions() generated.UserOptionsResolver { return &userOptionsResolver{r} }

type userOptionsResolver struct{ *Resolver }
