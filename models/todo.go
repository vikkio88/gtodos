package models

type Todo struct {
	Description string `clover:"description"`
	Done        bool   `clover:"done"`
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
