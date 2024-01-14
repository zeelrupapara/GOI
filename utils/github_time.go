package utils

import (
	"time"
)

func GetWeekTimestamps() (startOfWeek time.Time, endOfWeek time.Time) {
	now := time.Now()
	before7Days := now.AddDate(0, 0, -7) // Subtract 7 days
	return now, before7Days
}
