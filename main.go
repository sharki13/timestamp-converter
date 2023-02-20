package main

import (
	"com.sharki13/timestamp.converter/timestamp_converter"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.NewWithID("com.github.sharki13.timestamp-converter")
	tc := timestamp_converter.NewTimestampConverter(app)

	tc.ShowAndRun()
}
