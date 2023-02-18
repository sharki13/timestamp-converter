package timezone

import (
	"encoding/json"
	"strconv"
	"time"
)

type TimezoneType int

const (
	LocalTimezoneType TimezoneType = iota
	WithLocationTimzoneType
	UnixTimezoneType
)

type TimezoneDefinition struct {
	Id               int
	LocationAsString string
	Label            string
	Type             TimezoneType
}

func (td TimezoneDefinition) StringTime(t time.Time, format string) string {
	if td.Type == UnixTimezoneType {
		return strconv.FormatInt(t.Unix(), 10)
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

type Preset struct {
	Id        int    `json:"-"`
	Label     string `json:"label"`
	Timezones []int  `json:"timezones"`
}

func SerializePresets(e []Preset) (string, error) {
	ret, err := json.Marshal(e)
	if err != nil {
		return "", err
	}

	return string(ret), nil
}

func DeserializePresets(s string) ([]Preset, error) {
	var ret []Preset
	err := json.Unmarshal([]byte(s), &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// Id 0 is reserved for the None preset

const (
	DeveloperPreset int = iota + 1
	USPreset
	EuropePreset
	USEuropePreset
	LastInternalId
)

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
	CET_CEST_Central_Europe_France
	EET_EEST_Eastern_Europe_Finland
	MSK_Moscow_Russia
	IST_India_India
	CST_China_China
	AEST_AEDT_Australia_Australia
	UTC
)

var DefaultPresets = []Preset{
	{
		Id:        DeveloperPreset,
		Label:     "Developer",
		Timezones: []int{Local, Unix, UTC, PST_PDT_Pacific_US, CST_CDT_Central_US, CET_CEST_Central_Europe_France},
	},
	{
		Id:        USPreset,
		Label:     "US",
		Timezones: []int{Local, HST_Pacific_Honolulu_US, AKST_AKDT_Alaska_US, PST_PDT_Pacific_US, MST_Mountain_US, CST_CDT_Central_US, EST_EDT_Eastern_US, AST_Atlantic_GD},
	},
	{
		Id:        EuropePreset,
		Label:     "Europe",
		Timezones: []int{Local, GMT_BST_Greenwich_UK, CET_CEST_Central_Europe_France, EET_EEST_Eastern_Europe_Finland, MSK_Moscow_Russia},
	},
	{
		Id:        USEuropePreset,
		Label:     "US & Europe",
		Timezones: []int{Local, HST_Pacific_Honolulu_US, AKST_AKDT_Alaska_US, PST_PDT_Pacific_US, MST_Mountain_US, CST_CDT_Central_US, EST_EDT_Eastern_US, GMT_BST_Greenwich_UK, CET_CEST_Central_Europe_France, EET_EEST_Eastern_Europe_Finland, MSK_Moscow_Russia},
	},
}

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
		Id:               UTC,
		LocationAsString: "UTC",
		Label:            "UTC",
		Type:             WithLocationTimzoneType,
	},
}