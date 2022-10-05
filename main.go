package main

import (
	// "gtodos/db"

	"gtodos/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// db := db.MakeDb("db_files")
	// todos := db.GetAllTodos()
	a := app.NewWithID("gtodos")
	w := a.NewWindow("GTodos")

	w.Resize(fyne.NewSize(480, 320))
	var data = binding.NewUntypedList()

	list := widget.NewListWithData(
		data,
		func() fyne.CanvasObject {
			return widget.NewLabel("placeholder")
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			di.(*models.Todo).Description
			co.(*widget.Label).Bind(binding.BindString())
		},
	)

	input := widget.NewEntry()
	button := widget.NewButton("Add", func() {
		data = append(data, input.Text)
	})
	bottom := container.New(layout.NewVBoxLayout(), input, button)

	content := container.New(layout.NewBorderLayout(nil, bottom, nil, nil), bottom, list)
	w.SetContent(content)

	w.ShowAndRun()
}
