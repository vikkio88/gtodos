package services

import (
	"gtodos/models"

	"fyne.io/fyne/v2/data/binding"
)

type Todos struct {
	Data binding.UntypedList
}

func NewTodos() Todos {
	return Todos{
		Data: binding.NewUntypedList(),
	}
}

func (t *Todos) Drop() {
	list, _ := t.Data.Get()
	list = list[:0]
	t.Data.Set(list)
}

func (t *Todos) All() []*models.Todo {
	result := []*models.Todo{}
	for i := 0; i < t.Data.Length(); i++ {
		di, err := t.Data.GetItem(i)
		if err != nil {
			break
		}
		result = append(result, models.NewTodoFromDataItem(di))
	}

	return result
}

func (t *Todos) Add(todo *models.Todo) {
	t.Data.Prepend(todo)
}
