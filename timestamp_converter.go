package main

import (
	"fmt"
	"net/url"
	"runtime"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
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

		if timezone.Type != UnixTimezoneType {
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
		if timezone.Type != UnixTimezoneType {
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

	if timezone.Type == LocalTimezoneType {
		deleteBtn.Disable()
	}

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

func GetFormatMenu(t *TimestampConverter) *fyne.Menu {

	formatMenu := fyne.NewMenu("Format", make([]*fyne.MenuItem, 0)...)

	for _, format := range SupportedFormats {
		format := format
		formatMenu.Items = append(formatMenu.Items, fyne.NewMenuItem(format.Label, func() {
			t.Format.Set(format.Format)

			for _, item := range formatMenu.Items {
				if item.Label != format.Label {
					item.Checked = false
				} else {
					item.Checked = true
				}
			}
		}))
	}

	if len(formatMenu.Items) == 0 {
		panic("no format found")
	}

	formatMenu.Items[0].Checked = true
	t.Format.Set(SupportedFormats[0].Format)

	return formatMenu
}

func contains[K comparable](s []K, e K) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetPresetsMenu(t *TimestampConverter) *fyne.Menu {
	presetsMenu := fyne.NewMenu("Presets", make([]*fyne.MenuItem, 0)...)

	for _, preset := range TimezonePresets {
		preset := preset
		presetsMenu.Items = append(presetsMenu.Items, fyne.NewMenuItem(preset.Label, func() {

			for _, currentTimezoneVisibility := range t.VisibleChanger {
				currentTimezoneVisibility.Set(false)
			}

			for _, item := range presetsMenu.Items {
				if item.Label != preset.Label {
					item.Checked = false
				} else {
					item.Checked = true
				}
			}

			for _, presetDef := range TimezonePresets {
				if presetDef.Id == preset.Id {
					for _, id := range presetDef.Timezones {
						// check if id key exists
						if _, ok := t.VisibleChanger[id]; !ok {
							continue
						}

						t.VisibleChanger[id].Set(true)
					}
				}
			}

			t.SetStatus(fmt.Sprintf("Preset %s", preset.Label))
		}))
	}

	if len(presetsMenu.Items) == 0 {
		panic("no presets found")
	}

	presetsMenu.Items[0].Checked = true
	presetsMenu.Items[0].Action()

	return presetsMenu
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
			t.SetStatus("Updated to clipboard content")
		}),
	}

	toolbar := container.NewBorder(nil, nil, container.NewHBox(leftSideToolbarItems...), container.NewHBox(rightSideToolbarItems...), widget.NewLabel(""))

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

	scrollableMiddle := container.NewVScroll(container.NewBorder(nil, nil, leftSide, nil, middle))
	mainContainer := container.NewBorder(toolbar, status, nil, nil, scrollableMiddle)

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

				currentTimestamp, err := t.Timestamp.Get()
				if err != nil {
					panic(err)
				}

				if timestamp.Unix() == currentTimestamp.(int64) {
					continue
				}

				t.Timestamp.Set(timestamp.Unix())
				t.SetStatus("Updated from clipboard")
			}
		}
	}()

	about := fyne.NewMenuItem("GitHub page", func() {
		u, _ := url.Parse("https://github.com/sharki13/timestamp-converter")
		_ = app.OpenURL(u)
	})

	infoMenu := fyne.NewMenu("Help", about)

	openSettings := func() {
		w := app.NewWindow("Scale and Appearance")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	}
	settingsItem := fyne.NewMenuItem("Scale and Appearance", openSettings)
	settingsItem.Icon = theme.SettingsIcon()

	settingsMenu := fyne.NewMenu("Settings", settingsItem, GetThemeMenu(app))
	formatMenu := GetFormatMenu(t)
	presetMenu := GetPresetsMenu(t)
	menu := fyne.NewMainMenu(make([]*fyne.Menu, 0)...)

	if runtime.GOOS != "darwin" {
		fileMenu := fyne.NewMenu("File", fyne.NewMenuItem("Quit", func() {
			app.Quit()
		}))

		menu.Items = append(menu.Items, fileMenu)
	}

	menu.Items = append(menu.Items, presetMenu, formatMenu, settingsMenu, infoMenu)

	window.SetMainMenu(menu)
	window.SetContent(mainContainer)
	window.Resize(fyne.NewSize(600, 400))
	window.ShowAndRun()
}
