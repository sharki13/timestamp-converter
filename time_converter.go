package main

import (
	"time"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

)

type TimeConverter struct {
	items []TimestampItemsSet
	status *widget.Label
}

func (i *TimeConverter) Make() {
	i.items = append(i.items, MakeTimestampItemsSet("Unix", Unix, time.UTC))
	i.items = append(i.items, MakeTimestampItemsSet("UTC", RFC3339, time.UTC))
	i.items = append(i.items, MakeTimestampItemsSet("Local", RFC3339, time.Local))
	i.items = append(i.items, MakeTimestampItemsSet("EST", RFC3339, time.FixedZone("EST", -5*60*60)))
	i.items = append(i.items, MakeTimestampItemsSet("PST", RFC3339, time.FixedZone("PST", -8*60*60)))
	i.Update(time.Now())
	i.status = widget.NewLabel(fmt.Sprintf("[%s] Ready", time.Now().Format("15:04:05")))
}

func (i *TimeConverter) SetStatus(text string) {
	if i.status != nil {
		i.status.SetText(fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), text))
	}
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