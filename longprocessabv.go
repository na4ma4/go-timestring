package timestring

import (
	"fmt"
	"math"
	"time"
)

var LongProcessAbv = LongProcessAbvFormatter{}

type LongProcessAbvFormatter struct{}

func (a LongProcessAbvFormatter) String(td time.Duration) string {
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
		o += fmt.Sprintf("%dd ", days)
	}
	if hours > 0 {
		o += fmt.Sprintf("%dh ", hours)
	}
	if minutes > 0 {
		o += fmt.Sprintf("%dm ", minutes)
	}
	if seconds > 0 {
		o += fmt.Sprintf("%ds ", seconds)
	}

	if len(o) == 0 {
		return "0s"
	}

	return o[:len(o)-1]
}
