package repository

import (
	"database/sql"
	"github.com/FunnyDevP/Go-TodoListV1/internal/model"
)

type TodoList interface {
	GetList() ([]*model.Tasks, error)
}

type todoList struct {
	db *sql.DB
}

func NewTodoList(repo *sql.DB) *todoList {
	return &todoList{db: repo}
}

func (t todoList) GetList() ([]*model.Tasks, error) {
	return nil, nil
}
