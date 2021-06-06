package main

import (
	kenobihttpserver "github.com/ereb-or-od/kenobi-http-server"
	"github.com/ereb-or-od/kenobi-http-server/walking-skeleton/todo-service/pkg/api"
)

func main() {
	kenobiServer := kenobihttpserver.New("todo_app").
		WithDefaultLogger().
		UseHttp().
		WithLoggingMiddleware().
		WithRecoverMiddleware().
		WithRequestIDMiddleware().
		WithAllowAnyCORSMiddleware().
		WithGzipMiddleware().
		WithHealthCheckMiddleware("/ping", "pong!").
		WithController(api.NewTodoController())
	kenobiServer.Start()

}
