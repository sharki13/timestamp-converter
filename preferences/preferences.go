package preferences

import (
	"encoding/json"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"github.com/sharki13/timestamp-converter/xbinding"
)

type Keyed interface {
	GetKey() string
}

// Preference that is stored as a string
// key: the key of the preference, has to be unique across all preferences
type StringPreference struct {
	Key      string
	Value    binding.String
	Fallback string
}

func (s StringPreference) GetKey() string {
	return s.Key
}

// Preference that is stored as an int
// key: the key of the preference, has to be unique across all preferences
type IntPreference struct {
	Key      string
	Value    binding.Int
	Fallback int
}

func (i IntPreference) GetKey() string {
	return i.Key
}

type IntArrayPreference struct {
	Key      string
	Value    xbinding.IntArray
	Fallback []int
}

func (i IntArrayPreference) GetKey() string {
	return i.Key
}

// Preference that is stored as a boolean
// key: the key of the preference, has to be unique across all preferences
type BoolPreference struct {
	Key      string
	Value    binding.Bool
	Fallback bool
}

func (b BoolPreference) GetKey() string {
	return b.Key
}

// PreferencesSynchronizer is used to sync preferences
// between bindings and the fyne preferences
type PreferencesSynchronizer struct {
	stringPreferences   []StringPreference
	intPreferences      []IntPreference
	boolPreferences     []BoolPreference
	intArrayPreferences []IntArrayPreference
	app                 fyne.App
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
	pref.intArrayPreferences = make([]IntArrayPreference, 0)

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

// Adds a new int array preference to the synchronizer
// and sets the value to the current value of the preference
// or the fallback value if the preference is not set
func (p *PreferencesSynchronizer) AddIntArray(e IntArrayPreference) error {
	if p.isKeyExisting(e.Key) {
		return fmt.Errorf("key %s is already in use", e.Key)
	}

	serialized := p.app.Preferences().StringWithFallback(e.Key, "[]")
	deserialized := make([]int, 0)

	if serialized != "[]" {
		if err := json.Unmarshal([]byte(serialized), &deserialized); err != nil {
			return err
		}
	} else {
		deserialized = e.Fallback
	}

	e.Value.Set(deserialized)

	p.intArrayPreferences = append(p.intArrayPreferences, e)

	e.Value.AddListener(binding.NewDataListener(func() {
		v, err := e.Value.Get()
		if err != nil {
			panic(err)
		}

		serialized, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		p.app.Preferences().SetString(e.Key, string(serialized))
	}))

	return nil
}

func isKeyExistInCollection[T Keyed](key string, collection []T) bool {
	for _, pref := range collection {
		if pref.GetKey() == key {
			return true
		}
	}

	return false
}

func (p *PreferencesSynchronizer) isKeyExisting(key string) bool {
	if exist := isKeyExistInCollection(key, p.stringPreferences); exist {
		return true
	}

	if exist := isKeyExistInCollection(key, p.intPreferences); exist {
		return true
	}

	if exist := isKeyExistInCollection(key, p.boolPreferences); exist {
		return true
	}

	if exist := isKeyExistInCollection(key, p.intArrayPreferences); exist {
		return true
	}

	return false
}
