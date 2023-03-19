package main

import (
	"github.com/Nurdildag/todo-app"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server", err.Error())
	}

}
