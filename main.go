package main

import (
	"fmt"
	"strconv"
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

type Inputs struct {
	items []TimestampItemsSet
}

func (i *Inputs) MakeInputs() {

	i.items = append(i.items, MakeTimestampItemsSet("Unix", Unix, time.UTC))
	i.items = append(i.items, MakeTimestampItemsSet("UTC", RFC3339, time.UTC))
	i.items = append(i.items, MakeTimestampItemsSet("Local", RFC3339, time.Local))
	i.items = append(i.items, MakeTimestampItemsSet("EST", RFC3339, time.FixedZone("EST", -5*60*60)))
	i.items = append(i.items, MakeTimestampItemsSet("PST", RFC3339, time.FixedZone("PST", -8*60*60)))
}

// function which takes a time.Time and updates the inputs
func (i *Inputs) UpdateInputs(t time.Time) {
	for _, item := range i.items {
		item.Update(t)
	}
}

// return all inputs as a slice of fyne widgets
func (i *Inputs) ReturnInputs() []fyne.CanvasObject {
	var items []fyne.CanvasObject

	for _, item := range i.items {
		items = append(items, item.ReturnItems()...)
	}

	return items
}

func PraseStringToTime(s string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, s)
	if err == nil {
		return t, nil
	}

	t, err = time.Parse(time.RFC3339Nano, s)
	if err == nil {
		return t, nil
	}

	intT, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return time.Unix(intT, 0), nil
	}

	return time.Time{}, fmt.Errorf("invalid time format")
}

func main() {

	app := app.New()

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	mainWindow := app.NewWindow("Timestamp converter")
	mainWindow.Resize(fyne.NewSize(600, 100))

	inputs := Inputs{}
	inputs.MakeInputs()

	status := widget.NewLabel(fmt.Sprintf("[%s] Ready", time.Now().Format("15:04:05")))

	setStatus := func(s string) {
		status.SetText(fmt.Sprintf("[%s]: %s", time.Now().Format("15:04:05"), s))
	}

	inputs.UpdateInputs(time.Now())
	setStatus("Set to now")

	nowButton := widget.NewButton("Now", func() {
		inputs.UpdateInputs(time.Now())
		setStatus("Set to now")
	})

	fromCliboardButton := widget.NewButton("From clipboard", func() {
		clipboardContent := string(clipboard.Read(clipboard.FmtText))

		t, err := PraseStringToTime(clipboardContent)
		if err == nil {
			inputs.UpdateInputs(t)
			setStatus("Set to clipboard content")
			return
		}

		setStatus("Clipboard content is not a valid timestamp")
	})

	content := container.New(layout.NewVBoxLayout(), nowButton, fromCliboardButton, container.New(layout.NewGridLayout(3), inputs.ReturnInputs()...), status)

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}
