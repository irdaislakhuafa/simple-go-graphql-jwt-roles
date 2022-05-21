package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/config"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/directives"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/graph/generated"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/middlewares"
	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/services"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	// app mode
	var mode string
	flag.StringVar(&mode, "app-mode", "dev", "Put app mode here, use DEV for development mode and PROD for production")
	flag.Parse()

	// load .env file
	if strings.EqualFold(mode, "prod") {
		godotenv.Load("env/prod.env")
	} else if strings.EqualFold(mode, "dev") {
		godotenv.Load("env/dev.env")
	} else {
		log.Println("app-mode not valid, using default mode DEV")
		godotenv.Load("env/dev.env")
	}

	// initialize database
	config.InitDB()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = defaultPort
	}

	// initialize mux router
	router := mux.NewRouter()

	// use auth middleware
	router.Use(middlewares.AuthMiddleware)

	// graphql resolver config
	resolverConfig := &generated.Config{
		Resolvers: &graph.Resolver{
			RoleService: services.GetRoleService(),
			UserService: services.GetUserService(),
			AuthService: services.GetAuthService(),
		},
	}

	// implement @auth(roles: [String!]!) directive
	authDirective := directives.GetAuth()
	resolverConfig.Directives.Auth = authDirective.AuthDirective

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(*resolverConfig))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
