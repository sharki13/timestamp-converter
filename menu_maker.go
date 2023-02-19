package main

import (
	"net/url"
	"runtime"

	"com.sharki13/timestamp.converter/timezone"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
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

	menus = append(menus,
		t.MakePresetMenu(),
		t.MakeFormatMenu(app),
		t.MakeThemeMenu(app),
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

func (t *TimestampConverter) MakeThemeMenu(app fyne.App) *fyne.Menu {
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

	return fyne.NewMenu("Theme", system, fyne.NewMenuItemSeparator(), light, dark)
}

func (t *TimestampConverter) MakeFormatMenu(app fyne.App) *fyne.Menu {

	formatMenu := fyne.NewMenu("Format", make([]*fyne.MenuItem, 0)...)
	currentFormat, _ := t.format.Get()

	for k, label := range FormatLabelMap {
		format := k
		label := label
		formatMenuItem := fyne.NewMenuItem(label, func() {
			t.format.Set(format)
		})

		if format == currentFormat {
			formatMenuItem.Checked = true
		} else {
			formatMenuItem.Checked = false
		}

		formatMenu.Items = append(formatMenu.Items, formatMenuItem)
	}

	t.format.AddListener(binding.NewDataListener(func() {
		currentFormat, err := t.format.Get()
		if err != nil {
			panic(err)
		}

		label := FormatLabelMap[currentFormat]

		for _, item := range formatMenu.Items {
			if item.Label == label {
				item.Checked = true
			} else {
				item.Checked = false
			}
		}
	}))

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

	presetMenuItem.Checked = true

	return presetMenuItem
}

func (t *TimestampConverter) MakePresetMenu() *fyne.Menu {
	items := make([]*fyne.MenuItem, 0)

	for _, preset := range timezone.DefaultPresets {
		preset := preset
		items = append(items, makePresetMenuItem(preset.Label, preset.Id, t.preset))
	}

	return fyne.NewMenu("Presets", items...)
}
