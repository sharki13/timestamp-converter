package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.design/x/clipboard"
)

const statusTimeFormat = "15:04:05"

type TimeConverter struct {
	items []TimestampItemsSet
	status *widget.Label
	nowButton *widget.Button
	fromCliboardButton *widget.Button
}

func (i *TimeConverter) Make() {
	// get local timezone shift
	localTzName, offset := time.Now().Zone()
	

	i.items = append(i.items, MakeTimestampItemsSet("Unix", Unix, time.UTC, i.Update, i.SetStatus))
	i.items = append(i.items, MakeTimestampItemsSet("UTC", RFC3339, time.UTC, i.Update, i.SetStatus))
	i.items = append(i.items, MakeTimestampItemsSet(fmt.Sprintf("Local: %s (%d:00)", localTzName, offset/3600), RFC3339, time.Local, i.Update, i.SetStatus))
	i.items = append(i.items, MakeTimestampItemsSet("EST (-5:00)", RFC3339, time.FixedZone("EST", -5*60*60), i.Update, i.SetStatus))
	i.items = append(i.items, MakeTimestampItemsSet("PST (-8:00)", RFC3339, time.FixedZone("PST", -8*60*60), i.Update, i.SetStatus))
	i.status = widget.NewLabel("")
	i.Update(time.Now())
	i.SetStatus("Ready")
	i.nowButton = widget.NewButton("Now !", func() {
		i.Update(time.Now())
		i.SetStatus("Updated to now")
	})
	i.fromCliboardButton = widget.NewButtonWithIcon("", theme.ContentPasteIcon(), func() {
		clipboardContent := string(clipboard.Read(clipboard.FmtText))

		t, err := PraseStringToTime(clipboardContent)
		if err == nil {
			i.Update(t)
			i.SetStatus("Updated from clipboard")
			return
		}
		i.SetStatus("Clipboard content is not a valid timestamp")
	})
}

func (i *TimeConverter) SetStatus(text string) {
	i.status.SetText(fmt.Sprintf("[%s] %s", time.Now().Format(statusTimeFormat), text))
}

// function which takes a time.Time and updates the inputs
func (i *TimeConverter) Update(t time.Time) {
	for _, item := range i.items {
		item.Update(t)
	}

	i.SetStatus("Updated")
}

// return all inputs as a slice of fyne widgets
func (i *TimeConverter) ReturnTimestampSets() []fyne.CanvasObject {
	var items []fyne.CanvasObject

	for _, item := range i.items {
		items = append(items, item.ReturnItems()...)
	}

	return items
}

func (i *TimeConverter) ReturnStatus() fyne.CanvasObject {
	return i.status
}

func (i *TimeConverter) ReturnButtons() []fyne.CanvasObject {
	return []fyne.CanvasObject{i.nowButton, i.fromCliboardButton}
}