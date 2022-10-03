package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildingTodoShouldSetDoneAsFalse(t *testing.T) {
	description := "Clean the Void on Your Soul"

	todo := NewTodo(description)

	assert.Equal(t, todo.Description, description)
	assert.Equal(t, todo.Done, false)
}

func TestTodoMarkAsDoneAndAsToDoWorksCorrectly(t *testing.T) {
	description := "Some Todo"
	todo := NewTodo(description)

	assert.Equal(t, todo.Description, description)
	assert.Equal(t, false, todo.Done)

	todo.MarkAsDone()
	assert.Equal(t, true, todo.Done)

	todo.MarkAsToDo()
	assert.Equal(t, false, todo.Done)
}

func TestTodoToMapReturnsCorrectUpdatableFields(t *testing.T) {
	description := "A todo"
	todo := NewTodo(description)
	mapped := todo.ToMap()

	desc, hasDesc := mapped["description"]
	assert.True(t, hasDesc)
	assert.Equal(t, "A todo", desc)
	done, hasDone := mapped["done"]
	assert.True(t, hasDone)
	assert.Equal(t, done, false)
}
