package xbinding

import (
	"time"

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
