package service

import (
	todo "github.com/Njrctr/restapi-todo"
	"github.com/Njrctr/restapi-todo/pkg/repository"
)

type Autorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
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
