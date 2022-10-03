package main

import (
	// "gtodos/db"

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

	title := widget.NewLabel("GTodos")
	title1 := widget.NewLabel("GTodos1")
	title2 := widget.NewLabel("GTodos2")

	grid := container.New(layout.NewGridLayout(2), title, title1, title2)
	w.SetContent(grid)

	w.ShowAndRun()
}
