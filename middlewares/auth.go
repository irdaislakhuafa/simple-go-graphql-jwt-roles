package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/keys"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/services"
)

// TODO : implement AuthMiddleware
func AuthMiddleware(next http.Handler) http.Handler {
	log.Println("entering middleware section")

	// create handler function
	handlerFunc := func(writer http.ResponseWriter, request *http.Request) {

		// constant/final variables
		const (
			HEADER_AUTHORIZATION = "Authorization"
			BEARER               = "Bearer "
		)

		log.Println("get `Authorization` header from request")
		auth := strings.Trim(request.Header.Get(HEADER_AUTHORIZATION), " ")
		// if auth is empty
		if len(auth) <= 0 || auth == "" || len(auth) <= len(BEARER) {
			errMessage := fmt.Sprintf("`%s` is empty or not valid", HEADER_AUTHORIZATION)
			log.Println(errMessage)
			request = overrideRequest(request, "errAuthEmpty", &errMessage)
			next.ServeHTTP(writer, request)
			return
		}

		// get bearer token string from HEADER_AUTHORIZATION
		tokenString := auth[len(BEARER):]
		if tokenString == "" { // if empty
			errMessage := "token string is empty or not valid"
			log.Println(errMessage)
			request = overrideRequest(request, "errTokenEmpty", &errMessage)
			return
		}

		// TODO: check token string validation
		jwtToken, err := services.ValidateTokenString(context.Background(), &tokenString)
		if err != nil || !jwtToken.Valid {
			log.Println("token string is not valid:", err)
			request = overrideRequest(request, "errTokenInvalid", err)
			next.ServeHTTP(writer, request)
			return
		}

		// get claims
		claims, err := services.GetAllClaimsFromJwtToken(context.Background(), jwtToken)
		if err != nil {
			log.Println(err.Error())
			request = overrideRequest(request, "errClaimsInvalid", claims)
			next.ServeHTTP(writer, request)
			return
		}

		// TODO: insert token claims to context
		request = overrideRequest(request, "claims", claims)
		next.ServeHTTP(writer, request)
	}

	// return handler func
	return http.HandlerFunc(handlerFunc)
}

// method to override in comming request
func overrideRequest[T any](request *http.Request, key string, value T) *http.Request {
	log.Println("overriding request with new context")
	newCtx := context.WithValue(request.Context(), keys.String(key), value)
	request = request.WithContext(newCtx)
	return request
}
