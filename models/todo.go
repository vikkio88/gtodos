package models

import (
	"fmt"
)

type Todo struct {
	Id          string
	Description string `clover:"description"`
	Done        bool   `clover:"done"`
}

func (t Todo) String() string {
	return fmt.Sprintf("%s - %s  - %t", t.Id, t.Description, t.Done)
}

func (t *Todo) MarkAsDone() {
	t.Done = true
}

func (t *Todo) MarkAsToDo() {
	t.Done = false
}

func (t *Todo) ToMap() map[string]interface{} {
	result := map[string]interface{}{}

	result["description"] = t.Description
	result["done"] = t.Done

	return result
}

func NewTodo(description string) Todo {
	return Todo{Description: description, Done: false}
}
