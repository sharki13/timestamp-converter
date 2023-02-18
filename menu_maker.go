package main

import (
	"errors"
	"fmt"
	"net/url"
	"runtime"

	"com.sharki13/timestamp.converter/timezone"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (t *TimestampConverter) MakeMenu(app fyne.App) *fyne.MainMenu {

	menus := make([]*fyne.Menu, 0)

	// Mac OS has a built in quit menu,
	// on other platforms Fyne will add Quit to first menu if it is not defined
	if runtime.GOOS != "darwin" {
		fileMenu := fyne.NewMenu("File", fyne.NewMenuItem("Quit", func() {
			app.Quit()
		}))

		menus = append(menus, fileMenu)
	}

	t.presetMenu = t.MakePresetMenu()

	menus = append(menus,
		t.presetMenu,
		t.MakeFormatMenu(app),
		t.MakeSettingsMenu(app),
		t.MakeInfoMenu(app),
	)

	return fyne.NewMainMenu(menus...)
}

func (t *TimestampConverter) MakeInfoMenu(app fyne.App) *fyne.Menu {
	about := fyne.NewMenuItem("GitHub page", func() {
		u, _ := url.Parse("https://github.com/sharki13/timestamp-converter")
		_ = app.OpenURL(u)
	})

	return fyne.NewMenu("Help", about)
}

func (t *TimestampConverter) MakeSettingsMenu(app fyne.App) *fyne.Menu {
	openSettings := func() {
		w := app.NewWindow("Scale and Appearance")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	}
	settingsItem := fyne.NewMenuItem("Scale and Appearance", openSettings)
	settingsItem.Icon = theme.SettingsIcon()

	return fyne.NewMenu("Settings", settingsItem, t.MakeThemeMenu(app))
}

func (t *TimestampConverter) MakeThemeMenu(app fyne.App) *fyne.MenuItem {
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

func (t *TimestampConverter) MakeFormatMenu(app fyne.App) *fyne.Menu {

	formatMenu := fyne.NewMenu("Format", make([]*fyne.MenuItem, 0)...)

	for _, format := range SupportedFormats {
		format := format
		formatMenuItem := fyne.NewMenuItem(format.Label, func() {
			t.format.Set(format.Format)
		})

		t.format.AddListener(binding.NewDataListener(func() {
			currentFormat, err := t.format.Get()
			if err != nil {
				panic(err)
			}

			if currentFormat == format.Format {
				formatMenuItem.Checked = true
			} else {
				formatMenuItem.Checked = false
			}
		}))

		formatMenu.Items = append(formatMenu.Items, formatMenuItem)
	}

	return formatMenu
}

func makePresetMenuItem(label string, id int, currentPresetBound binding.Int) *fyne.MenuItem {
	presetMenuItem := fyne.NewMenuItem(label, func() {
		currentPresetBound.Set(id)
	})

	currentPresetBound.AddListener(binding.NewDataListener(func() {
		currentPreset, err := currentPresetBound.Get()
		if err != nil {
			panic(err)
		}

		if currentPreset == id {
			presetMenuItem.Checked = true
		} else {
			presetMenuItem.Checked = false
		}
	}))

	return presetMenuItem
}

func (t *TimestampConverter) MakePresetMenu() *fyne.Menu {
	items := make([]*fyne.MenuItem, 0)

	for _, preset := range timezone.DefaultPresets {
		preset := preset
		items = append(items, makePresetMenuItem(preset.Label, preset.Id, t.preset))
	}

	items = append(items, fyne.NewMenuItemSeparator())

	userPresets, _ := t.userPresets.Get()

	if len(userPresets) == 0 {
		noneItem := fyne.NewMenuItem("(None)", func() {})
		noneItem.Disabled = true
		items = append(items, noneItem)
	} else {
		for _, preset := range userPresets {
			preset := preset
			items = append(items, makePresetMenuItem(preset.Label, preset.Id, t.preset))
		}
	}

	addPereset := fyne.NewMenuItem("Add current as preset", t.MakeAndShowAddPresetWindow)

	removePreset := fyne.NewMenuItem("Remove current preset", func() {
		fmt.Printf("Remove current preset")
	})

	removePreset.Disabled = true

	items = append(items, addPereset, removePreset)

	return fyne.NewMenu("Presets", items...)
}

func (t *TimestampConverter) MakeAndShowAddPresetWindow() {
	w := t.app.NewWindow("Add Preset")

	presetName := widget.NewEntry()

	presetName.Validator = func(s string) error {
		if len(s) == 0 {
			return errors.New("Preset name cannot be empty")
		}

		return nil
	}

	errorLabel := widget.NewLabel("")

	presetName.SetOnValidationChanged(func(err error) {
		if err != nil {
			errorLabel.SetText(err.Error())
		} else {
			errorLabel.SetText("")
		}
	})

	okBtn := widget.NewButton("OK", func() {
		userPresets, _ := t.userPresets.Get()

		activeTimezones := []int{}

		for k, tz := range t.timezonesVisibleState {
			visible, _ := tz.Get()
			if visible {
				activeTimezones = append(activeTimezones, k)
			}
		}

		preset := timezone.Preset{
			Id:        timezone.LastInternalId + len(userPresets) + 1,
			Label:     presetName.Text,
			Timezones: activeTimezones,
		}

		userPresets = append(userPresets, preset)

		_ = t.userPresets.Set(userPresets)
		t.window.SetMainMenu(t.MakeMenu(t.app))
		t.preset.Set(preset.Id)
		w.Close()
	})

	okBtn.Importance = widget.HighImportance
	okBtn.Icon = theme.ConfirmIcon()

	cancelBtn := widget.NewButton("Cancel", func() {
		w.Close()
	})

	cancelBtn.Icon = theme.CancelIcon()

	w.SetContent(container.NewVBox(
		widget.NewLabel("Enter preset name"),
		presetName,
		errorLabel,
		container.NewCenter(container.NewHBox(container.NewMax(okBtn), cancelBtn)),
	))

	w.SetIcon(theme.ContentAddIcon())
	w.Resize(fyne.NewSize(480, 10))
	w.SetFixedSize(true)
	w.Show()
}
