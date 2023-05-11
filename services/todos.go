package services

import (
	"errors"
	"gtodos/db"
	"gtodos/models"

	"fyne.io/fyne/v2/data/binding"
)

type Todos struct {
	db    *db.Db
	todos map[string]*models.Todo
	Data  binding.StringList
}

func NewTodosFromDb(db *db.Db) Todos {
	todoList := db.GetAllTodos()
	return NewTodos(db, todoList)
}

func NewTodos(db *db.Db, todos []models.Todo) Todos {
	t := Todos{
		db:    db,
		todos: map[string]*models.Todo{},
		Data:  binding.NewStringList(),
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
	t.todos = map[string]*models.Todo{}
	t.Data.Set([]string{})
}

func (t *Todos) Persist() {
	t.db.Save(t.All())
}

func (t *Todos) Get(id string) (*models.Todo, error) {
	if t, ok := t.todos[id]; ok {
		return t, nil
	}

	return nil, errors.New("No Todo")
}

func (t *Todos) All() []*models.Todo {
	result := []*models.Todo{}
	for _, v := range t.todos {
		result = append(result, v)
	}

	return result
}

func (t *Todos) Add(todo *models.Todo) {
	if todo.Id == "" {
		t.db.InsertTodo(todo)
	}
	t.todos[todo.Id] = todo
	t.Data.Append(todo.Id)
}
