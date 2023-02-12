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

	mainWindow.Resize(fyne.NewSize(650, 100))

	timeConverter := TimeConverter{}
	timeConverter.Make()

	content := container.New(layout.NewVBoxLayout(), container.New(layout.NewHBoxLayout(), timeConverter.ReturnButtons()...), container.New(layout.NewGridLayout(3), timeConverter.ReturnTimestampSets()...), timeConverter.ReturnStatus())

	// run funtion in backgound to check clipboard
	go func(enable *bool) {
		for {
			if *enable {
				clipboardContent := string(clipboard.Read(clipboard.FmtText))
				clipTime, err := PraseStringToTime(clipboardContent)
				if err == nil && clipTime != timeConverter.CurrentTimestamp {
					timeConverter.Update(clipTime)
					timeConverter.SetStatus("Updated from clipboard")
				}
				time.Sleep(1 * time.Second)
			}
		}
	}(&timeConverter.watchClipboard)

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}
