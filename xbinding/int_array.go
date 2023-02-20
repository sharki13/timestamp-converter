package xbinding

import "fyne.io/fyne/v2/data/binding"

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
