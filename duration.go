package timestring

import (
	"math"
	"time"
)

// Duration contains the absolute number of each field with all fields adding up to
// the total duration, so Hours is only the amount of hours to be displayed,
// unlike (time.Duration).Hours().
type Duration struct {
	Days         int64
	Hours        int64
	Minutes      int64
	Seconds      int64
	Milliseconds int64
	Microseconds int64
	Nanoseconds  int64
}

// TimeDurationToDuration converts a time.Duration to the timestring.Duration for easier
// display of durations.
//
//nolint:mnd // These _are_ magic numbers.
func TimeDurationToDuration(td time.Duration) Duration {
	// d := Duration{}
	// d.Hours = int64(math.Trunc(td.Hours()))
	// if d.Hours >= 24 {
	// 	d.Days = int64(math.Trunc(float64(d.Hours) / 24))
	// 	d.Hours -= d.Days * 24
	// }
	// d.Minutes = int64(math.Trunc(math.Mod(td.Minutes(), 60)))
	// d.Seconds = int64(math.Trunc(math.Mod(td.Seconds(), 60)))
	// return d
	return Duration{
		Days:         int64(math.Trunc(td.Hours() / 24)),
		Hours:        int64(math.Trunc(math.Mod(td.Hours(), 24))),
		Minutes:      int64(math.Trunc(math.Mod(td.Minutes(), 60))),
		Seconds:      int64(math.Trunc(math.Mod(td.Seconds(), 60))),
		Milliseconds: int64(math.Trunc(math.Mod(float64(td.Milliseconds()), 1000))),
		Microseconds: int64(math.Trunc(math.Mod(float64(td.Microseconds()), 1000))),
		Nanoseconds:  int64(math.Trunc(math.Mod(float64(td.Nanoseconds()), 1000))),
	}
}
