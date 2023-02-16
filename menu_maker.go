package main

import (
	"fmt"
	"net/url"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
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
		t.MakePresetMenu(app),
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

func (t* TimestampConverter) MakeSettingsMenu(app fyne.App) *fyne.Menu {
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

func (t* TimestampConverter) MakeThemeMenu(app fyne.App) *fyne.MenuItem {
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

	savedFormat := app.Preferences().String("format")

	for _, format := range SupportedFormats {
		format := format
		formatMenuItem := fyne.NewMenuItem(format.Label, func() {
			t.Format.Set(format.Format)

			for _, item := range formatMenu.Items {
				if item.Label != format.Label {
					item.Checked = false
				} else {
					item.Checked = true
				}
			}

			app.Preferences().SetString("format", format.Format)
		})
		formatMenu.Items = append(formatMenu.Items, formatMenuItem)

		if savedFormat == format.Format {
			formatMenuItem.Checked = true
			t.Format.Set(format.Format)
		}
	}

	if len(formatMenu.Items) == 0 {
		panic("no format found")
	}

	if savedFormat == "" {
		formatMenu.Items[0].Checked = true
		t.Format.Set(SupportedFormats[0].Format)
	}

	return formatMenu
}

func (t *TimestampConverter) MakePresetMenu(app fyne.App) *fyne.Menu {
	presetsMenu := fyne.NewMenu("Presets", make([]*fyne.MenuItem, 0)...)

	savedPreset := app.Preferences().Int("preset")

	for _, preset := range TimezonePresets {
		preset := preset

		presetsMenuItem := fyne.NewMenuItem(preset.Label, func() {
			for _, e := range t.TimezonesVisbleState {
				e.Set(false)
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
						if _, ok := t.TimezonesVisbleState[id]; !ok {
							continue
						}

						t.TimezonesVisbleState[id].Set(true)
					}
				}
			}

			t.SetStatus(fmt.Sprintf("Preset %s", preset.Label))
			app.Preferences().SetInt("preset", preset.Id)
		})

		if savedPreset == preset.Id {
			presetsMenuItem.Checked = true
			presetsMenuItem.Action()
		}

		presetsMenu.Items = append(presetsMenu.Items, presetsMenuItem)
	}

	if len(presetsMenu.Items) == 0 {
		panic("no presets found")
	}

	if savedPreset == 0 {
		presetsMenu.Items[0].Checked = true
		presetsMenu.Items[0].Action()
	}

	return presetsMenu
}