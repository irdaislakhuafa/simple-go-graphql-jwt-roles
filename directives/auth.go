package directives

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/keys"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/services"
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
		"errTokenInvalid",  // for check is token invalid
		"errClaimsInvalid", // for check is claim invalid
	}
	for _, value := range errErrors {
		if v := ctx.Value(keys.String(value)); v != nil {
			return nil, v.(error)
		}
	}

	// get claims from context
	var claims *services.TokenClaims
	if v := ctx.Value(keys.String("claims")); v != nil {
		claims = v.(*services.TokenClaims)
	}

	// checking roles
	err := ad.CheckRoles(roles, claims.Roles)
	if err != nil {
		return nil, err
	}

	// direct to next resolvers
	return next(ctx)
}

func (ad *authDirective) CheckRoles(expectedRoles []string, actualRoles []string) error {
	log.Println("checking roles")

	// if dont have any role
	if len(expectedRoles) <= 0 {
		log.Println("don't have any roles")
		return nil
	}

	var err error
	isPermited := false
	for i := 0; i < len(expectedRoles); i++ {
		for j := 0; j < len(actualRoles); j++ {
			log.Println("Expected:", actualRoles[j], ">>", "Actual:", expectedRoles[i], "=>", actualRoles[j] == expectedRoles[i])

			if strings.EqualFold(actualRoles[j], expectedRoles[i]) {
				log.Println("role", strings.ToUpper(actualRoles[j]), "is allowed")
				err = nil
				isPermited = true
				break
			} else {
				err = fmt.Errorf("access denied, role %s not allowed", strings.ToUpper(actualRoles[j]))
				isPermited = false
				log.Println(err.Error())
			}
		}

		// is permited
		if isPermited {
			break
		}
	}

	return err
}
