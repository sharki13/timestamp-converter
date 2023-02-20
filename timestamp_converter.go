package main

import (
	"strings"
	"time"

	prefSync "com.sharki13/timestamp.converter/preferences"
	"com.sharki13/timestamp.converter/timezone"
	"com.sharki13/timestamp.converter/xbinding"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

type TimestampConverter struct {
	timezonesVisibleState map[int]binding.Bool
	visibleTimezones      xbinding.IntArray
	timestamp             xbinding.Time
	format                binding.String
	watchClipboard        bool
	window                fyne.Window
	app                   fyne.App
	preferences           *prefSync.PreferencesSynchronizer
}

func NewTimestampConverter(app fyne.App) *TimestampConverter {
	ret := TimestampConverter{
		app:    app,
		window: app.NewWindow("Timestamp converter"),
	}

	ret.timezonesVisibleState = make(map[int]binding.Bool, 0)
	ret.timestamp = xbinding.NewTime()
	ret.visibleTimezones = xbinding.NewIntArray()
	ret.timestamp.Set(time.Now())
	ret.format = binding.NewString()
	ret.preferences = prefSync.NewPreferencesSynchronizer(app)

	return &ret
}

// Sets up the preferences and loads them
// from the fyne preferences
// Might panic if keys are not unique
func (t *TimestampConverter) SetupAndLoadPreferences() {
	err := t.preferences.AddString(prefSync.StringPreference{
		Key:      "format",
		Value:    t.format,
		Fallback: time.RFC3339,
	})

	if err != nil {
		panic(err)
	}

	err = t.preferences.AddIntArray(prefSync.IntArrayPreference{
		Key:      "visibleTimezones",
		Value:    t.visibleTimezones,
		Fallback: []int{0},
	})

	savedTimezones, _ := t.visibleTimezones.Get()
	for _, timezoneIndex := range savedTimezones {
		if visible, ok := t.timezonesVisibleState[timezoneIndex]; ok {
			visible.Set(true)
		}
	}

	if err != nil {
		panic(err)
	}
}

type TimestampItemsSet struct {
	DeleteBtnLabelContainer *fyne.Container
	EntryCopyBtnContainer   *fyne.Container
	Visible                 binding.Bool
}

func (t *TimestampConverter) AttachEntryToFormatOrTimestampChange(entry *widget.Entry, timezoneDefinition timezone.TimezoneDefinition) {
	onFormatOrTimestampChange := binding.NewDataListener(func() {
		timestamp, err := t.timestamp.Get()
		if err != nil {
			panic(err)
		}

		format, err := t.format.Get()
		if err != nil {
			panic(err)
		}

		new_text := timezoneDefinition.StringTime(timestamp, format)

		if new_text != entry.Text {
			entry.SetText(new_text)
		}
	})

	t.timestamp.AddListener(onFormatOrTimestampChange)
	t.format.AddListener(onFormatOrTimestampChange)
}

func (t *TimestampConverter) MakeCopyButtonForEntry(entry *widget.Entry, window fyne.Window) *widget.Button {
	return widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		clip := window.Clipboard()

		if clip == nil {
			return
		}

		clip.SetContent(entry.Text)
	})
}

func RemoveIntFromSlice(slice []int, value int) []int {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}

func (t *TimestampConverter) NewTimestampSetItems(tz timezone.TimezoneDefinition, window fyne.Window) TimestampItemsSet {
	entry := widget.NewEntry()
	t.AttachEntryToFormatOrTimestampChange(entry, tz)

	entry.OnChanged = func(text string) {
		timestamp, err := PraseStringToTime(text)
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

	entry.Validator = func(text string) error {
		_, err := PraseStringToTime(text)
		if err != nil {
			return err
		}

		return nil
	}

	visibleBind := binding.NewBool()

	deleteBtn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		visibleBind.Set(false)

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

	label := widget.NewLabel(tz.Label)
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

	entry.OnChanged = func(timezone string) {
		options := getOptions(timezone)

		entry.SetOptions(options)
		entry.ShowCompletion()
	}

	entry.OnSubmitted = func(tz string) {
		options := getOptions(tz)

		if len(options) != 0 {
			for _, timeZoneDefinition := range timezone.Timezones {
				if timeZoneDefinition.Label == options[0] {
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

func (t *TimestampConverter) NewToolbar(window fyne.Window) *fyne.Container {
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
			clip := window.Clipboard()

			if clip == nil {
				return
			}

			clipboardContent := clip.Content()
			if clipboardContent == "" {
				return
			}

			timestamp, err := PraseStringToTime(clipboardContent)
			if err != nil {
				return
			}

			t.timestamp.Set(timestamp)
		}),
	}

	return container.NewBorder(nil, nil, container.NewHBox(leftSideToolbarItems...), container.NewHBox(rightSideToolbarItems...), t.NewCompletionAddEntry())
}

func (t *TimestampConverter) SetupAndRun() {
	leftSide := container.NewVBox()
	middle := container.NewVBox()

	for _, tz := range timezone.Timezones {
		items := t.NewTimestampSetItems(tz, t.window)

		leftSide.Add(items.DeleteBtnLabelContainer)
		middle.Add(items.EntryCopyBtnContainer)

		if tz.Type == timezone.LocalTimezoneType {
			items.Visible.Set(true)
		} else {
			items.Visible.Set(false)
		}

		// add to visible changer
		t.timezonesVisibleState[tz.Id] = items.Visible
	}

	scrollableMiddle := container.NewVScroll(container.NewBorder(nil, nil, leftSide, nil, middle))
	mainContainer := container.NewBorder(t.NewToolbar(t.window), nil, nil, nil, scrollableMiddle)

	go func() {
		for {
			time.Sleep(time.Second)
			if t.watchClipboard {
				clip := t.window.Clipboard()

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

				currentTimestamp, err := t.timestamp.Get()
				if err != nil {
					panic(err)
				}

				if timestamp == currentTimestamp {
					continue
				}

				t.timestamp.Set(timestamp)
			}
		}
	}()

	t.window.SetMainMenu(t.MakeMenu())
	t.window.SetContent(mainContainer)
	t.SetupAndLoadPreferences()
	t.window.Resize(fyne.NewSize(600, 400))
	t.window.ShowAndRun()
}
