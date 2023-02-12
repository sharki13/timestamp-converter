package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
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


	nowButton := widget.NewButton("Now", func() {
		timeConverter.Update(time.Now())
	})

	fromCliboardButton := widget.NewButton("From clipboard", func() {
		clipboardContent := string(clipboard.Read(clipboard.FmtText))

		t, err := PraseStringToTime(clipboardContent)
		if err == nil {
			timeConverter.Update(t)
			return
		}

		timeConverter.SetStatus("Clipboard content is not a valid timestamp")
	})

	content := container.New(layout.NewVBoxLayout(), nowButton, fromCliboardButton, container.New(layout.NewGridLayout(3), timeConverter.ReturnTimestampSets()...), timeConverter.ReturnStatus())

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}
