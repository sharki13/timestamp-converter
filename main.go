package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()

	mainWindow := app.NewWindow("Timestamp converter")
	timestampConverter := &TimestampConverter{}

	timestampConverter.SetupAndRun(mainWindow, app)
}
