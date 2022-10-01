package main

import (
	"fmt"
	"gtodos/models"
)

func main() {
	t := models.Todo{
		Description: "some",
		Done:        false,
	}

	fmt.Println(t)
}
