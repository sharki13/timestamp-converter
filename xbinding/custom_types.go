package xbinding

import (
	"time"

	"com.sharki13/timestamp.converter/timezone"
	"fyne.io/fyne/v2/data/binding"
)

// Minimal implementation of a time.Time binding
type Time struct {
	value binding.Untyped
}

func NewTime() Time {
	return Time{
		value: binding.NewUntyped(),
	}
}

func (t *Time) Set(value time.Time) error {
	return t.value.Set(value)
}

func (t *Time) Get() (time.Time, error) {
	value, err := t.value.Get()
	if err != nil {
		return time.Time{}, err
	}

	return value.(time.Time), nil
}

func (t *Time) AddListener(listener binding.DataListener) {
	t.value.AddListener(listener)
}

type Presets struct {
	value binding.UntypedList
}

func NewPresets() Presets {
	return Presets{
		value: binding.NewUntypedList(),
	}
}

func (e *Presets) Set(value []timezone.Preset) error {
	toSet := make([]interface{}, len(value))
	for i, v := range value {
		toSet[i] = v
	}

	return e.value.Set(toSet)
}

func (e *Presets) Get() ([]timezone.Preset, error) {
	value, err := e.value.Get()
	if err != nil {
		return []timezone.Preset{}, err
	}

	if value == nil {
		return []timezone.Preset{}, nil
	}

	toReturn := make([]timezone.Preset, len(value))
	for i, v := range value {
		toReturn[i] = v.(timezone.Preset)
	}

	return toReturn, nil
}

func (e *Presets) AddListener(listener binding.DataListener) {
	e.value.AddListener(listener)
}
