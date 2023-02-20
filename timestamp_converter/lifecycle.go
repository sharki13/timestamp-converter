package timestamp_converter

import (
	"time"

	prefSync "com.sharki13/timestamp.converter/preferences"
	"com.sharki13/timestamp.converter/xbinding"
	"fyne.io/fyne/v2/data/binding"
)

// Sets up the preferences and loads them
// from the fyne preferences
// Should be called just before ShowAndRun
func (t *TimestampConverter) setupAndLoadPreferences() {
	err := t.preferences.AddString(prefSync.StringPreference{
		Key:      "format",
		Value:    t.format,
		Fallback: time.RFC3339,
	})

	if err != nil {
		panic(err)
	}

	err = t.preferences.AddString(prefSync.StringPreference{
		Key:      "theme",
		Value:    t.theme,
		Fallback: SystemTheme,
	})

	if err != nil {
		panic(err)
	}

	err = t.preferences.AddIntArray(prefSync.IntArrayPreference{
		Key:      "visibleTimezones",
		Value:    t.visibleTimezones,
		Fallback: []int{0},
	})

	if err != nil {
		panic(err)
	}

	savedTimezones, err := t.visibleTimezones.Get()
	for _, timezoneIndex := range savedTimezones {
		if visible, ok := t.timezonesVisibleState[timezoneIndex]; ok {
			visible.Set(true)
		}
	}

	if err != nil {
		panic(err)
	}

	// run background loop to watch for clipboard changes
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

				timestamp, err := praseStringToTime(cliboardContent)
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
}

func (t *TimestampConverter) initialize() {
	t.timezonesVisibleState = make(map[int]binding.Bool)
	t.visibleTimezones = xbinding.NewIntArray()
	t.timestamp = xbinding.NewTime()
	t.timestamp.Set(time.Now())
	t.format = binding.NewString()
	t.theme = binding.NewString()
	t.preferences = prefSync.NewPreferencesSynchronizer(t.app)
}
