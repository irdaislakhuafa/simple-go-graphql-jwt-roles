package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/keys"
)

type authDirective struct{}
type AuthDirectiveInterface interface {
	AuthDirective(ctx context.Context, object any, next graphql.Resolver, roles []string) (any, error)
}

var authDirectiveVal *authDirective = &authDirective{}

func GetAuth() AuthDirectiveInterface {
	return authDirectiveVal
}

// TODO : implement @auth(roles [String!]!) directive
func (ad *authDirective) AuthDirective(ctx context.Context, object any, next graphql.Resolver, roles []string) (any, error) {

	errStringPointers := []string{
		"errAuthEmpty",  // for check is `errAuthEmpty` exists
		"errTokenEmpty", // for check is `errTokenEmpty` exists
	}
	for _, value := range errStringPointers {
		if v := ctx.Value(keys.String(value)); v != nil {
			return nil, fmt.Errorf(*(v.(*string)))
		}
	}

	errErrors := []string{
		"errTokenInvalid", // for check is token invalid
	}
	for _, value := range errErrors {
		if v := ctx.Value(keys.String(value)); v != nil {
			return nil, v.(error)
		}
	}

	panic("not implemented")
}
