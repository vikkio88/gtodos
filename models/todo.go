package models

import (
	"fmt"

	"fyne.io/fyne/v2/data/binding"
)

type Todo struct {
	Id          string
	Description string `clover:"description"`
	Done        bool   `clover:"done"`
}

func NewTodoFromDataItem(item binding.DataItem) *Todo {
	v, _ := item.(binding.Untyped).Get()
	return v.(*Todo)
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

func NewTodo(description string) Todo {
	return Todo{Description: description, Done: false}
}
