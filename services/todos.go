package services

import (
	"gtodos/db"
	"gtodos/models"

	"fyne.io/fyne/v2/data/binding"
)

type Todos struct {
	binding.UntypedList
	db *db.Db
}

func NewTodosFromDb(db *db.Db) Todos {
	todoList := db.GetAllTodos()
	return NewTodos(db, todoList)
}

func NewTodos(db *db.Db, todos []models.Todo) Todos {
	t := Todos{
		binding.NewUntypedList(),
		db,
	}

	for _, td := range todos {
		// I dont know but this is needed
		// may be this https://levelup.gitconnected.com/go-for-range-slice-bug-lessons-learned-fa401d5d8c9a
		td1 := td
		t.Add(&td1)
	}

	return t
}

func (t *Todos) Drop() {
	t.db.Drop()
	list, _ := t.Get()
	list = list[:0]
	t.Set(list)
}

func (t *Todos) Persist() {
	t.db.Save(t.All())
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
	if todo.Id == "" {
		t.db.InsertTodo(todo)
	}
	t.Prepend(todo)
}
