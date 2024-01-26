package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	H, M int
}

func New(h, m int) Clock {
	// bring down the total minutes to be within a day
	minutes := (h*60 + m) % (24 * 60)
	// add a day if the minutes are negative
	if minutes < 0 {
		minutes += 24 * 60
	}
	return Clock{minutes / 60, minutes % 60}
}

func (c Clock) Add(m int) Clock {
	return New(c.H, c.M+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.H, c.M-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.H, c.M)
}
