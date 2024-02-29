package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First WeekSchedule = iota + 1
	Second
	Third
	Fourth
	Fifth
	Last
	Teenth
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	var date time.Time

	// last weekday is a special case
	if wSched == Last {
		// start at the end of the month
		date := time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local)

		// roll back until we find the weekday
		for date.Weekday() != wDay {
			date = date.Add(time.Hour * -24)
		}
		return date.Day()
	}

	// start at the beginning of the month
	date = time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

	// roll forward to the first weekday
	for date.Weekday() != wDay {
		date = date.Add(time.Hour * 24)
	}

	switch wSched {
	case First:
		// nothing to do
	case Second:
		// roll forward one week
		date = date.Add(time.Hour * 24 * 7)
	case Third:
		// roll forward two weeks
		date = date.Add(time.Hour * 24 * 14)
	case Fourth:
		// roll forward three weeks
		date = date.Add(time.Hour * 24 * 21)
	case Fifth:
		// roll forward four weeks
		date = date.Add(time.Hour * 24 * 28)
	case Teenth:
		// roll forward until we hit the correct weekday and day
		for {
			if date.Weekday() == wDay && date.Day() >= 13 {
				break
			}
			date = date.Add(time.Hour * 24)
		}
	default:
		// we would'nt have to do this if the function also returned an error
		return -1
	}
	return date.Day()
}
