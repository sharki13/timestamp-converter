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
	value binding.Untyped
}

func NewPresets() Presets {
	return Presets{
		value: binding.NewUntyped(),
	}
}

func (e *Presets) Set(value timezone.Presets) error {
	return e.value.Set(value)
}

func (e *Presets) Get() (timezone.Presets, error) {
	value, err := e.value.Get()
	if err != nil {
		return timezone.Presets{}, err
	}

	return value.(timezone.Presets), nil
}

func (e *Presets) AddListener(listener binding.DataListener) {
	e.value.AddListener(listener)
}
