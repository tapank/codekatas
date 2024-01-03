package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	if t, err := time.Parse("1/02/2006 15:04:05", date); err != nil {
		return time.Now()
	} else {
		return t
	}
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	if t, err := time.Parse("January 2, 2006 15:04:05", date); err != nil {
		return false
	} else {
		return time.Now().After(t)
	}
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	if t, err := time.Parse("Monday, January 2, 2006 15:04:05", date); err != nil {
		return false
	} else {
		return t.Hour() >= 12 && t.Hour() < 18
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	if t, err := time.Parse("1/2/2006 15:04:05", date); err != nil {
		return ""
	} else {
		return t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
	}
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
