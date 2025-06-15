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
