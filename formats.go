package main

import "time"

type FormatDefinition struct {
	Format string
	Label  string
}

var SupportedFormats = []FormatDefinition{
	{
		Format: time.RFC3339,
		Label:  "RFC3339 (2006-01-02T15:04:05Z07:00)",
	},
	{
		Format: time.RubyDate,
		Label:  "Ruby Date (Mon Jan 2 15:04:05 -0700 2006)",
	},
	{
		Format: time.RFC822Z,
		Label:  "RFC822Z (02 Jan 06 15:04 -0700)",
	},
	{
		Format: time.RFC1123Z,
		Label:  "RFC1123Z (Mon, 02 Jan 2006 15:04:05 -0700)",
	},
}
