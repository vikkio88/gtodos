package main

import (
	"gtodos/db"
	"gtodos/models"
	"gtodos/services"
	"gtodos/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	db := db.MakeDb("db_files")
	defer db.Close()

	todos := services.NewTodosFromDb(&db)
	a := app.NewWithID("gtodos")
	w := a.NewWindow("GTodos")

	w.Resize(fyne.NewSize(480, 600))

	list := widget.NewListWithData(
		todos.Data,
		rederListLine,
		bindDataToListLine(&todos),
	)

	input := widget.NewEntry()
	addButton := widget.NewButton("Add", func() {
		t := models.NewTodo(input.Text)
		todos.Add(&t)
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
	resetButton := widget.NewButton("Delete All", func() {
		todos.Drop()
	})
	saveButton := widget.NewButton("Save", func() {
		todos.Persist()
	})
	bottom := container.New(layout.NewVBoxLayout(), container.New(layout.NewAdaptiveGridLayout(2), input, addButton), resetButton, saveButton)
	content := container.New(layout.NewBorderLayout(nil, bottom, nil, nil), bottom, list)

	w.SetContent(content)
	w.ShowAndRun()
}

func bindDataToListLine(todos *services.Todos) func(di binding.DataItem, co fyne.CanvasObject) {
	return func(di binding.DataItem, co fyne.CanvasObject) {
		t := utils.TodoFromDataItem(di)
		container := co.(*fyne.Container)
		label := container.Objects[1].(*widget.Label)
		check := container.Objects[0].(*widget.Check)
		label.Bind(binding.BindString(&t.Description))
		check.Bind(binding.BindBool(&t.Done))
	}
}
func rederListLine() fyne.CanvasObject {
	c := widget.NewCheck("", nil)
	return container.New(layout.NewBorderLayout(nil, nil, nil, c), c, widget.NewLabel(""))
}
