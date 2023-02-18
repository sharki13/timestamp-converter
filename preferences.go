package main

import (
	"fmt"

	"com.sharki13/timestamp.converter/timezone"
	"com.sharki13/timestamp.converter/xbinding"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

// Preference that is stored as a string
// key: the key of the preference, has to be unique across all preferences
type StringPreference struct {
	Key      string
	Value    binding.String
	Fallback string
}

// Preference that is stored as an int
// key: the key of the preference, has to be unique across all preferences
type IntPreference struct {
	Key      string
	Value    binding.Int
	Fallback int
}

// Preference that is stored as a boolean
// key: the key of the preference, has to be unique across all preferences
type BoolPreference struct {
	Key      string
	Value    binding.Bool
	Fallback bool
}

// Preference that is stored as a list of presets
// key: the key of the preference, has to be unique across all preferences
type PresetsPreference struct {
	Key      string
	Value    xbinding.Presets
	Fallback []timezone.Preset
}

// PreferencesSynchronizer is used to sync preferences
// between bindings and the fyne preferences
type PreferencesSynchronizer struct {
	stringPreferences []StringPreference
	intPreferences    []IntPreference
	boolPreferences   []BoolPreference
	userPresets       PresetsPreference
	app               fyne.App
}

// Creates a new preferences sync
// that can be used to sync preferences with the fyne preferences
// Remark: all bindings have to be initialized before calling this function
func NewPreferencesSynchronizer(app fyne.App) *PreferencesSynchronizer {
	pref := PreferencesSynchronizer{
		app: app,
	}

	pref.stringPreferences = make([]StringPreference, 0)
	pref.intPreferences = make([]IntPreference, 0)
	pref.boolPreferences = make([]BoolPreference, 0)

	return &pref
}

// Adds a new string preference to the synchronizer
// and sets the value to the current value of the preference
// or the fallback value if the preference is not set
func (p *PreferencesSynchronizer) AddString(e StringPreference) error {
	if p.isKeyExisting(e.Key) {
		return fmt.Errorf("key %s is already in use", e.Key)
	}

	e.Value.Set(p.app.Preferences().StringWithFallback(e.Key, e.Fallback))

	p.stringPreferences = append(p.stringPreferences, e)

	e.Value.AddListener(binding.NewDataListener(func() {
		v, err := e.Value.Get()
		if err != nil {
			panic(err)
		}

		p.app.Preferences().SetString(e.Key, v)
	}))

	return nil
}

// Adds a new int preference to the synchronizer
// and sets the value to the current value of the preference
// or the fallback value if the preference is not set
func (p *PreferencesSynchronizer) AddInt(e IntPreference) error {
	if p.isKeyExisting(e.Key) {
		return fmt.Errorf("key %s is already in use", e.Key)
	}

	e.Value.Set(p.app.Preferences().IntWithFallback(e.Key, e.Fallback))

	p.intPreferences = append(p.intPreferences, e)

	e.Value.AddListener(binding.NewDataListener(func() {
		v, err := e.Value.Get()
		if err != nil {
			panic(err)
		}

		p.app.Preferences().SetInt(e.Key, v)
	}))

	return nil
}

// Adds a new bool preference to the synchronizer
// and sets the value to the current value of the preference
// or the fallback value if the preference is not set
func (p *PreferencesSynchronizer) AddBool(e BoolPreference) error {
	if p.isKeyExisting(e.Key) {
		return fmt.Errorf("key %s is already in use", e.Key)
	}

	e.Value.Set(p.app.Preferences().BoolWithFallback(e.Key, e.Fallback))

	p.boolPreferences = append(p.boolPreferences, e)

	e.Value.AddListener(binding.NewDataListener(func() {
		v, err := e.Value.Get()
		if err != nil {
			panic(err)
		}

		p.app.Preferences().SetBool(e.Key, v)
	}))

	return nil
}

// Adds a new presets preference to the synchronizer
// and sets the value to the current value of the preference
// or the fallback value if the preference is not set
func (p *PreferencesSynchronizer) AddPresets(e PresetsPreference) error {
	if p.isKeyExisting(e.Key) {
		return fmt.Errorf("key %s is already in use", e.Key)
	}

	storageValue := p.app.Preferences().StringWithFallback(e.Key, "[]")

	if storageValue == "[]" {
		e.Value.Set(e.Fallback)
	} else {
		presets, err := timezone.DeserializePresets(storageValue)
		if err != nil {
			panic(err)
		}

		e.Value.Set(presets)
	}

	e.Value.AddListener(binding.NewDataListener(func() {
		v, err := e.Value.Get()
		if err != nil {
			panic(err)
		}

		serialized, err := timezone.SerializePresets(v)
		if err != nil {
			panic(err)
		}

		p.app.Preferences().SetString(e.Key, serialized)
	}))

	return nil
}

func (p *PreferencesSynchronizer) isKeyExisting(key string) bool {
	for _, pref := range p.stringPreferences {
		if pref.Key == key {
			return true
		}
	}

	for _, pref := range p.intPreferences {
		if pref.Key == key {
			return true
		}
	}

	for _, pref := range p.boolPreferences {
		if pref.Key == key {
			return true
		}
	}

	if p.userPresets.Key == key {
		return true
	}

	return false
}
