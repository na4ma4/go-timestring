package timestring

import (
	"fmt"
	"math"
	"time"
)

var LongProcess = LongProcessFormatter{}

type LongProcessFormatter struct{}

func (a LongProcessFormatter) String(td time.Duration) string {
	o := ""
	days := int64(0)
	hours := int64(math.Trunc(td.Hours()))
	if hours >= 24 {
		days = int64(math.Trunc(float64(hours) / 24))
		hours -= days * 24
	}
	minutes := int64(math.Trunc(math.Mod(td.Minutes(), 60)))
	seconds := int64(math.Trunc(math.Mod(td.Seconds(), 60)))
	if days > 0 {
		if days == 1 {
			o += fmt.Sprintf("%d day ", days)
		} else {
			o += fmt.Sprintf("%d days ", days)
		}
	}
	if hours > 0 {
		if hours == 1 {
			o += fmt.Sprintf("%d hour ", hours)
		} else {
			o += fmt.Sprintf("%d hours ", hours)
		}
	}
	if minutes > 0 {
		if minutes == 1 {
			o += fmt.Sprintf("%d minute ", minutes)
		} else {
			o += fmt.Sprintf("%d minutes ", minutes)
		}
	}
	if seconds > 0 {
		if seconds == 1 {
			o += fmt.Sprintf("%d second ", seconds)
		} else {
			o += fmt.Sprintf("%d seconds ", seconds)
		}
	}

	if len(o) == 0 {
		return "0 seconds"
	}

	return o[:len(o)-1]
}
