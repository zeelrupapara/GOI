package utils

import (
	"strconv"
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

func ConvertEpochToTime(t string) (time.Time, error) {
	epoch, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	seconds := epoch / 1000
	datetime := time.Unix(seconds, 0).UTC()
	return datetime, nil
}

func ConvertIntToTime(t int) (time.Time, error) {
	seconds := t / 1000
	datetime := time.Unix(int64(seconds), 0).UTC()
	return datetime, nil
}

// Split time by specific time duration
func SplitTimeRange(start, end time.Time, interval time.Duration) [][2]time.Time {
	var subRanges [][2]time.Time

	currentTime := start
	for currentTime.Before(end) {
		subEnd := currentTime.Add(interval)
		if subEnd.After(end) {
			subEnd = end
		}
		subRanges = append(subRanges, [2]time.Time{currentTime, subEnd})
		currentTime = subEnd
	}

	return subRanges
}
