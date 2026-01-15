package timestring

import (
	"fmt"
	"strings"
	"time"
)

// Absolute is the ready-to-use Absolute Formatter.
//
//nolint:gochecknoglobals // pre initialised formatter.
var Absolute Formatter = AbsoluteFormatter{abbreviated: true}

// AbsoluteFormatter is a Absolute Formatter.
// It provides a concise and precise representation of time.Duration,
// always using abbreviated units and omitting zero-value units.
type AbsoluteFormatter struct {
	nospaces     bool
	nounitspaces bool
	abbreviated  bool // Always true for AbsoluteFormatter, but kept for interface compatibility
}

// Option returns a Absolute Formatter with the applied options.
// For AbsoluteFormatter, Abbreviated is always true.
// ShowMSOnSeconds is not applicable.
func (s AbsoluteFormatter) Option(opts ...FormatterOption) Formatter {
	s.abbreviated = true // Ensure abbreviated is always true
	for _, opt := range opts {
		switch opt {
		case NoSpaces:
			s.nospaces = true
		case NoUnitSpaces:
			s.nounitspaces = true
		case Abbreviated:
			// Already true, do nothing
		case ShowMSOnSeconds:
			// Not applicable for AbsoluteFormatter
		}
	}
	return s
}

// String returns a human readable string using the Absolute Formatter.
// It always uses abbreviated units and omits zero-value units.
//
// Example: "1d 2h 3m 4s", "2h 3m", "4s", "0s".
func (s AbsoluteFormatter) String(td time.Duration) string {
	if td == 0 {
		return "0s" // Absolute formatter usually returns 0s even with spaces.
	}

	d := TimeDurationToDuration(td)
	units := []timeUnit{
		unitDay.toTimeUnit(d.Days),
		unitHour.toTimeUnit(d.Hours),
		unitMinute.toTimeUnit(d.Minutes),
		unitSecond.toTimeUnit(d.Seconds),
		unitMillisecond.toTimeUnit(d.Milliseconds),
		unitMicrosecond.toTimeUnit(d.Microseconds),
		unitNanosecond.toTimeUnit(d.Nanoseconds),
	}
	sb := &strings.Builder{}

	for _, unit := range units {
		if unit.value > 0 {
			if s.nospaces {
				fmt.Fprintf(sb, "%d%s", unit.value, unit.GetNameAbbrev())
			} else {
				fmt.Fprintf(sb, "%d%s ", unit.value, unit.GetNameAbbrev())
			}
		}
	}

	if sb.Len() == 0 {
		return "0s"
	}

	return strings.TrimSpace(sb.String())
}
