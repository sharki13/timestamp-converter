package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	app := app.NewWithID("com.github.sharki13.timestamp-converter")

	mainWindow := app.NewWindow("Timestamp converter")
	timestampConverter := &TimestampConverter{}

	timestampConverter.SetupAndRun(mainWindow, app)
}
