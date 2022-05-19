package graph

import "github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	RoleService *services.RoleService
	UserService services.UserServiceInterface
}
