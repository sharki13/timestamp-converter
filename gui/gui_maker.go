package gui

import (
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
	"github.com/sharki13/timestamp-converter/timezone"
)

type timestampItemsSet struct {
	deleteBtnLabelContainer *fyne.Container
	entryCopyBtnContainer   *fyne.Container
	visible                 binding.Bool
}

func (t *TimestampConverter) makeCopyButtonForEntry(entry *widget.Entry) *widget.Button {
	return widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		clip := t.window.Clipboard()

		if clip == nil {
			return
		}

		clip.SetContent(entry.Text)
	})
}

func (t *TimestampConverter) newTimestampSetItems(tz timezone.TimezoneDefinition, window fyne.Window) timestampItemsSet {
	timestampEntry := widget.NewEntry()

	timestampEntry.OnChanged = func(text string) {
		timestamp, err := praseStringToTime(text)
		if err != nil {
			return
		}

		currentTimestamp, err := t.timestamp.Get()
		if err != nil {
			panic(err)
		}

		if currentTimestamp != timestamp {
			t.timestamp.Set(timestamp)
		}
	}

	timestampEntry.Validator = func(text string) error {
		_, err := praseStringToTime(text)
		if err != nil {
			return err
		}

		return nil
	}

	onFormatOrTimestampChange := binding.NewDataListener(func() {
		timestamp, err := t.timestamp.Get()
		if err != nil {
			panic(err)
		}

		format, err := t.format.Get()
		if err != nil {
			panic(err)
		}

		new_text := tz.StringTime(timestamp, format)

		if new_text != timestampEntry.Text {
			timestampEntry.SetText(new_text)
		}
	})

	t.timestamp.AddListener(onFormatOrTimestampChange)
	t.format.AddListener(onFormatOrTimestampChange)

	visibleState := binding.NewBool()

	deleteBtn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		visibleState.Set(false)

		for id, state := range t.timezonesVisibleState {
			visibleIds := make([]int, 0)

			visible, _ := state.Get()
			if visible {
				visibleIds = append(visibleIds, id)
			}

			t.visibleTimezones.Set(visibleIds)
		}
	})

	if tz.Type == timezone.LocalTimezoneType {
		deleteBtn.Disable()
	}

	deleteBtnLabelContainer := container.NewHBox(deleteBtn, widget.NewLabel(tz.Label))

	entryCopyBtnContainer := container.NewBorder(nil, nil, nil, t.makeCopyButtonForEntry(timestampEntry), timestampEntry)

	visibleHandler := binding.NewDataListener(func() {
		visible, err := visibleState.Get()
		if err != nil {
			panic(err)
		}

		deleteBtnLabelContainer.Hidden = !visible
		entryCopyBtnContainer.Hidden = !visible
	})
	visibleState.AddListener(visibleHandler)

	return timestampItemsSet{
		deleteBtnLabelContainer: deleteBtnLabelContainer,
		entryCopyBtnContainer:   entryCopyBtnContainer,
		visible:                 visibleState,
	}
}

func (t *TimestampConverter) getOptions(text string) []string {
	options := []string{}

	visibleIds := []int{}

	for k, e := range t.timezonesVisibleState {
		visible, _ := e.Get()
		if visible {
			visibleIds = append(visibleIds, k)
		}
	}

	for _, timeZoneDefinition := range timezone.Timezones {
		if strings.Contains(strings.ToLower(timeZoneDefinition.Label), strings.ToLower(text)) && !contains(visibleIds, timeZoneDefinition.Id) {
			options = append(options, timeZoneDefinition.Label)
		}
	}

	return options
}

func (t *TimestampConverter) newTimezoneAddEntry() *xwidget.CompletionEntry {
	entry := xwidget.NewCompletionEntry([]string{})
	entry.PlaceHolder = "Add"

	entry.OnChanged = func(text string) {
		entry.SetOptions(t.getOptions(text))
		entry.ShowCompletion()
	}

	entry.OnSubmitted = func(string) {
		if len(entry.Options) != 0 {
			for _, timeZoneDefinition := range timezone.Timezones {
				if timeZoneDefinition.Label == entry.Options[0] {
					t.timezonesVisibleState[timeZoneDefinition.Id].Set(true)
					break
				}
			}

			visibleTimezones := make([]int, 0)

			for k, e := range t.timezonesVisibleState {
				visible, _ := e.Get()
				if visible {
					visibleTimezones = append(visibleTimezones, k)
				}
			}

			t.visibleTimezones.Set(visibleTimezones)

			entry.SetText("")
			entry.HideCompletion()
		}
	}

	// due to bug in this widget: https://github.com/fyne-io/fyne-x/issues/38
	entry.CustomUpdate = func(i widget.ListItemID, o fyne.CanvasObject) {
		options := entry.Options

		if i >= len(options) {
			return
		}

		o.(*widget.Label).SetText(options[i])
	}

	return entry
}

func (t *TimestampConverter) newToolbar() *fyne.Container {
	nowBtn := widget.NewButtonWithIcon("Now", theme.ViewRefreshIcon(), func() {
		t.timestamp.Set(time.Now())
	})
	nowBtn.Importance = widget.HighImportance

	leftSideToolbarItems := []fyne.CanvasObject{
		nowBtn,
	}

	rightSideToolbarItems := []fyne.CanvasObject{
		widget.NewCheck("Watch clipboard", func(checked bool) { t.watchClipboard = checked }),
		widget.NewButtonWithIcon("", theme.ContentPasteIcon(), func() {
			clip := t.window.Clipboard()

			if clip == nil {
				return
			}

			clipboardContent := clip.Content()
			if clipboardContent == "" {
				return
			}

			timestamp, err := praseStringToTime(clipboardContent)
			if err != nil {
				return
			}

			t.timestamp.Set(timestamp)
		}),
	}

	return container.NewBorder(nil, nil, container.NewHBox(leftSideToolbarItems...), container.NewHBox(rightSideToolbarItems...), t.newTimezoneAddEntry())
}

func (t *TimestampConverter) makeContent() *fyne.Container {
	leftSide := container.NewVBox()
	middle := container.NewVBox()

	for _, tz := range timezone.Timezones {
		items := t.newTimestampSetItems(tz, t.window)

		leftSide.Add(items.deleteBtnLabelContainer)
		middle.Add(items.entryCopyBtnContainer)

		if tz.Type == timezone.LocalTimezoneType {
			items.visible.Set(true)
		} else {
			items.visible.Set(false)
		}

		// add to visible changer
		t.timezonesVisibleState[tz.Id] = items.visible
	}

	scrollableMiddle := container.NewVScroll(container.NewBorder(nil, nil, leftSide, nil, middle))
	return container.NewBorder(t.newToolbar(), nil, nil, nil, scrollableMiddle)
}
