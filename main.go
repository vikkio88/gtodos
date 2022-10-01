package main

import (
	"gtodos/db"
	"gtodos/models"
)

func main() {
	t := models.Todo{
		Description: "some",
		Done:        false,
	}

	data := db.MakeDb("db_files")
	data.GetAllTodos()
	for range [1000]int{} {
		data.InsertTodo(t)
	}
	// ts := data.GetAllTodos()
	// for _, t := range ts {
	// 	fmt.Println(t)
	// }
}
