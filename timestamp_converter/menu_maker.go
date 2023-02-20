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
		fileMenu := fyne.NewMenu(FileLabel, fyne.NewMenuItem(QuitLabel, func() {
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
	about := fyne.NewMenuItem(GitHubPageLabel, func() {
		u, _ := url.Parse(ProjectPageURL)
		_ = t.app.OpenURL(u)
	})

	return fyne.NewMenu(HelpLabel, about)
}

func (t *TimestampConverter) makeThemeMenu() *fyne.Menu {
	system := fyne.NewMenuItem(SystemLabel, nil)
	light := fyne.NewMenuItem(LightLabel, nil)
	dark := fyne.NewMenuItem(DarkLabel, nil)

	light.Action = func() {
		t.theme.Set(LightTheme)
	}

	dark.Action = func() {
		t.theme.Set(DarkTheme)
	}

	system.Action = func() {
		t.theme.Set(SystemTheme)
	}

	t.theme.AddListener(binding.NewDataListener(func() {
		themeVariant, err := t.theme.Get()
		if err != nil {
			panic(err)
		}

		switch themeVariant {
		case LightTheme:
			t.app.Settings().SetTheme(&myTheme{variant: LightTheme})
			light.Checked = true
			dark.Checked = false
			system.Checked = false
		case DarkTheme:
			t.app.Settings().SetTheme(&myTheme{variant: DarkTheme})
			light.Checked = false
			dark.Checked = true
			system.Checked = false
		default:
			if t.app.Settings().ThemeVariant() == theme.VariantLight {
				t.app.Settings().SetTheme(&myTheme{variant: LightTheme})
			} else {
				t.app.Settings().SetTheme(&myTheme{variant: DarkTheme})
			}
			light.Checked = false
			dark.Checked = false
			system.Checked = true
		}
	}))

	return fyne.NewMenu(ThemeLabel, system, fyne.NewMenuItemSeparator(), light, dark)
}

func (t *TimestampConverter) makeFormatMenu() *fyne.Menu {
	formatMenu := fyne.NewMenu(FormatLabel, make([]*fyne.MenuItem, 0)...)

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
