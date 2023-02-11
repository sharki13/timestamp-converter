package main

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
	"golang.design/x/clipboard"
)

func log(text string) {
	now := time.Now()
	timestamp := now.UTC().Format(time.RFC3339)

	fmt.Printf("[%s] %s\n", timestamp, text)
}

type Inputs struct {
	unixTimestampEntry *widget.Entry
	utcRFC3339Entry    *widget.Entry
	localRFC3339Entry  *widget.Entry
	estRFC3339Entry    *widget.Entry
	pstRFC3339Entry    *widget.Entry

	unixTimestampLabel *widget.Label
	utcRFC3339Label    *widget.Label
	localRFC3339Label  *widget.Label
	estRFC3339Label    *widget.Label
	pstRFC3339Label    *widget.Label
}

func (i *Inputs) MakeInputs() {
	now := time.Now()

	i.unixTimestampEntry = widget.NewEntry()
	i.utcRFC3339Entry = widget.NewEntry()
	i.localRFC3339Entry = widget.NewEntry()
	i.estRFC3339Entry = widget.NewEntry()
	i.pstRFC3339Entry = widget.NewEntry()

	i.unixTimestampLabel = widget.NewLabel("Unix")
	i.utcRFC3339Label = widget.NewLabel("UTC")
	i.localRFC3339Label = widget.NewLabel("Local")
	i.estRFC3339Label = widget.NewLabel("EST")
	i.pstRFC3339Label = widget.NewLabel("PST")

	style := fyne.TextStyle{
		Bold: 	true,
	}

	i.unixTimestampLabel.TextStyle = style
	i.utcRFC3339Label.TextStyle = style
	i.localRFC3339Label.TextStyle = style
	i.estRFC3339Label.TextStyle = style
	i.pstRFC3339Label.TextStyle = style

	i.unixTimestampEntry.TextStyle = style
	i.utcRFC3339Entry.TextStyle = style
	i.localRFC3339Entry.TextStyle = style
	i.estRFC3339Entry.TextStyle = style
	i.pstRFC3339Entry.TextStyle = style

	i.unixTimestampEntry.SetText(strconv.FormatInt(now.Unix(), 10))
	i.utcRFC3339Entry.SetText(now.UTC().Format(time.RFC3339))
	i.localRFC3339Entry.SetText(now.Local().Format(time.RFC3339))
	i.estRFC3339Entry.SetText(now.In(time.FixedZone("EST", -5*60*60)).Format(time.RFC3339))
	i.pstRFC3339Entry.SetText(now.In(time.FixedZone("PST", -8*60*60)).Format(time.RFC3339))
}

// function which takes a time.Time and updates the inputs
func (i *Inputs) UpdateInputs(t time.Time) {
	i.unixTimestampEntry.SetText(strconv.FormatInt(t.Unix(), 10))
	i.utcRFC3339Entry.SetText(t.UTC().Format(time.RFC3339))
	i.localRFC3339Entry.SetText(t.Local().Format(time.RFC3339))
	i.estRFC3339Entry.SetText(t.In(time.FixedZone("EST", -5*60*60)).Format(time.RFC3339))
	i.pstRFC3339Entry.SetText(t.In(time.FixedZone("PST", -8*60*60)).Format(time.RFC3339))
}

// return all inputs as a slice of fyne widgets
func (i *Inputs) ReturnInputs() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		i.unixTimestampLabel, i.unixTimestampEntry,
		i.utcRFC3339Label, i.utcRFC3339Entry,
		i.localRFC3339Label, i.localRFC3339Entry,
		i.estRFC3339Label, i.estRFC3339Entry,
		i.pstRFC3339Label, i.pstRFC3339Entry,
	}
}


func main() {

	app := app.New()
	mainWindow := app.NewWindow("Timestamp converter")
	mainWindow.Resize(fyne.NewSize(400, 100))

	inputs := Inputs{}
	inputs.MakeInputs()

	nowButton := widget.NewButton("Now", func() {
		inputs.UpdateInputs(time.Now())
	})

	fromCliboardButton := widget.NewButton("From clipboard", func() {
		err := clipboard.Init()
		if err != nil {
			panic(err)
		}

		clipboardContent := string(clipboard.Read(clipboard.FmtText))
		log("Clipboard content: " + clipboardContent)

		t, err := time.Parse(time.RFC3339, clipboardContent)
		if err == nil {
			log(fmt.Sprintf("Clipboard content is a valid RFC3339 timestamp, %s", t.Format(time.RFC3339)))
			inputs.UpdateInputs(t)
			return
		}

		t, err = time.Parse(time.RFC3339Nano, clipboardContent)
		if err == nil {
			log(fmt.Sprintf("Clipboard content is a valid RFC3339Nano timestamp, %s", t.Format(time.RFC3339Nano)))
			inputs.UpdateInputs(t)
			return
		}

		intT, err := strconv.ParseInt(clipboardContent, 10, 64)
		if err == nil {
			log(fmt.Sprintf("Clipboard content is a valid Unix timestamp, %d", intT))
			inputs.UpdateInputs(time.Unix(intT, 0))
			return
		}

		log("Clipboard content is not a valid timestamp")
	})

	content := container.New(layout.NewVBoxLayout(), nowButton, fromCliboardButton, container.New(layout.NewGridLayout(2), inputs.ReturnInputs()...) )

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}
