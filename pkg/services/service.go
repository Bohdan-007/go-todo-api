package service

import repository "github.com/Bohdan-007/go-todo-api/pkg/repositories"

type Authorization interface {
}

type TdoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TdoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
