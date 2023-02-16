package main

import (
	"encoding/json"
	"fmt"

	"fyne.io/fyne/v2"
)

func (t* TimestampConverter) FetchUserPresents(app fyne.App) []TimezonePreset {
	ret := make([]TimezonePreset, 0)

	presets := app.Preferences().StringWithFallback("presets", "[]")

	err := json.Unmarshal([]byte(presets), &ret)
	if err != nil {
		t.SetStatus(fmt.Sprintf("Error loading presets: %s", err))
		return nil
	}

	return ret
}

func (t* TimestampConverter) SaveUserPresents(app fyne.App, presets []TimezonePreset) {
	presetsJson, err := json.Marshal(presets)
	if err != nil {
		t.SetStatus(fmt.Sprintf("Error saving presets: %s", err))
		return
	}

	app.Preferences().SetString("presets", string(presetsJson))
}

