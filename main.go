package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/sharki13/timestamp-converter/gui"
)

func main() {
	app := app.NewWithID("com.github.sharki13.timestamp-converter")
	tc := gui.NewTimestampConverter(app)

	tc.ShowAndRun()
}
