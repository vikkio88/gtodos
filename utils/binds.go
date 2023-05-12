package utils

import (
	"gtodos/models"

	"fyne.io/fyne/v2/data/binding"
)

func TodoFromDataItem(item binding.DataItem) *models.Todo {
	v, _ := item.(binding.Untyped).Get()
	return v.(*models.Todo)
}
