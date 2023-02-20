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

type IntArray struct {
	value binding.UntypedList
}

func NewIntArray() IntArray {
	return IntArray{
		value: binding.NewUntypedList(),
	}
}

func (t *IntArray) Set(value []int) error {
	interfaces := make([]interface{}, len(value))
	for i, v := range value {
		interfaces[i] = v
	}

	return t.value.Set(interfaces)
}

func (t *IntArray) Get() ([]int, error) {
	values, err := t.value.Get()
	if err != nil {
		return nil, err
	}

	// make from values a []int
	ret := make([]int, len(values))
	for i, v := range values {
		ret[i] = v.(int)
	}

	return ret, nil
}

func (t *IntArray) AddListener(listener binding.DataListener) {
	t.value.AddListener(listener)
}
