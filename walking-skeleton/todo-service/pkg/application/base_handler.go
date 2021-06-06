package application

import (
	"github.com/ereb-or-od/kenobi-http-server/walking-skeleton/todo-service/pkg/domain/repository/interfaces"
	"github.com/ereb-or-od/kenobi-http-server/walking-skeleton/todo-service/pkg/infrastructure/repository"
)

type BaseHandler struct {
	repository interfaces.TodoRepository
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{
		repository: repository.NewTodoRepository(),
	}
}
