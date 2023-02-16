package main

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

type TimestampConverter struct {
	TimezonesVisbleState map[int]binding.Bool
	Timestamp            BoundTime
	Format               binding.String
	Status               binding.String
	WatchClipboard       bool
	Preset               binding.Int
}

// Intialize bindings, have to be called at the very beginning,
// otherwise code might try to reach bindings that are not yet initialized
func (t *TimestampConverter) CreateBindings() {
	t.TimezonesVisbleState = make(map[int]binding.Bool)
	t.Timestamp = NewBoundTime()
	t.Format = binding.NewString()
	t.Status = binding.NewString()
	t.Preset = binding.NewInt()
}

// Should be called after UI setup
func (t *TimestampConverter) BindStateToPreferencesAndUI(app fyne.App) {
	format := app.Preferences().StringWithFallback("format", time.RFC3339)

	t.Format.AddListener(binding.NewDataListener(func() {
		f, err := t.Format.Get()
		if err != nil {
			panic(err)
		}

		app.Preferences().SetString("format", f)
		t.SetStatus(fmt.Sprintf("Format set to: %s", f))
	}))

	t.Format.Set(format)

	preset := app.Preferences().IntWithFallback("preset", 1)

	t.Preset.AddListener(binding.NewDataListener(func() {
		p, err := t.Preset.Get()
		if err != nil {
			panic(err)
		}

		app.Preferences().SetInt("preset", p)

		for _, e := range t.TimezonesVisbleState {
			e.Set(false)
		}

		for _, presetDef := range TimezonePresets {
			if presetDef.Id == p {
				for _, id := range presetDef.Timezones {
					// check if id key exists
					if _, ok := t.TimezonesVisbleState[id]; !ok {
						continue
					}

					t.TimezonesVisbleState[id].Set(true)
				}
				t.SetStatus(fmt.Sprintf("Preset %s", presetDef.Label))
			}
		}
	}))

	t.Preset.Set(preset)

	t.Timestamp.AddListener(binding.NewDataListener(func() {
		t.SetStatus("Timestamp updated")
	}))
}

type TimestampItemsSet struct {
	DeleteBtnLabelContainer *fyne.Container
	EntryCopyBtnContainer   *fyne.Container
	Visible                 binding.Bool
}

func (t *TimestampConverter) SetStatus(status string) {
	now := time.Now()
	t.Status.Set(fmt.Sprintf("[%s]: %s", now.Format("15:04:05"), status))
}

func (t *TimestampConverter) AttachEntryToFormatOrTimestampChange(entry *widget.Entry, timezoneDefinition TimezoneDefinition) {
	onFormatOrTimestampChange := binding.NewDataListener(func() {
		timestamp, err := t.Timestamp.Get()
		if err != nil {
			panic(err)
		}

		format, err := t.Format.Get()
		if err != nil {
			panic(err)
		}

		new_text := timezoneDefinition.StringTime(timestamp, format)

		if new_text != entry.Text {
			entry.SetText(new_text)
		}
	})

	t.Timestamp.AddListener(onFormatOrTimestampChange)
	t.Format.AddListener(onFormatOrTimestampChange)
}

func (t *TimestampConverter) MakeCopyButtonForEntry(entry *widget.Entry, window fyne.Window) *widget.Button {
	return widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		clip := window.Clipboard()

		if clip == nil {
			return
		}

		clip.SetContent(entry.Text)
		t.SetStatus("Copied to clipboard")
	})
}

func (t *TimestampConverter) NewTimestampSetItems(timezone TimezoneDefinition, window fyne.Window) TimestampItemsSet {
	entry := widget.NewEntry()
	t.AttachEntryToFormatOrTimestampChange(entry, timezone)

	entry.OnChanged = func(text string) {
		timestamp, err := PraseStringToTime(text)
		if err != nil {
			t.SetStatus("Invalid timestamp")
			return
		}

		// due to issues with parsing time by Go, expierementally found that
		min_allowed_epoch := int64(0 - (50 * 31556926))

		if timestamp.Unix() <= min_allowed_epoch {
			t.SetStatus("Invalid timestamp")
			return
		}

		currentTimestamp, err := t.Timestamp.Get()
		if err != nil {
			panic(err)
		}

		if currentTimestamp != timestamp {
			t.Timestamp.Set(timestamp)
		}
	}

	visibleBind := binding.NewBool()

	deleteBtn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		visibleBind.Set(false)
	})

	if timezone.Type == LocalTimezoneType {
		deleteBtn.Disable()
	}

	label := widget.NewLabel(timezone.Label)
	deleteBtnLabelContainer := container.NewHBox(deleteBtn, label)

	entryCopyBtnContainer := container.NewBorder(nil, nil, nil, t.MakeCopyButtonForEntry(entry, window), entry)

	visibleHandler := binding.NewDataListener(func() {
		visible, err := visibleBind.Get()
		if err != nil {
			panic(err)
		}

		deleteBtnLabelContainer.Hidden = !visible
		entryCopyBtnContainer.Hidden = !visible
	})
	visibleBind.AddListener(visibleHandler)

	return TimestampItemsSet{
		DeleteBtnLabelContainer: deleteBtnLabelContainer,
		EntryCopyBtnContainer:   entryCopyBtnContainer,
		Visible:                 visibleBind,
	}
}

func (t *TimestampConverter) NewCompletionAddEntry() *xwidget.CompletionEntry {
	entry := xwidget.NewCompletionEntry([]string{})
	entry.PlaceHolder = "Add"

	getOptions := func(text string) []string {
		options := []string{}

		visibleIds := []int{}

		for k, e := range t.TimezonesVisbleState {
			visible, _ := e.Get()
			if visible {
				visibleIds = append(visibleIds, k)
			}
		}

		for _, timeZoneDefinition := range Timezones {
			if strings.Contains(strings.ToLower(timeZoneDefinition.Label), strings.ToLower(text)) && !contains(visibleIds, timeZoneDefinition.Id) {
				options = append(options, timeZoneDefinition.Label)
			}
		}

		return options
	}

	entry.OnChanged = func(timezone string) {
		options := getOptions(timezone)

		entry.SetOptions(options)
		entry.ShowCompletion()
	}

	entry.OnSubmitted = func(timezone string) {
		options := getOptions(timezone)

		if len(options) != 0 {
			for _, timeZoneDefinition := range Timezones {
				if timeZoneDefinition.Label == options[0] {
					t.TimezonesVisbleState[timeZoneDefinition.Id].Set(true)
					break
				}
			}

			t.SetStatus(fmt.Sprintf("Added %s", options[0]))
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

func (t *TimestampConverter) NewToolbar(window fyne.Window) *fyne.Container {
	nowBtn := widget.NewButtonWithIcon("Now", theme.ViewRefreshIcon(), func() {
		t.Timestamp.Set(time.Now())
		t.SetStatus("Updated to now")
	})
	nowBtn.Importance = widget.HighImportance

	leftSideToolbarItems := []fyne.CanvasObject{
		nowBtn,
	}

	rightSideToolbarItems := []fyne.CanvasObject{
		widget.NewCheck("Watch clipboard", func(checked bool) { t.WatchClipboard = checked }),
		widget.NewButtonWithIcon("", theme.ContentPasteIcon(), func() {
			clip := window.Clipboard()

			if clip == nil {
				t.SetStatus("Clipboard not initialized")
				return
			}

			clipboardContent := clip.Content()
			if clipboardContent == "" {
				return
			}

			timestamp, err := PraseStringToTime(clipboardContent)
			if err != nil {
				t.SetStatus("Invalid timestamp")
				return
			}

			t.Timestamp.Set(timestamp)
			t.SetStatus("Updated from clipboard")
		}),
	}

	return container.NewBorder(nil, nil, container.NewHBox(leftSideToolbarItems...), container.NewHBox(rightSideToolbarItems...), t.NewCompletionAddEntry())
}

func (t *TimestampConverter) SetupAndRun(window fyne.Window, app fyne.App) {
	t.CreateBindings()
	t.Timestamp.Set(time.Now())

	status := widget.NewLabelWithData(t.Status)
	t.SetStatus("Ready")

	leftSide := container.NewVBox()
	middle := container.NewVBox()

	for _, timezone := range Timezones {
		items := t.NewTimestampSetItems(timezone, window)

		leftSide.Add(items.DeleteBtnLabelContainer)
		middle.Add(items.EntryCopyBtnContainer)
		items.Visible.Set(true)

		// add to visible changer
		t.TimezonesVisbleState[timezone.Id] = items.Visible
	}

	scrollableMiddle := container.NewVScroll(container.NewBorder(nil, nil, leftSide, nil, middle))
	mainContainer := container.NewBorder(t.NewToolbar(window), status, nil, nil, scrollableMiddle)

	go func() {
		for {
			time.Sleep(time.Second)
			if t.WatchClipboard {
				clip := window.Clipboard()

				if clip == nil {
					continue
				}

				cliboardContent := clip.Content()
				if cliboardContent == "" {
					continue
				}

				timestamp, err := PraseStringToTime(cliboardContent)
				if err != nil {
					continue
				}

				currentTimestamp, err := t.Timestamp.Get()
				if err != nil {
					panic(err)
				}

				if timestamp == currentTimestamp {
					continue
				}

				t.Timestamp.Set(timestamp)
				t.SetStatus("Updated from clipboard")
			}
		}
	}()

	t.BindStateToPreferencesAndUI(app)
	window.SetMainMenu(t.MakeMenu(app))
	window.SetContent(mainContainer)
	window.Resize(fyne.NewSize(600, 400))
	window.ShowAndRun()
}
