package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.design/x/clipboard"
)

type TimezoneDefinition struct {
	Name string
	Location string
}

func getLocalTimezoneName() string {
	name, _ := time.Now().Zone()

	return name
}

var Timezones = []TimezoneDefinition{
	{"Unix", "Unix"},
	{"UTC", "UTC"},
	{getLocalTimezoneName(), "Local"},
	{"HST", "Pacific/Honolulu"},
	{"PST/PDT", "America/Los_Angeles"},
	{"CST/CDT", "America/Chicago"},
	{"EST/EDT", "America/New_York"},
	{"CET/CEST", "Europe/Paris"},
}

const statusTimeFormat = "15:04:05"

type TimeConverter struct {
	items []TimestampItemsSet
	status *widget.Label
	nowButton *widget.Button
	fromCliboardButton *widget.Button
	CurrentTimestamp time.Time
	watchClipboard bool
	watchClipboardCheck *widget.Check
}

func (i *TimeConverter) Make() {
	// get local timezone shift
	
	
	for _, tz := range Timezones {
		if tz.Location == "Unix" {
			i.items = append(i.items, MakeTimestampItemsSet(tz.Name, Unix, time.UTC, i.Update, i.SetStatus))
		} else {
			loc, err := time.LoadLocation(tz.Location)
			if err != nil {
				panic(err)
			}
			i.items = append(i.items, MakeTimestampItemsSet(tz.Name, RFC3339, loc, i.Update, i.SetStatus))
		}
	}

	i.status = widget.NewLabel("")
	i.Update(time.Now())
	i.CurrentTimestamp = time.Now()
	i.SetStatus("Ready")
	i.nowButton = widget.NewButton("Now !", func() {
		i.Update(time.Now())
		i.SetStatus("Updated to now")
		i.watchClipboard = false
		i.watchClipboardCheck.SetChecked(false)
	})

	i.nowButton.Importance = widget.WarningImportance

	i.watchClipboardCheck = widget.NewCheck("Watch clipboard", func(b bool) {
		i.watchClipboard = b
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
	i.CurrentTimestamp = t

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
	return []fyne.CanvasObject{i.nowButton, i.fromCliboardButton, i.watchClipboardCheck}
}