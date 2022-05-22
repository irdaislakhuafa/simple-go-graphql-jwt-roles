# About

This project is for learning for me about how to implement `GraphQL` with `JWT (JSON Web Token)` authorization in `GO` programming language. Here i use a specific roles for some endpoint that may contains sensitive data. In this project, roles management just for simulation and learning but can also applied in the real world.

## Dependencies

    - github.com/99designs/gqlgen => for generate GraphQL resolver with Schema First concept
    - github.com/joho/godotenv => to load enviroment variables with `.env` file for application config
    - github.com/google/uuid => to generate ID with String `uuid`
    - github.com/gorilla/mux => for routing
    - golang.org/x/crypto/bcrypt => to hashing password
    - github.com/dgrijalva/jwt-go => library to generate `JWT Token`, validate token and more. 

## How to run

To make it easier to run this appliaction, i use docker for container. So make sure you have docker installed or if you haven't installed docker, you can download docker and install it [here](https://docs.docker.com/get-docker/) first.

Run this command and docker will prepare the enviroment to run this app

```bash
cd env/
docker compose -f docker-compose.yml up -d
```

## How to use

To use this application you can access with url `http://localhost:8080` and write some `GrapQL Query` or `Mutation`. You can see simple documentation in the upper right corner.

Below are some `Query`/`Mutation` that i want to show you

- ##  Roles

    ``` graphql
    # to create new role (currently without @auth)
    mutation NewRole {
      role {
        save(newRole: {name: "admin", description: "-"}) {
          id
          name
          description
        }
      }
    }
    
    # to get all roles (just roles "ADMIN" and "USER" can access it)
    query GetAllRoles {
      role {
        getAll {
          id
          name
          description
        }
      }
    }
    ```

- ## User

    ```graphql
    # to get all user (just roles "ADMIN" can access it)
    query GetAllUser {
      user {
        getAll {
          id
          name
          email
          password
          roles
        }
      }
    }

    # to get user by id (just roles "USER" and "ADMIN" can access it )
    query GetUserByID {
      user {
        getById(userId: "439eb4e8-e7f3-40af-adeb-825e836b2da4") {
          id
          name
          email
          password
        }
      }
    }


    ```

- ## Auth

    ```graphql
    # to login a user, this endpoint is without @auth
    mutation Login {
      auth {
        login(
          user: {
            email:"admin@gmail.com"
            password: "admin"
          }

        ) {
          token
        }
      }
    }

    # to register new user, this endpoint is without @auth for authentication
    mutation Register {
      auth {
        register (
          newUser: {
            name: "Me is ADMIN"
            email:"admin@gmail.com"
            password:"admin"
            roles: [
              "user", "admin"
            ]
          }
        ) {
          token
        }
      }
    }
    ```

    For mor details you can see `Docs`.


- TODO: Design Flowchart
- TODO: Description for service
