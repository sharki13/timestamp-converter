package main

import (
	"time"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"golang.design/x/clipboard"
)

// define enum for the different sets types, rfc and unix
type TimestampItemsSetType int

const (
	RFC3339 TimestampItemsSetType = iota
	RFC3339Nano
	Unix
)

type TimestampItemsSet struct {
	entry *widget.Entry
	label *widget.Label
	copyButton *widget.Button
	setType TimestampItemsSetType
	loc *time.Location
}

func MakeTimestampItemsSet(labelText string, setType TimestampItemsSetType, timezone *time.Location) TimestampItemsSet {
	if timezone == nil {
		panic("timezone cannot be nil")
	}

	t := TimestampItemsSet{}

	t.entry = widget.NewEntry()
	t.label = widget.NewLabel(labelText)
	t.copyButton = widget.NewButton("Copy", func() {
		clipboard.Write(clipboard.FmtText, []byte(t.entry.Text))
	})

	style := fyne.TextStyle{
		Bold: true,
	}

	t.label.TextStyle = style
	t.entry.TextStyle = style

	t.setType = setType
	t.loc = timezone

	return t
}

func (t *TimestampItemsSet) Update(e time.Time) {
	switch t.setType {
	case RFC3339:
		t.entry.SetText(e.In(t.loc).Format(time.RFC3339))
	case RFC3339Nano:
		t.entry.SetText(e.In(t.loc).Format(time.RFC3339Nano))
	case Unix:
		t.entry.SetText(strconv.FormatInt(e.Unix(), 10))
	}
}

func (t *TimestampItemsSet) ReturnItems() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		t.label,
		t.entry,
		t.copyButton,
	}
}
