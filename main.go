package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()

	mainWindow := app.NewWindow("Timestamp converter")
	mainWindow.Resize(fyne.NewSize(600, 10))
	timestampConverter := &TimestampConverter{}

	timestampConverter.SetContent(mainWindow, app)

	mainWindow.ShowAndRun()

}
