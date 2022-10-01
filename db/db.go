package db

import (
	c "github.com/ostafen/clover"
)

type Db struct {
	db *c.DB
}

func MakeDb(fileName string) Db {
	db, _ := c.Open(fileName)
	return Db{db}
}
