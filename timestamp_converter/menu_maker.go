package timestamp_converter

import (
	"net/url"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
)

func (t *TimestampConverter) makeMenu() *fyne.MainMenu {

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
		t.makeFormatMenu(),
		t.makeThemeMenu(),
		t.makeInfoMenu(),
	)

	return fyne.NewMainMenu(menus...)
}

func (t *TimestampConverter) makeInfoMenu() *fyne.Menu {
	about := fyne.NewMenuItem("GitHub page", func() {
		u, _ := url.Parse("https://github.com/sharki13/timestamp-converter")
		_ = t.app.OpenURL(u)
	})

	return fyne.NewMenu("Help", about)
}

func (t *TimestampConverter) makeThemeMenu() *fyne.Menu {
	system := fyne.NewMenuItem("System", nil)
	light := fyne.NewMenuItem("Light", nil)
	dark := fyne.NewMenuItem("Dark", nil)

	system.Checked = true

	light.Action = func() {
		t.theme.Set("light")
	}

	dark.Action = func() {
		t.theme.Set("dark")
	}

	system.Action = func() {
		t.theme.Set("system")
	}

	t.theme.AddListener(binding.NewDataListener(func() {
		themeVariant, err := t.theme.Get()
		if err != nil {
			panic(err)
		}

		switch themeVariant {
		case "light":
			t.app.Settings().SetTheme(&myTheme{variant: "light"})
			light.Checked = true
			dark.Checked = false
			system.Checked = false
		case "dark":
			t.app.Settings().SetTheme(&myTheme{variant: "dark"})
			light.Checked = false
			dark.Checked = true
			system.Checked = false
		default:
			if t.app.Settings().ThemeVariant() == theme.VariantLight {
				t.app.Settings().SetTheme(&myTheme{variant: "light"})
			} else {
				t.app.Settings().SetTheme(&myTheme{variant: "dark"})
			}
			light.Checked = false
			dark.Checked = false
			system.Checked = true
		}
	}))

	return fyne.NewMenu("Theme", system, fyne.NewMenuItemSeparator(), light, dark)
}

func (t *TimestampConverter) makeFormatMenu() *fyne.Menu {

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
