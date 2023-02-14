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
}

var Timezones = []TimezoneDefinition{
	{
		Id:               0,
		LocationAsString: "Local",
		Label:            "Local",
		Type:             Local,
	},
	{
		Id:               1,
		LocationAsString: "UTC",
		Label:            "Unix",
		Type:             Unix,
	},
	{
		Id:               2,
		LocationAsString: "Pacific/Honolulu",
		Label:            "HST (Hawaii), US",
		Type:             WithLocation,
	},
	{
		Id:               3,
		LocationAsString: "America/Anchorage",
		Label:            "AKST/AKDT (Alaska), US",
		Type:             WithLocation,
	},
	{
		Id:               4,
		LocationAsString: "America/Los_Angeles",
		Label:            "PST/PDT (Pacific), US",
		Type:             WithLocation,
	},
	{
		Id:               5,
		LocationAsString: "America/Phoenix",
		Label:            "MST (Mountain), US",
		Type:             WithLocation,
	},
	{
		Id:               6,
		LocationAsString: "America/Chicago",
		Label:            "CST/CDT (Central), US",
		Type:             WithLocation,
	},
	{
		Id:               7,
		LocationAsString: "America/New_York",
		Label:            "EST/EDT (Eastern), US",
		Type:             WithLocation,
	},
	{
		Id:               8,
		LocationAsString: "America/Grenada",
		Label:            "AST (Atlantic), GD",
		Type:             WithLocation,
	},
	{
		Id:               9,
		LocationAsString: "Europe/London",
		Label:            "GMT/BST (Greenwich), UK",
		Type:             WithLocation,
	},
	{
		Id:               10,
		LocationAsString: "Europe/Paris",
		Label:            "CET/CEST (Central Europe), France",
		Type:             WithLocation,
	},
	{
		Id:               11,
		LocationAsString: "Europe/Helsinki",
		Label:            "EET/EEST (Eastern Europe), Finland",
		Type:             WithLocation,
	},
	{
		Id:               12,
		LocationAsString: "Europe/Moscow",
		Label:            "MSK (Moscow), Russia",
		Type:             WithLocation,
	},
	{
		Id:               13,
		LocationAsString: "Asia/Kolkata",
		Label:            "IST (India), India",
		Type:             WithLocation,
	},
	{
		Id:               14,
		LocationAsString: "Asia/Chongqing",
		Label:            "CST (China), China",
		Type:             WithLocation,
	},
	{
		Id:               15,
		LocationAsString: "Australia/Sydney",
		Label:            "AEST/AEDT (Australia), Australia",
		Type:             WithLocation,
	},
	{
		Id:               16,
		LocationAsString: "UTC",
		Label:            "UTC",
		Type:             WithLocation,
	},
}
