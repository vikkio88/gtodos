package models

type Todo struct {
	Description string
	Done        bool
}

func NewTodo(description string) Todo {
	return Todo{Description: description, Done: false}
}
