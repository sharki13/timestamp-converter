package preferences

import (
	"reflect"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/test"
	"github.com/sharki13/timestamp-converter/xbinding"

	"testing"
)

type assert struct {
	t *testing.T
}

func (a *assert) True(value bool, message string) {
	if !value {
		a.t.Error(message)
	}
}

func (a *assert) False(value bool, message string) {
	if value {
		a.t.Error(message)
	}
}

func (a *assert) Equal(expected interface{}, actual interface{}, message string) {
	if !reflect.DeepEqual(expected, actual) {
		a.t.Errorf("%s: expected %v, got %v", message, expected, actual)
	}
}

func (a *assert) NotEqual(expected interface{}, actual interface{}, message string) {
	if reflect.DeepEqual(expected, actual) {
		a.t.Errorf("%s: expected %v, got %v", message, expected, actual)
	}
}

func (a *assert) Nil(value interface{}, message string) {
	if value != nil {
		a.t.Errorf("%s: expected nil, got %v", message, value)
	}
}

func (a *assert) NotNil(value interface{}, message string) {
	if value == nil {
		a.t.Errorf("%s: expected not nil, got nil", message)
	}
}

func (a *assert) Error(err error, message string) {
	if err == nil {
		a.t.Errorf("%s: expected error, got nil", message)
	}
}

func (a *assert) NoError(err error, message string) {
	if err != nil {
		a.t.Errorf("%s: expected no error, got %v", message, err)
	}
}

func TestPreferences_Bool_Empty(t *testing.T) {
	assert := assert{t}
	testApp := test.NewApp()

	prefSync := NewPreferencesSynchronizer(testApp)

	testBoolBinding := binding.NewBool()

	err := prefSync.AddBool(BoolPreference{
		Key:      "testBool",
		Value:    testBoolBinding,
		Fallback: false,
	})

	assert.NoError(err, "AddBool should not return an error")

	value, err := testBoolBinding.Get()
	assert.NoError(err, "Get should not return an error")

	assert.False(value, "Value should be false")

	testBoolBinding.Set(true)

	value, err = testBoolBinding.Get()
	assert.NoError(err, "Get should not return an error")
	assert.True(value, "Value should be true")

	err = prefSync.AddBool(BoolPreference{
		Key:      "testBool",
		Value:    binding.NewBool(),
		Fallback: false,
	})

	assert.Error(err, "AddBool should return an error")
}

func TestPreferences_Bool_NonEmpty(t *testing.T) {
	assert := assert{t}
	testApp := test.NewApp()

	testApp.Preferences().SetBool("testBool", true)

	prefSync := NewPreferencesSynchronizer(testApp)

	testBoolBinding := binding.NewBool()

	err := prefSync.AddBool(BoolPreference{
		Key:      "testBool",
		Value:    testBoolBinding,
		Fallback: false,
	})

	assert.NoError(err, "AddBool should not return an error")

	value, err := testBoolBinding.Get()
	assert.NoError(err, "Get should not return an error")

	assert.True(value, "Value should be ture")

	testBoolBinding.Set(false)

	value, err = testBoolBinding.Get()
	assert.NoError(err, "Get should not return an error")
	assert.False(value, "Value should be false")

	err = prefSync.AddBool(BoolPreference{
		Key:      "testBool",
		Value:    binding.NewBool(),
		Fallback: false,
	})

	assert.Error(err, "AddBool should return an error")
}

func TestPreferences_Int_Empty(t *testing.T) {
	assert := assert{t}
	testApp := test.NewApp()

	prefSync := NewPreferencesSynchronizer(testApp)

	testIntBinding := binding.NewInt()

	err := prefSync.AddInt(IntPreference{
		Key:      "testInt",
		Value:    testIntBinding,
		Fallback: 20,
	})

	assert.NoError(err, "AddInt should not return an error")

	value, err := testIntBinding.Get()
	assert.NoError(err, "Get should not return an error")

	assert.Equal(20, value, "Value should be 20")

	testIntBinding.Set(1)

	value, err = testIntBinding.Get()
	assert.NoError(err, "Get should not return an error")
	assert.Equal(1, value, "Value should be 1")

	err = prefSync.AddInt(IntPreference{
		Key:      "testInt",
		Value:    binding.NewInt(),
		Fallback: 0,
	})

	assert.Error(err, "AddInt should return an error")
}

func TestPreferences_Int_NonEmpty(t *testing.T) {
	assert := assert{t}
	testApp := test.NewApp()
	testApp.Preferences().SetInt("testInt", 16)

	prefSync := NewPreferencesSynchronizer(testApp)

	testIntBinding := binding.NewInt()

	err := prefSync.AddInt(IntPreference{
		Key:      "testInt",
		Value:    testIntBinding,
		Fallback: 20,
	})

	assert.NoError(err, "AddInt should not return an error")

	value, err := testIntBinding.Get()
	assert.NoError(err, "Get should not return an error")

	assert.Equal(16, value, "Value should be 16")

	testIntBinding.Set(1)

	value, err = testIntBinding.Get()
	assert.NoError(err, "Get should not return an error")
	assert.Equal(1, value, "Value should be 1")

	err = prefSync.AddInt(IntPreference{
		Key:      "testInt",
		Value:    binding.NewInt(),
		Fallback: 0,
	})

	assert.Error(err, "AddInt should return an error")
}

func TestPreferences_IntArray_Empty(t *testing.T) {
	assert := assert{t}
	testApp := test.NewApp()

	prefSync := NewPreferencesSynchronizer(testApp)

	testIntArrayBinding := xbinding.NewIntArray()

	err := prefSync.AddIntArray(IntArrayPreference{
		Key:      "testIntArray",
		Value:    testIntArrayBinding,
		Fallback: []int{1, 2, 3},
	})

	assert.NoError(err, "AddIntArray should not return an error")

	valueIntArray, err := testIntArrayBinding.Get()
	assert.NoError(err, "Get should not return an error")

	assert.Equal([]int{1, 2, 3}, valueIntArray, "Value should be [1, 2, 3]")

	err = testIntArrayBinding.Set([]int{4, 5, 6})
	assert.NoError(err, "Set should not return an error")

	valueIntArray, err = testIntArrayBinding.Get()
	assert.NoError(err, "Get should not return an error")
	assert.Equal([]int{4, 5, 6}, valueIntArray, "Value should be [4, 5, 6]")

}

func TestPreferences_IntArray_NonEmpty(t *testing.T) {
	assert := assert{t}
	testApp := test.NewApp()
	testApp.Preferences().SetString("testIntArray", "[3, 5, 8]")

	prefSync := NewPreferencesSynchronizer(testApp)

	testIntArrayBinding := xbinding.NewIntArray()

	err := prefSync.AddIntArray(IntArrayPreference{
		Key:      "testIntArray",
		Value:    testIntArrayBinding,
		Fallback: []int{1, 2, 3},
	})

	assert.NoError(err, "AddIntArray should not return an error")

	valueIntArray, err := testIntArrayBinding.Get()
	assert.NoError(err, "Get should not return an error")

	assert.Equal([]int{3, 5, 8}, valueIntArray, "Value should be [3, 5, 8]")

	err = testIntArrayBinding.Set([]int{4, 5, 6})
	assert.NoError(err, "Set should not return an error")

	valueIntArray, err = testIntArrayBinding.Get()
	assert.NoError(err, "Get should not return an error")
	assert.Equal([]int{4, 5, 6}, valueIntArray, "Value should be [4, 5, 6]")

}
