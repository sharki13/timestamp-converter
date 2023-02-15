package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.NewWithID("com.github.sharki13.timestamp-converter")

	mainWindow := app.NewWindow("Timestamp converter")
	timestampConverter := &TimestampConverter{}

	timestampConverter.SetupAndRun(mainWindow, app)
}
