package main

import "time"

var FormatLabelMap = map[string]string{
	time.RFC3339:  "RFC3339 (2006-01-02T15:04:05Z07:00)",
	time.RubyDate: "Ruby Date (Mon Jan 2 15:04:05 -0700 2006)",
	time.RFC822Z:  "RFC822Z (02 Jan 06 15:04 -0700)",
	time.RFC1123Z: "RFC1123Z (Mon, 02 Jan 2006 15:04:05 -0700)",
}
