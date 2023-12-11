package utils

import (
	"time"
)

func StringToTime(str string) time.Time {
	layout := "2006-01-02 15:04:05" // Specify the layout of the string
	t, err := time.Parse(layout, str)
	if err != nil {
		return time.Time{}
	}
	return t
}
