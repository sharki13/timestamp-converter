package main

import (
	"net/url"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
)

func (t *TimestampConverter) MakeMenu() *fyne.MainMenu {

	menus := make([]*fyne.Menu, 0)

	// Mac OS has a built in quit menu,
	// on other platforms Fyne will add Quit to first menu if it is not defined
	if runtime.GOOS != "darwin" {
		fileMenu := fyne.NewMenu("File", fyne.NewMenuItem("Quit", func() {
			t.app.Quit()
		}))

		menus = append(menus, fileMenu)
	}

	menus = append(menus,
		t.MakeFormatMenu(),
		t.MakeThemeMenu(),
		t.MakeInfoMenu(),
	)

	return fyne.NewMainMenu(menus...)
}

func (t *TimestampConverter) MakeInfoMenu() *fyne.Menu {
	about := fyne.NewMenuItem("GitHub page", func() {
		u, _ := url.Parse("https://github.com/sharki13/timestamp-converter")
		_ = t.app.OpenURL(u)
	})

	return fyne.NewMenu("Help", about)
}

func (t *TimestampConverter) MakeThemeMenu() *fyne.Menu {
	system := fyne.NewMenuItem("System", nil)
	light := fyne.NewMenuItem("Light", nil)
	dark := fyne.NewMenuItem("Dark", nil)

	if t.app.Settings().Theme() == theme.LightTheme() {
		light.Checked = true
	} else if t.app.Settings().Theme() == theme.DarkTheme() {
		dark.Checked = true
	} else {
		system.Checked = true
	}

	light.Action = func() {
		t.app.Settings().SetTheme(theme.LightTheme())
		light.Checked = true
		dark.Checked = false
		system.Checked = false
	}

	dark.Action = func() {
		t.app.Settings().SetTheme(theme.DarkTheme())
		light.Checked = false
		dark.Checked = true
		system.Checked = false
	}

	system.Action = func() {
		t.app.Settings().SetTheme(theme.DefaultTheme())
		light.Checked = false
		dark.Checked = false
		system.Checked = true
	}

	return fyne.NewMenu("Theme", system, fyne.NewMenuItemSeparator(), light, dark)
}

func (t *TimestampConverter) MakeFormatMenu() *fyne.Menu {

	formatMenu := fyne.NewMenu("Format", make([]*fyne.MenuItem, 0)...)

	for k, label := range FormatLabelMap {
		format := k
		label := label
		formatMenuItem := fyne.NewMenuItem(label, func() {
			t.format.Set(format)
		})

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
