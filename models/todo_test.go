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
	assert.Equal(t, todo.Done, false)

	todo.MarkAsDone()
	assert.Equal(t, todo.Done, true)

	todo.MarkAsToDo()
	assert.Equal(t, todo.Done, false)
}
