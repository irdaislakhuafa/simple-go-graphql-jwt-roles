// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthOptions struct {
	Register *ResponseToken `json:"register"`
	Login    *ResponseToken `json:"login"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewRole struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NewUser struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

type ResponseToken struct {
	Token string `json:"token"`
}

type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RoleMutationOptions struct {
	Save *Role `json:"save"`
}

type RoleQueryOptions struct {
	GetAll []*Role `json:"getAll"`
}

type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

type UserOptions struct {
	GetAll  []*User `json:"getAll"`
	GetByID *User   `json:"getById"`
}
