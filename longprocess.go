package timestring

import (
	"fmt"
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
func (a LongProcessFormatter) String(td time.Duration) string {
	o := ""
	d := TimeDurationToDuration(td)

	if v, ok := a.days(d); ok {
		o += v
	}

	if v, ok := a.hours(d); ok {
		o += v
	}

	if v, ok := a.minutes(d); ok {
		o += v
	}

	if v, ok := a.seconds(d); ok {
		o += v
	}

	if a.showmsonsec && td.Seconds() < 60 {
		if v, ok := a.milliseconds(d); ok {
			o += v
		}
	}

	if len(o) == 0 {
		if a.abbreviated {
			return "0s"
		}

		return "0 seconds"
	}

	if a.nospaces {
		return o
	}

	return o[:len(o)-1]
}

// particle takes a unit and the display units and returns either the formatted unit and true or a empty
// string and false.
func (a LongProcessFormatter) particle(v int64, unitAbr, unitSingle, unitMultiple string) (string, bool) {
	var (
		sp  string
		sep string
	)

	if !a.nospaces {
		sp = " "
	}

	if !a.nounitspaces {
		sep = " "
	}

	if v > 0 {
		if a.abbreviated {
			return fmt.Sprintf("%d%s%s", v, unitAbr, sp), true
		}

		if v == 1 {
			return fmt.Sprintf("%d%s%s%s", v, sep, unitSingle, sp), true
		}

		return fmt.Sprintf("%d%s%s%s", v, sep, unitMultiple, sp), true
	}

	return "", false
}

func (a LongProcessFormatter) days(d Duration) (string, bool) {
	return a.particle(d.Days, "d", "day", "days")
}

func (a LongProcessFormatter) hours(d Duration) (string, bool) {
	return a.particle(d.Hours, "h", "hour", "hours")
}

func (a LongProcessFormatter) minutes(d Duration) (string, bool) {
	return a.particle(d.Minutes, "m", "minute", "minutes")
}

func (a LongProcessFormatter) seconds(d Duration) (string, bool) {
	return a.particle(d.Seconds, "s", "second", "seconds")
}

func (a LongProcessFormatter) milliseconds(d Duration) (string, bool) {
	return a.particle(d.Milliseconds, "ms", "millisecond", "milliseconds")
}
