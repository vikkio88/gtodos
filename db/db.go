package db

import (
	"gtodos/models"

	c "github.com/ostafen/clover/v2"
)

const TODO_COLLECTION = "TODOS"

type Db struct {
	db *c.DB
}

func (db *Db) Close() error {
	return db.db.Close()
}

func formatTodo(doc *c.Document) models.Todo {
	t := models.Todo{}
	doc.Unmarshal(&t)
	t.Id = doc.ObjectId()
	return t
}

func (db *Db) GetAllTodos() []models.Todo {
	docs, err := db.db.FindAll(c.NewQuery(TODO_COLLECTION))
	if err != nil {
		panic(err)
	}
	result := make([]models.Todo, 0)
	for _, doc := range docs {
		result = append(result, formatTodo(doc))
	}

	return result
}

func (db *Db) FindTodoById(id string) *models.Todo {
	doc, err := db.db.FindById(TODO_COLLECTION, id)
	if err != nil {
		return nil
	}

	t := formatTodo(doc)

	return &t

}

func (db *Db) InsertTodo(todo *models.Todo) bool {
	doc := c.NewDocumentOf(todo.ToMap())
	id, err := db.db.InsertOne(TODO_COLLECTION, doc)
	todo.Id = id
	return err == nil
}

func (db *Db) Save(todos []*models.Todo) bool {
	for _, t := range todos {
		if t.Id == "" {
			db.InsertTodo(t)
		} else {
			db.UpdateTodo(t)
		}
	}
	return true
}

func (db *Db) UpdateTodo(todo *models.Todo) bool {
	mapped := todo.ToMap()
	err := db.db.UpdateById(TODO_COLLECTION, todo.Id, mapped)

	return err == nil
}

func (db *Db) Drop() bool {
	err := db.db.Delete(c.NewQuery(TODO_COLLECTION))

	return err == nil
}

func (db *Db) DeleteTodo(todo *models.Todo) bool {
	err := db.db.DeleteById(TODO_COLLECTION, todo.Id)

	return err == nil
}

func MakeDb(fileName string) Db {
	db, _ := c.Open(fileName)
	db.CreateCollection(TODO_COLLECTION)
	return Db{db}
}
