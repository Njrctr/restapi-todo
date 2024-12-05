package service

import (
	todo "github.com/Njrctr/restapi-todo"
	"github.com/Njrctr/restapi-todo/pkg/repository"
)

type Autorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateJWTToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
}

type TodoItem interface {
}

type Service struct {
	Autorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
	}
}
