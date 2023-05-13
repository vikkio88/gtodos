package services

import (
	"gtodos/models"

	"fyne.io/fyne/v2/data/binding"
)

type Todos struct {
	binding.UntypedList
}

func NewTodos() Todos {
	return Todos{
		binding.NewUntypedList(),
	}
}

func (t *Todos) Drop() {
	list, _ := t.Get()
	list = list[:0]
	t.Set(list)
}

func (t *Todos) All() []*models.Todo {
	result := []*models.Todo{}
	for i := 0; i < t.Length(); i++ {
		di, err := t.GetItem(i)
		if err != nil {
			break
		}
		result = append(result, models.NewTodoFromDataItem(di))
	}

	return result
}

func (t *Todos) Add(todo *models.Todo) {
	t.Prepend(todo)
}
