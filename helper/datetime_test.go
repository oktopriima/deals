package helper

import (
	"testing"
	"time"
)

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected bool
	}{
		{
			name:     "Saturday",
			date:     time.Date(2024, 6, 8, 0, 0, 0, 0, time.UTC), // Saturday
			expected: true,
		},
		{
			name:     "Sunday",
			date:     time.Date(2024, 6, 9, 0, 0, 0, 0, time.UTC), // Sunday
			expected: true,
		},
		{
			name:     "Monday",
			date:     time.Date(2024, 6, 10, 0, 0, 0, 0, time.UTC), // Monday
			expected: false,
		},
		{
			name:     "Wednesday",
			date:     time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC), // Wednesday
			expected: false,
		},
		{
			name:     "Friday",
			date:     time.Date(2024, 6, 14, 0, 0, 0, 0, time.UTC), // Friday
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsWeekend(tt.date)
			if result != tt.expected {
				t.Errorf("IsWeekend(%v) = %v; want %v", tt.date.Weekday(), result, tt.expected)
			}
		})
	}
}
