package clock

import (
	"fmt"
)

// Time holds minutes and seconds
type Time struct {
	h, m int
}

func timeToMin(t Time) int {
	return t.h * 60 + t.m
}

func minToTime(m int) Time {
	m %= 1440 // 1440 = 24 * 60
	if m < 0 {
		m += 1440
	}
	return Time {
		h: m / 60,
		m: m % 60,
	};
}

// String returns the string representation of a Time
func (t Time) String() string {
	return fmt.Sprintf("%02d:%02d", t.h, t.m)
}

// Add adds minutes to a Time
func (t Time) Add(m int) Time {
	return minToTime(timeToMin(t) + m)
}

// Subtract subtracts minutes out of a Time
func (t Time) Subtract(m int) Time {
	return minToTime(timeToMin(t) - m)
}

// New returns a Time struct
func New(h, m int) Time {
	return minToTime(h * 60 + m)
}