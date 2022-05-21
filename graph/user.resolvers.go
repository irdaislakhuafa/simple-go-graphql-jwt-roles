package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph/generated"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph/model"
)

func (r *userOptionsResolver) GetAll(ctx context.Context, obj *model.UserOptions) ([]*model.User, error) {
	// get all user
	users, err := r.UserService.GetAll()
	if err != nil {
		return nil, err
	}

	// create array of struct model.User to return response
	var modelUsers []*model.User
	for _, v := range users {
		// convert Entity User to Model User
		modelUsers = append(modelUsers, r.UserService.ConvertEntityUserToModelUser(v))
	}

	return modelUsers, nil
}

func (r *userOptionsResolver) GetByID(ctx context.Context, obj *model.UserOptions, userID string) (*model.User, error) {
	// get user by id
	user, err := r.UserService.GetUserByID(ctx, &userID)
	if err != nil {
		return nil, err
	}

	modelUser := r.UserService.ConvertEntityUserToModelUser(user)
	return modelUser, nil
}

// UserOptions returns generated.UserOptionsResolver implementation.
func (r *Resolver) UserOptions() generated.UserOptionsResolver { return &userOptionsResolver{r} }

type userOptionsResolver struct{ *Resolver }
