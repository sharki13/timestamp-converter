package main

import (
	"time"
	"strconv"
	"fmt"
)

func PraseStringToTime(s string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, s)
	if err == nil {
		return t, nil
	}

	t, err = time.Parse(time.RFC3339Nano, s)
	if err == nil {
		return t, nil
	}

	intT, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return time.Unix(intT, 0), nil
	}

	return time.Time{}, fmt.Errorf("invalid time format")
}