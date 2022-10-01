package db

import (
	"fmt"
	"gtodos/models"
	"time"

	c "github.com/ostafen/clover/v2"
)

const TODO_COLLECTION = "TODOS"

type Db struct {
	db *c.DB
}

func (db *Db) GetAllTodos() []models.Todo {
	docs, err := db.db.FindAll(c.NewQuery(TODO_COLLECTION))
	if err != nil {
		panic(err)
	}
	result := make([]models.Todo, 0)
	for _, doc := range docs {
		t := models.Todo{}
		doc.Unmarshal(&t)
		t.Id = doc.ObjectId()
		result = append(result, t)
	}

	return result
}

func (db *Db) InsertTodo(todo models.Todo) {
	doc := c.NewDocumentOf(todo)
	x := time.Now()
	_, err := db.db.InsertOne(TODO_COLLECTION, doc)
	fmt.Println(time.Since(x))
	if err != nil {
		panic(err)
	}

}

func MakeDb(fileName string) Db {
	db, _ := c.Open(fileName)
	db.CreateCollection(TODO_COLLECTION)
	return Db{db}
}
