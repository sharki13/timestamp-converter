package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.NewWithID("com.github.sharki13.timestamp-converter")
	timestampConverter := NewTimestampConverter(app)

	timestampConverter.SetupAndRun()
}
