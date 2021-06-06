package repository

import (
	"github.com/ereb-or-od/kenobi-http-server/walking-skeleton/todo-service/pkg/domain"
	"github.com/ereb-or-od/kenobi-http-server/walking-skeleton/todo-service/pkg/domain/repository/interfaces"
)

type todoRepository struct {
	database map[string]*domain.Todo
}

func (t todoRepository) Delete(id string) {
	t.database[id]  = nil
}

func (t todoRepository) FindById(id string) *domain.Todo {
	return t.database[id]
}

func (t todoRepository) Create(todo *domain.Todo) {
	t.database[todo.Id] = todo
}

func NewTodoRepository() interfaces.TodoRepository {
	return &todoRepository{
		database: map[string]*domain.Todo{},
	}
}
