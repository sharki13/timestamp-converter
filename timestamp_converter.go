package main

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

type TimestampConverter struct {
	VisibleChanger map[int]binding.Bool
	Timestamp      binding.Untyped
	Format         binding.String
	Status         binding.String
	WachClipboard  bool
}

type TimestampItemsSet struct {
	DeleteBtnLabelContainer *fyne.Container
	EntryCopyBtnContainer   *fyne.Container
	Checkbox                *widget.Check
	Visible                 binding.Bool
}

func PraseStringToTime(s string) (time.Time, error) {
	for _, format := range SupportedFormats {
		t, err := time.Parse(format.Format, s)
		if err == nil {
			return t, nil
		}
	}

	intT, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return time.Unix(intT, 0), nil
	}

	return time.Time{}, fmt.Errorf("invalid time format")
}

func (t *TimestampConverter) SetStatus(status string) {
	now := time.Now()
	t.Status.Set(fmt.Sprintf("[%s]: %s", now.Format("15:04:05"), status))
}

func (t *TimestampConverter) MakeTimestampSetItmes(timezone TimezoneDefinition, window fyne.Window) TimestampItemsSet {
	deleteBtn := widget.NewButtonWithIcon("", theme.DeleteIcon(), nil)
	entry := widget.NewEntry()

	loc, err := time.LoadLocation(timezone.LocationAsString)
	if err != nil {
		panic(err)
	}

	entryListener := binding.NewDataListener(func() {
		timestampInt64, err := t.Timestamp.Get()
		if err != nil {
			panic(err)
		}

		format, err := t.Format.Get()
		if err != nil {
			panic(err)
		}

		timestamp := time.Unix(timestampInt64.(int64), 0)

		new_text := ""

		if timezone.Type != Unix {
			new_text = timestamp.In(loc).Format(format)
		} else {
			new_text = strconv.FormatInt(timestamp.Unix(), 10)
		}
		if new_text != entry.Text {
			entry.SetText(new_text)
		}
	})

	t.Timestamp.AddListener(entryListener)

	entryOnFormatChanged := binding.NewDataListener(func() {
		if timezone.Type != Unix {
			format, err := t.Format.Get()
			if err != nil {
				panic(err)
			}

			timestampInt64, err := t.Timestamp.Get()
			if err != nil {
				panic(err)
			}

			// parse timestamp to time.Time
			timestamp := time.Unix(timestampInt64.(int64), 0)

			new_content := timestamp.In(loc).Format(format)
			if new_content != entry.Text {
				entry.SetText(new_content)
			}
		}
	})

	t.Format.AddListener(entryOnFormatChanged)

	entry.OnChanged = func(text string) {
		timestamp, err := PraseStringToTime(text)
		if err != nil {
			t.SetStatus("Invalid timestamp")
			return
		}

		min_allowed_epoch := int64(0 - (50 * 31556926))

		if timestamp.Unix() <= min_allowed_epoch {
			t.SetStatus("Invalid timestamp")
			return
		}

		currentTimestamp, err := t.Timestamp.Get()
		if err != nil {
			panic(err)
		}

		currentTimestampUnix := currentTimestamp.(int64)
		timestampUnix := timestamp.Unix()

		if currentTimestampUnix != timestampUnix {
			t.Timestamp.Set(timestampUnix)
			t.SetStatus("Timestamp updated")
		}
	}

	label := widget.NewLabel(timezone.Label)
	copyBtn := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		window.Clipboard().SetContent(entry.Text)
		t.SetStatus("Copied to clipboard")
	})

	deleteBtnLabelContainer := container.NewHBox(deleteBtn, label)
	entryCopyBtnContainer := container.NewBorder(nil, nil, nil, copyBtn, entry)

	visibleBind := binding.NewBool()
	visibleHandler := binding.NewDataListener(func() {
		visible, err := visibleBind.Get()
		if err != nil {
			panic(err)
		}

		deleteBtnLabelContainer.Hidden = !visible
		entryCopyBtnContainer.Hidden = !visible
	})
	visibleBind.AddListener(visibleHandler)

	deleteBtn.OnTapped = func() {
		visibleBind.Set(false)
	}

	checkBox := widget.NewCheckWithData(timezone.Label, visibleBind)

	return TimestampItemsSet{
		DeleteBtnLabelContainer: deleteBtnLabelContainer,
		EntryCopyBtnContainer:   entryCopyBtnContainer,
		Checkbox:                checkBox,
		Visible:                 visibleBind,
	}
}

func GetThemeMenu(app fyne.App) *fyne.MenuItem {
	system := fyne.NewMenuItem("System", nil)
	light := fyne.NewMenuItem("Light", nil)
	dark := fyne.NewMenuItem("Dark", nil)

	if app.Settings().Theme() == theme.LightTheme() {
		light.Checked = true
	} else if app.Settings().Theme() == theme.DarkTheme() {
		dark.Checked = true
	} else {
		system.Checked = true
	}

	light.Action = func() {
		app.Settings().SetTheme(theme.LightTheme())
		light.Checked = true
		dark.Checked = false
		system.Checked = false
	}

	dark.Action = func() {
		app.Settings().SetTheme(theme.DarkTheme())
		light.Checked = false
		dark.Checked = true
		system.Checked = false
	}

	system.Action = func() {
		app.Settings().SetTheme(theme.DefaultTheme())
		light.Checked = false
		dark.Checked = false
		system.Checked = true
	}

	themeSubMenu := fyne.NewMenuItem("Theme", func() {})
	themeSubMenu.ChildMenu = fyne.NewMenu("", system, fyne.NewMenuItemSeparator(), light, dark)

	return themeSubMenu
}

func (t *TimestampConverter) SetupAndRun(window fyne.Window, app fyne.App) {
	t.Status = binding.NewString()
	t.SetStatus("Ready")
	t.VisibleChanger = make(map[int]binding.Bool)
	t.Timestamp = binding.NewUntyped()
	err := t.Timestamp.Set(time.Now().Unix())
	if err != nil {
		panic(err)
	}

	t.Format = binding.NewString()
	t.Format.Set(time.RFC3339)

	settingsMenu := fyne.NewMenu("Settings", GetThemeMenu(app))
	menu := fyne.NewMainMenu(settingsMenu)

	addEntry := xwidget.NewCompletionEntry([]string{})
	addEntry.PlaceHolder = "Add"
	// addEntry.OnChanged = func(timezone string) {
	// 	addEntry.SetOptions(t.Timezones)
	// 	addEntry.ShowCompletion()
	// }
	// addEntry.OnSubmitted = func(timezone string) {
	// 	// check if timezone key is present in map
	// 	for _, timeZoneDefinition := range Timezones {
	// 		if

	// 	if _, ok := t.VisibleChanger[timezone.id]; ok {
	// 		t.VisibleChanger[timezone].Set(true)
	// 	}
	// 	addEntry.SetText("")
	// 	addEntry.HideCompletion()
	// }

	nowBtn := widget.NewButtonWithIcon("Now", theme.ViewRefreshIcon(), func() {
		t.Timestamp.Set(time.Now().Unix())
		t.SetStatus("Updated to now")
	})
	nowBtn.Importance = widget.HighImportance

	leftSideToolbarItems := []fyne.CanvasObject{
		nowBtn,
	}

	rightSideToolbarItems := []fyne.CanvasObject{
		widget.NewCheck("Watch clipboard", func(checked bool) { t.WachClipboard = checked }),
		widget.NewButtonWithIcon("", theme.ContentPasteIcon(), func() {
			clipboardContent := window.Clipboard().Content()
			if clipboardContent == "" {
				return
			}

			timestamp, err := PraseStringToTime(clipboardContent)
			if err != nil {
				t.SetStatus("Invalid timestamp")
				return
			}

			t.Timestamp.Set(timestamp.Unix())
		}),
	}

	toolbar := container.NewBorder(nil, nil, container.NewHBox(leftSideToolbarItems...), container.NewHBox(rightSideToolbarItems...), addEntry)

	leftSide := container.NewVBox()
	middle := container.NewVBox()

	for _, timezone := range Timezones {
		items := t.MakeTimestampSetItmes(timezone, window)

		leftSide.Add(items.DeleteBtnLabelContainer)
		middle.Add(items.EntryCopyBtnContainer)
		items.Visible.Set(true)

		// add to visible changer
		t.VisibleChanger[timezone.Id] = items.Visible
	}

	status := widget.NewLabelWithData(t.Status)

	container := container.NewBorder(toolbar, status, leftSide, nil, middle)

	go func() {
		for {
			time.Sleep(time.Second)
			if t.WachClipboard {
				cliboardContent := window.Clipboard().Content()
				if cliboardContent == "" {
					continue
				}

				timestamp, err := PraseStringToTime(cliboardContent)
				if err != nil {
					continue
				}

				t.Timestamp.Set(timestamp.Unix())
				t.SetStatus("Updated from clipboard")
			}
		}
	}()

	window.SetMainMenu(menu)
	window.SetContent(container)
	window.Resize(fyne.NewSize(600, 10))
	window.ShowAndRun()
}
