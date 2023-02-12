package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"golang.design/x/clipboard"
)

func log(text string) {
	now := time.Now()
	timestamp := now.UTC().Format(time.RFC3339)

	fmt.Printf("[%s] %s\n", timestamp, text)
}

func main() {

	app := app.New()

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	mainWindow := app.NewWindow("Timestamp converter")
	mainWindow.Resize(fyne.NewSize(600, 100))

	timeConverter := TimeConverter{}
	timeConverter.Make()

	content := container.New(layout.NewVBoxLayout(), container.New(layout.NewHBoxLayout(), timeConverter.ReturnButtons()...), container.New(layout.NewGridLayout(3), timeConverter.ReturnTimestampSets()...), timeConverter.ReturnStatus())

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}
