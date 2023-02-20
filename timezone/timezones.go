package timezone

import (
	"strconv"
	"time"
)

type TimezoneType int

const (
	LocalTimezoneType TimezoneType = iota
	WithLocationTimzoneType
	UnixTimezoneType
	FixedOffsetTimezoneType
)

type TimezoneDefinition struct {
	Id               int
	LocationAsString string
	Label            string
	Offset           int
	Type             TimezoneType
}

func (td TimezoneDefinition) StringTime(t time.Time, format string) string {
	if td.Type == UnixTimezoneType {
		return strconv.FormatInt(t.Unix(), 10)
	} else if td.Type == FixedOffsetTimezoneType {
		return t.In(time.FixedZone(td.Label, td.Offset)).Format(format)
	} else {
		return t.In(td.Location()).Format(format)
	}
}

func (td TimezoneDefinition) Location() *time.Location {
	loc, err := time.LoadLocation(td.LocationAsString)
	if err != nil {
		panic(err)
	}

	return loc
}

type TimeonzeId int

const (
	Local int = iota
	Unix
	HST_Pacific_Honolulu_US
	AKST_AKDT_Alaska_US
	PST_PDT_Pacific_US
	MST_Mountain_US
	CST_CDT_Central_US
	EST_EDT_Eastern_US
	AST_Atlantic_GD
	GMT_BST_Greenwich_UK
	WET_WEST_Western_Europe
	CET_CEST_Central_Europe_France
	EET_EEST_Eastern_Europe_Finland
	MSK_Moscow_Russia
	IST_India_India
	CST_China_China
	AEST_AEDT_Australia_Australia
	UTC
	LastNamedId
)

var Timezones = []TimezoneDefinition{
	{
		Id:               Local,
		LocationAsString: "Local",
		Label:            "Local",
		Type:             LocalTimezoneType,
	},
	{
		Id:               Unix,
		LocationAsString: "UTC",
		Label:            "Unix",
		Type:             UnixTimezoneType,
	},
	{
		Id:               UTC,
		LocationAsString: "UTC",
		Label:            "UTC",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               HST_Pacific_Honolulu_US,
		LocationAsString: "Pacific/Honolulu",
		Label:            "HST (Hawaii), US",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               AKST_AKDT_Alaska_US,
		LocationAsString: "America/Anchorage",
		Label:            "AKST/AKDT (Alaska), US",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               PST_PDT_Pacific_US,
		LocationAsString: "America/Los_Angeles",
		Label:            "PST/PDT (Pacific), US",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               MST_Mountain_US,
		LocationAsString: "America/Phoenix",
		Label:            "MST (Mountain), US",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               CST_CDT_Central_US,
		LocationAsString: "America/Chicago",
		Label:            "CST/CDT (Central), US",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               EST_EDT_Eastern_US,
		LocationAsString: "America/New_York",
		Label:            "EST/EDT (Eastern), US",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               AST_Atlantic_GD,
		LocationAsString: "America/Grenada",
		Label:            "AST (Atlantic), GD",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               GMT_BST_Greenwich_UK,
		LocationAsString: "Europe/London",
		Label:            "GMT/BST (Greenwich), UK",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               WET_WEST_Western_Europe,
		LocationAsString: "Europe/Lisbon",
		Label:            "WET/WEST (Western Europe), Portugal",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               CET_CEST_Central_Europe_France,
		LocationAsString: "Europe/Paris",
		Label:            "CET/CEST (Central Europe), France",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               EET_EEST_Eastern_Europe_Finland,
		LocationAsString: "Europe/Helsinki",
		Label:            "EET/EEST (Eastern Europe), Finland",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               MSK_Moscow_Russia,
		LocationAsString: "Europe/Moscow",
		Label:            "MSK (Moscow), Russia",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               IST_India_India,
		LocationAsString: "Asia/Kolkata",
		Label:            "IST (India), India",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               CST_China_China,
		LocationAsString: "Asia/Chongqing",
		Label:            "CST (China), China",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               AEST_AEDT_Australia_Australia,
		LocationAsString: "Australia/Sydney",
		Label:            "AEST/AEDT (Australia), Australia",
		Type:             WithLocationTimzoneType,
	},
	{
		Id:               LastNamedId,
		LocationAsString: "UTC",
		Label:            "UTC-11",
		Offset:           -11 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 1,
		LocationAsString: "UTC",
		Label:            "UTC-10",
		Offset:           -10 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 2,
		LocationAsString: "UTC",
		Label:            "UTC-9",
		Offset:           -9 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 3,
		LocationAsString: "UTC",
		Label:            "UTC-8",
		Offset:           -8 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 4,
		LocationAsString: "UTC",
		Label:            "UTC-7",
		Offset:           -7 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 5,
		LocationAsString: "UTC",
		Label:            "UTC-6",
		Offset:           -6 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 6,
		LocationAsString: "UTC",
		Label:            "UTC-5",
		Offset:           -5 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 7,
		LocationAsString: "UTC",
		Label:            "UTC-4",
		Offset:           -4 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 8,
		LocationAsString: "UTC",
		Label:            "UTC-3",
		Offset:           -3 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 9,
		LocationAsString: "UTC",
		Label:            "UTC-2",
		Offset:           -2 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 10,
		LocationAsString: "UTC",
		Label:            "UTC-1",
		Offset:           -1 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 11,
		LocationAsString: "UTC",
		Label:            "UTC+1",
		Offset:           1 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 12,
		LocationAsString: "UTC",
		Label:            "UTC+2",
		Offset:           2 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 13,
		LocationAsString: "UTC",
		Label:            "UTC+3",
		Offset:           3 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 14,
		LocationAsString: "UTC",
		Label:            "UTC+4",
		Offset:           4 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 15,
		LocationAsString: "UTC",
		Label:            "UTC+5",
		Offset:           5 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 16,
		LocationAsString: "UTC",
		Label:            "UTC+6",
		Offset:           6 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 17,
		LocationAsString: "UTC",
		Label:            "UTC+7",
		Offset:           7 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 18,
		LocationAsString: "UTC",
		Label:            "UTC+8",
		Offset:           8 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 19,
		LocationAsString: "UTC",
		Label:            "UTC+9",
		Offset:           9 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 20,
		LocationAsString: "UTC",
		Label:            "UTC+10",
		Offset:           10 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
	{
		Id:               LastNamedId + 21,
		LocationAsString: "UTC",
		Label:            "UTC+11",
		Offset:           11 * 60 * 60,
		Type:             FixedOffsetTimezoneType,
	},
}
