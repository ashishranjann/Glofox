package utils

import (
	"errors"
	"time"
)

var layOut = "2006-01-02"

func ParseDate(dateStr string) (time.Time, error) {
	t, err := time.Parse(layOut, dateStr)
	if err != nil {
		return time.Time{}, errors.New("invalid time format")
	}
	return t, nil

}
