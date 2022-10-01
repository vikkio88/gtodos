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
