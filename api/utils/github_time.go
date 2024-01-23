package utils

import (
	"time"

	"github.com/Improwised/GPAT/constants"
)

func GetWeekTimestamps() (startOfWeek time.Time, endOfWeek time.Time) {
	now := time.Now().UTC()
	before7Days := now.AddDate(0, 0, -7) // Subtract 7 days
	return now, before7Days
}

func ParseTimeFromString(t string) (time.Time, error) {
	formatedTime, err := time.Parse(constants.DateTimeFormat, t)
	if err != nil {
		return formatedTime, err
	}
	return formatedTime, nil
}
