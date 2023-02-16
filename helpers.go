package main

import (
	"fmt"
	"strconv"
	"time"
)

func PraseStringToTime(s string) (time.Time, error) {
	for _, format := range SupportedFormats {
		t, err := time.Parse(format.Format, s)
		if err == nil {
			return t, nil
		}
	}

	intT, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return time.Unix(intT, 0), nil
	}

	return time.Time{}, fmt.Errorf("invalid time format")
}

func contains[K comparable](s []K, e K) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
