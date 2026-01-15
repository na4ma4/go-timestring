package timestring

import (
	"strings"
	"time"
)

// LongProcess is the ready-to-use Long Process Formatter.
//
//nolint:gochecknoglobals // pre initialised formatter.
var LongProcess Formatter = LongProcessFormatter{}

// LongProcessFormatter is a Long Process Formatter.
//
// It is a formatter that handles processes that would be considered long running,
// like displaying the uptime of a server or service.
type LongProcessFormatter struct {
	nospaces     bool
	nounitspaces bool
	showmsonsec  bool
	abbreviated  bool
}

// Option returns a Long Process Formatter with the applied options.
func (a LongProcessFormatter) Option(opts ...FormatterOption) Formatter {
	for _, opt := range opts {
		switch opt {
		case NoSpaces:
			a.nospaces = true
		case NoUnitSpaces:
			a.nounitspaces = true
		case Abbreviated:
			a.abbreviated = true
		case ShowMSOnSeconds:
			a.showmsonsec = true
		}
	}

	return a
}

// String returns a human readable string using the Long Process Formatter.
// It formats the duration into days, hours, minutes, seconds, and milliseconds,
// with options for abbreviated output, no spaces, and showing milliseconds on seconds.
func (a LongProcessFormatter) String(td time.Duration) string {
	d := TimeDurationToDuration(td)
	units := []timeUnit{
		unitDay.toTimeUnit(d.Days),
		unitHour.toTimeUnit(d.Hours),
		unitMinute.toTimeUnit(d.Minutes),
		unitSecond.toTimeUnit(d.Seconds),
		unitMillisecond.toTimeUnit(d.Milliseconds),
	}

	var hasContent bool
	o := &strings.Builder{}

	for _, unit := range units {
		// Skip ms if not showing ms on seconds or duration >= 60s
		if unit.IsOnlyIfSeconds() && (!a.showmsonsec || td.Seconds() >= 60) {
			continue
		}
		// Skip "0s" if showing ms and ms > 0
		if unit.IsGlobalUnit(unitSecond) && unit.value == 0 && a.showmsonsec && d.Milliseconds > 0 {
			continue
		}
		// Skip zero units unless showZero and no content yet
		if unit.value == 0 && (!unit.IsShowZero() || hasContent) {
			continue
		}

		if valStr := unit.String(a.abbreviated, !a.nounitspaces); len(valStr) > 0 {
			o.WriteString(valStr)
			if !a.nospaces {
				o.WriteString(" ")
			}

			hasContent = true
		}
	}

	if o.Len() == 0 {
		return unitSecond.toTimeUnit(0).String(a.abbreviated, !a.nounitspaces)
	}

	return strings.TrimSpace(o.String())
}
