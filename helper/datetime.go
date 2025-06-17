package helper

import (
	"time"
)

// IsWeekend returns true if the given date falls on a Saturday or Sunday.
// It checks the Weekday of the provided time.Time value and determines
// whether it is a weekend day.
func IsWeekend(date time.Time) bool {
	return date.Weekday() == time.Saturday || date.Weekday() == time.Sunday
}

// WorkingDaysCount returns the number of weekdays (Mon–Fri)
func WorkingDaysCount(startDate, endDate time.Time) int {
	if endDate.Before(startDate) {
		return 0
	}

	start := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	end := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, endDate.Location())

	count := 0
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday {
			count++
		}
	}
	return count
}
