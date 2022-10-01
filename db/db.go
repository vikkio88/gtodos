package db

import (
	"gtodos/models"

	c "github.com/ostafen/clover/v2"
)

const TODO_COLLECTION = "TODOS"

type Db struct {
	db *c.DB
}

func (db *Db) GetAllTodos() []models.Todo {
	docs, _ := db.db.FindAll(c.NewQuery(TODO_COLLECTION))
	result := make([]models.Todo, 0)
	for _, doc := range docs {
		t := models.Todo{}
		doc.Unmarshal(&t)
		result = append(result, t)
	}

	return result
}

func (db *Db) InsertTodo(todo models.Todo) {
	doc := c.NewDocumentOf(todo)
	db.db.InsertOne(TODO_COLLECTION, doc)
}

func MakeDb(fileName string) Db {
	db, _ := c.Open(fileName)
	db.CreateCollection(TODO_COLLECTION)
	return Db{db}
}
