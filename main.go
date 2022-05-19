package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/config"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph/generated"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/services"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	// load .env file
	godotenv.Load(".env")

	// initialize database
	config.InitDB()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = defaultPort
	}

	// graphql resolver config
	resolverConfig := &generated.Config{
		Resolvers: &graph.Resolver{
			RoleService: services.GetRoleService(),
			UserService: services.GetUserService(),
		},
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(*resolverConfig))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
