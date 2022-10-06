package main

import (
	// "gtodos/db"

	"fmt"
	"gtodos/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func todoFromDataItem(item binding.DataItem) *models.Todo {
	u, _ := item.(binding.Untyped)
	v, _ := u.Get()
	t, _ := v.(models.Todo)
	return &t
}

func main() {
	// db := db.MakeDb("db_files")
	// todos := db.GetAllTodos()
	a := app.NewWithID("gtodos")
	w := a.NewWindow("GTodos")

	w.Resize(fyne.NewSize(480, 600))
	var data = binding.NewUntypedList()
	data.Append(models.NewTodo("test"))

	list := widget.NewListWithData(
		data,
		func() fyne.CanvasObject {
			c := widget.NewCheck("", nil)
			return container.New(layout.NewBorderLayout(nil, nil, nil, c), c, widget.NewLabel(""))
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			t := todoFromDataItem(di)
			container := co.(*fyne.Container)
			label := container.Objects[1].(*widget.Label)
			check := container.Objects[0].(*widget.Check)
			label.Bind(binding.BindString(&t.Description))
			check.Bind(binding.BindBool(&t.Done))
		},
	)

	input := widget.NewEntry()
	addButton := widget.NewButton("Add", func() {
		data.Append(models.NewTodo(input.Text))
		input.SetText("")
	})
	addButton.Disable()
	input.OnChanged = func(s string) {
		if len(s) > 2 {
			addButton.Enable()
			return
		}

		addButton.Disable()
	}

	resetButton := widget.NewButton("Reset", func() {
		list, _ := data.Get()
		list = list[:0]
		data.Set(list)
	})
	saveButton := widget.NewButton("Save", func() {
		fmt.Println("saving")
	})

	bottom := container.New(layout.NewVBoxLayout(), container.New(layout.NewAdaptiveGridLayout(2), input, addButton), resetButton, saveButton)

	content := container.New(layout.NewBorderLayout(nil, bottom, nil, nil), bottom, list)
	w.SetContent(content)

	w.ShowAndRun()
}
