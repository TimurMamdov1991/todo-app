package main

import (
	"github.com/TimurMamdov1991/todo-app"
	"github.com/TimurMamdov1991/todo-app/pkg/handler"
	"github.com/TimurMamdov1991/todo-app/pkg/repository"
	service "github.com/TimurMamdov1991/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8888", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
