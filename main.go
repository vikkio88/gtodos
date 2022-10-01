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

	db.MakeDb("ciao")

	fmt.Println(t)
}
