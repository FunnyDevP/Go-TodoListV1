package handler

import (
	"github.com/FunnyDevP/Go-TodoListV1/internal/repository"
	"net/http"
)

type TodoList interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

type todoList struct {
	repo repository.TodoList
}

func NewTodoList(repo repository.TodoList) *todoList {
	return &todoList{repo: repo}
}

func (t todoList) GetList(w http.ResponseWriter, r *http.Request) {

}
