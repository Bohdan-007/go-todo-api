package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type TdoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TdoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
