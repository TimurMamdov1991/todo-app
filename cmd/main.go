package main

import (
	"github.com/TimurMamdov1991/todo-app"
	"github.com/TimurMamdov1991/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8888", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
