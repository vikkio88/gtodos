package main

import (
	"fmt"
	"gtodos/db"
	"gtodos/models"
)

func main() {
	t := models.Todo{
		Description: "some",
		Done:        false,
	}

	data := db.MakeDb("ciao")
	data.GetAllTodos()
	data.InsertTodo(t)
	ts := data.GetAllTodos()
	for _, t := range ts {
		fmt.Println(t)
	}
}
