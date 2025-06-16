package helper

import "time"

func CheckOvertimeStart(date time.Time) bool {
	// in weekend user can submit overtime in any hour
	if IsWeekend(date) {
		return true
	}
	// On weekdays, allow only after 5 PM
	hour := date.Hour()
	minute := date.Minute()
	sec := date.Second()

	// Check if time is after 17:00:00
	if hour > 17 || (hour == 17 && (minute > 0 || sec > 0)) {
		return true
	}

	return false
}
