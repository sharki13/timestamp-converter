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

const (
	All int = iota
	Developer
	US
	Europe
	USEurope
)

var TimezonePresets = []TimezonePreset{
	{
		Id:    All,
		Label: "All",
	},
	{
		Id:    Developer,
		Label: "Developer",
	},
	{
		Id:    US,
		Label: "US",
	},
	{
		Id:    Europe,
		Label: "Europe",
	},
	{
		Id:    USEurope,
		Label: "US & Europe",
	},
}

var Timezones = []TimezoneDefinition{
	{
		Id:               0,
		LocationAsString: "Local",
		Label:            "Local",
		Type:             Local,
		Presets:          []int{All, Developer, US, Europe, USEurope},
	},
	{
		Id:               1,
		LocationAsString: "UTC",
		Label:            "Unix",
		Type:             Unix,
		Presets:          []int{All, Developer},
	},
	{
		Id:               2,
		LocationAsString: "Pacific/Honolulu",
		Label:            "HST (Hawaii), US",
		Type:             WithLocation,
		Presets:          []int{All, US, USEurope},
	},
	{
		Id:               3,
		LocationAsString: "America/Anchorage",
		Label:            "AKST/AKDT (Alaska), US",
		Type:             WithLocation,
		Presets:          []int{All, US, USEurope},
	},
	{
		Id:               4,
		LocationAsString: "America/Los_Angeles",
		Label:            "PST/PDT (Pacific), US",
		Type:             WithLocation,
		Presets:          []int{All, Developer, US, USEurope},
	},
	{
		Id:               5,
		LocationAsString: "America/Phoenix",
		Label:            "MST (Mountain), US",
		Type:             WithLocation,
		Presets:          []int{All, US, USEurope},
	},
	{
		Id:               6,
		LocationAsString: "America/Chicago",
		Label:            "CST/CDT (Central), US",
		Type:             WithLocation,
		Presets:          []int{All, Developer, US, USEurope},
	},
	{
		Id:               7,
		LocationAsString: "America/New_York",
		Label:            "EST/EDT (Eastern), US",
		Type:             WithLocation,
		Presets:          []int{0, All, US, USEurope},
	},
	{
		Id:               8,
		LocationAsString: "America/Grenada",
		Label:            "AST (Atlantic), GD",
		Type:             WithLocation,
		Presets:          []int{All},
	},
	{
		Id:               9,
		LocationAsString: "Europe/London",
		Label:            "GMT/BST (Greenwich), UK",
		Type:             WithLocation,
		Presets:          []int{All, Europe, USEurope},
	},
	{
		Id:               10,
		LocationAsString: "Europe/Paris",
		Label:            "CET/CEST (Central Europe), France",
		Type:             WithLocation,
		Presets:          []int{All, Developer, Europe, USEurope},
	},
	{
		Id:               11,
		LocationAsString: "Europe/Helsinki",
		Label:            "EET/EEST (Eastern Europe), Finland",
		Type:             WithLocation,
		Presets:          []int{All, Europe, USEurope},
	},
	{
		Id:               12,
		LocationAsString: "Europe/Moscow",
		Label:            "MSK (Moscow), Russia",
		Type:             WithLocation,
		Presets:          []int{All, Europe, USEurope},
	},
	{
		Id:               13,
		LocationAsString: "Asia/Kolkata",
		Label:            "IST (India), India",
		Type:             WithLocation,
		Presets:          []int{All},
	},
	{
		Id:               14,
		LocationAsString: "Asia/Chongqing",
		Label:            "CST (China), China",
		Type:             WithLocation,
		Presets:          []int{All},
	},
	{
		Id:               15,
		LocationAsString: "Australia/Sydney",
		Label:            "AEST/AEDT (Australia), Australia",
		Type:             WithLocation,
		Presets:          []int{All},
	},
	{
		Id:               16,
		LocationAsString: "UTC",
		Label:            "UTC",
		Type:             WithLocation,
		Presets:          []int{All, Developer, US, Europe, USEurope},
	},
}
