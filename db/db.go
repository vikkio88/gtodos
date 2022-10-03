package db

import (
	"fmt"
	"gtodos/models"

	c "github.com/ostafen/clover/v2"
)

const TODO_COLLECTION = "TODOS"

type Db struct {
	db *c.DB
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
	doc := c.NewDocumentOf(todo)
	id, err := db.db.InsertOne(TODO_COLLECTION, doc)
	todo.Id = id
	return err == nil
}

func (db *Db) InsertManyTodos(todos []models.Todo) bool {
	docs := make([]*c.Document, len(todos))
	for i, t := range todos {
		docs[i] = c.NewDocumentOf(t)
	}
	err := db.db.Insert(TODO_COLLECTION, docs...)
	return err == nil
}

func (db *Db) UpdateTodo(todo *models.Todo) bool {
	mapped := todo.ToMap()
	fmt.Println(mapped)
	err := db.db.UpdateById(TODO_COLLECTION, todo.Id, mapped)

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
