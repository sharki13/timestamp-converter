package main

import (
	"time"

	"fyne.io/fyne/v2/data/binding"
)

// Minimal implementation of a time.Time binding
type BoundTime struct {
	value binding.Untyped
}

func NewBoundTime() BoundTime {
	return BoundTime{
		value: binding.NewUntyped(),
	}
}

func (t *BoundTime) Set(value time.Time) error {
	return t.value.Set(value)
}

func (t *BoundTime) Get() (time.Time, error) {
	value, err := t.value.Get()
	if err != nil {
		return time.Time{}, err
	}

	return value.(time.Time), nil
}

func (t *BoundTime) AddListener(listener binding.DataListener) {
	t.value.AddListener(listener)
}