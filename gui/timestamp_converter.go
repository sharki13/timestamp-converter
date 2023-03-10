package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	prefSync "github.com/sharki13/timestamp-converter/preferences"
	"github.com/sharki13/timestamp-converter/xbinding"

	// to include tzdata in the binary, not all OSes have it like windows
	_ "time/tzdata"
)

type TimestampConverter struct {
	timezonesVisibleState map[int]binding.Bool
	visibleTimezones      xbinding.IntArray
	timestamp             xbinding.Time
	format                binding.String
	watchClipboard        bool
	theme                 binding.String
	window                fyne.Window
	app                   fyne.App
	preferences           *prefSync.PreferencesSynchronizer
}

func NewTimestampConverter(app fyne.App) *TimestampConverter {
	ret := TimestampConverter{
		app:    app,
		window: app.NewWindow(TimestampConverterLabel),
	}

	ret.window.SetIcon(theme.HistoryIcon())

	ret.initialize()

	return &ret
}

// Should be called near the end of the function
// becasue it will block until the window is closed
func (t *TimestampConverter) ShowAndRun() {
	t.window.SetMainMenu(t.makeMenu())
	t.window.SetContent(t.makeContent())
	t.setupAndLoadPreferences()
	t.window.Resize(fyne.NewSize(600, 400))
	t.window.ShowAndRun()
}
