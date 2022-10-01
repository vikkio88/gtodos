package models

type Todo struct {
	Description string
	Done        bool
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
