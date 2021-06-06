package application

import (
	"context"
	"github.com/ereb-or-od/kenobi-http-server/walking-skeleton/todo-service/pkg/domain"
	mediator "github.com/ereb-or-od/kenobi-mediator"
)

type CreateTodoCommand struct {
	Name string
}

type CreateTodoContract struct {
	Id string
}

func (*CreateTodoCommand) Key() string { return "CreateTodoCommand" }

type CreateTodoCommandHandler struct {
	baseHandler *BaseHandler
}

func NewCreateTodoCommandHandler(baseHandler *BaseHandler) CreateTodoCommandHandler {
	return CreateTodoCommandHandler{baseHandler: baseHandler}
}

func (c CreateTodoCommandHandler) Handle(_ context.Context, command mediator.Message) (interface{}, error) {
	cmd := command.(*CreateTodoCommand)
	todo := domain.New(cmd.Name)
	c.baseHandler.repository.Create(todo)
	return &CreateTodoContract{
		Id: todo.Id,
	}, nil
}
