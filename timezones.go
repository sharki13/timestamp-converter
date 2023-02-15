package main

type TimezoneType int

const (
	Local TimezoneType = iota
	WithLocation
	Unix
)

type TimezoneDefinition struct {
	Id               int
	LocationAsString string
	Label            string
	Type             TimezoneType
	Presets          []int
}

type TimezonePreset struct {
	Id    int
	Label string
}

var TimezonePresets = []TimezonePreset{
	{
		Id:    0,
		Label: "All",
	},
	{
		Id:    1,
		Label: "Developer",
	},
	{
		Id:    2,
		Label: "US",
	},
	{
		Id:    3,
		Label: "Europe",
	},
	{
		Id:    4,
		Label: "US & Europe",
	},
}

var Timezones = []TimezoneDefinition{
	{
		Id:               0,
		LocationAsString: "Local",
		Label:            "Local",
		Type:             Local,
		Presets:          []int{0, 1, 2, 3, 4},
	},
	{
		Id:               1,
		LocationAsString: "UTC",
		Label:            "Unix",
		Type:             Unix,
		Presets:          []int{0},
	},
	{
		Id:               2,
		LocationAsString: "Pacific/Honolulu",
		Label:            "HST (Hawaii), US",
		Type:             WithLocation,
		Presets:          []int{0, 1, 2, 4},
	},
	{
		Id:               3,
		LocationAsString: "America/Anchorage",
		Label:            "AKST/AKDT (Alaska), US",
		Type:             WithLocation,
		Presets:          []int{0, 1, 2, 4},
	},
	{
		Id:               4,
		LocationAsString: "America/Los_Angeles",
		Label:            "PST/PDT (Pacific), US",
		Type:             WithLocation,
		Presets:          []int{0, 1, 2, 4},
	},
	{
		Id:               5,
		LocationAsString: "America/Phoenix",
		Label:            "MST (Mountain), US",
		Type:             WithLocation,
		Presets:          []int{0, 1, 2, 4},
	},
	{
		Id:               6,
		LocationAsString: "America/Chicago",
		Label:            "CST/CDT (Central), US",
		Type:             WithLocation,
		Presets:          []int{0, 1, 2, 4},
	},
	{
		Id:               7,
		LocationAsString: "America/New_York",
		Label:            "EST/EDT (Eastern), US",
		Type:             WithLocation,
		Presets:          []int{0, 1, 2, 4},
	},
	{
		Id:               8,
		LocationAsString: "America/Grenada",
		Label:            "AST (Atlantic), GD",
		Type:             WithLocation,
		Presets:          []int{0},
	},
	{
		Id:               9,
		LocationAsString: "Europe/London",
		Label:            "GMT/BST (Greenwich), UK",
		Type:             WithLocation,
		Presets:          []int{0, 3, 4},
	},
	{
		Id:               10,
		LocationAsString: "Europe/Paris",
		Label:            "CET/CEST (Central Europe), France",
		Type:             WithLocation,
		Presets:          []int{0, 1, 3, 4},
	},
	{
		Id:               11,
		LocationAsString: "Europe/Helsinki",
		Label:            "EET/EEST (Eastern Europe), Finland",
		Type:             WithLocation,
		Presets:          []int{0, 3, 4},
	},
	{
		Id:               12,
		LocationAsString: "Europe/Moscow",
		Label:            "MSK (Moscow), Russia",
		Type:             WithLocation,
		Presets:          []int{0, 3, 4},
	},
	{
		Id:               13,
		LocationAsString: "Asia/Kolkata",
		Label:            "IST (India), India",
		Type:             WithLocation,
		Presets:          []int{0},
	},
	{
		Id:               14,
		LocationAsString: "Asia/Chongqing",
		Label:            "CST (China), China",
		Type:             WithLocation,
		Presets:          []int{0},
	},
	{
		Id:               15,
		LocationAsString: "Australia/Sydney",
		Label:            "AEST/AEDT (Australia), Australia",
		Type:             WithLocation,
		Presets:          []int{0},
	},
	{
		Id:               16,
		LocationAsString: "UTC",
		Label:            "UTC",
		Type:             WithLocation,
		Presets:          []int{0, 1, 2, 3, 4},
	},
}
