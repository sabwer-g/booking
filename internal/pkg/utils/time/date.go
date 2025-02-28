package time

import (
	"time"
)

func DaysBetween(from time.Time, to time.Time) map[time.Time]struct{} {
	if from.After(to) {
		return nil
	}

	days := make(map[time.Time]struct{})
	for d := ToDay(from); !d.After(ToDay(to)); d = d.AddDate(0, 0, 1) {
		days[d] = struct{}{}
	}

	return days
}

func ToDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
