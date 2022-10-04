package main

import (
	// "gtodos/db"

	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// db := db.MakeDb("db_files")
	// todos := db.GetAllTodos()
	a := app.NewWithID("gtodos")
	w := a.NewWindow("GTodos")

	w.Resize(fyne.NewSize(480, 320))
	var data = []string{"a", "string", "list"}

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	topContent := container.New(layout.NewVBoxLayout(), list)

	top := container.New(layout.NewMaxLayout(), topContent)
	input := widget.NewEntry()
	button := widget.NewButton("Add", func() { fmt.Println("clicked") })
	bottom := container.New(layout.NewMaxLayout(), container.New(layout.NewVBoxLayout(), input, button))

	content := container.New(layout.NewVBoxLayout(), top, layout.NewSpacer(), container.New(layout.NewMaxLayout(), bottom))
	w.SetContent(content)

	w.ShowAndRun()
}
