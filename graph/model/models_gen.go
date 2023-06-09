// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Name   string `json:"name"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	User *User  `json:"user"`
}

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
