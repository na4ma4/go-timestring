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
// It formats the duration into days, hours, minutes, seconds, and milliseconds,
// with options for abbreviated output, no spaces, and showing milliseconds on seconds.
//
//nolint:gocognit // String is the main processing function for LongProcessFormatter.
func (a LongProcessFormatter) String(td time.Duration) string {
	d := TimeDurationToDuration(td)
	units := []timeUnit{
		day.toTimeUnit(d.Days),
		hour.toTimeUnit(d.Hours),
		minute.toTimeUnit(d.Minutes),
		second.toTimeUnit(d.Seconds),
		millisecond.toTimeUnit(d.Milliseconds),
	}

	var (
		o          string
		hasContent bool
	)

	for _, unit := range units {
		// Skip ms if not showing ms on seconds or duration >= 60s
		if unit.IsOnlyIfSeconds() && (!a.showmsonsec || td.Seconds() >= 60) {
			continue
		}
		// Skip "0s" if showing ms and ms > 0
		if unit.GetNameAbbrev() == "s" && unit.value == 0 && a.showmsonsec && d.Milliseconds > 0 {
			continue
		}
		// Skip zero units unless showZero and no content yet
		if unit.value == 0 && (!unit.IsShowZero() || hasContent) {
			continue
		}
		if valStr, ok := a.formatUnit(unit, hasContent); ok {
			o += valStr
			if len(valStr) > 0 {
				hasContent = true
			}
		}
	}

	if o == "" {
		if a.abbreviated {
			return "0s"
		}
		return "0 seconds"
	}

	if a.nospaces {
		return o
	}

	// Remove trailing space if present
	if o[len(o)-1] == ' ' {
		return o[:len(o)-1]
	}
	return o
}

// formatUnit formats a single time unit based on the formatter's options.
// hasContent indicates if any prior time units have already been added to the output string.
//
//nolint:nestif // the nesting is for handling zero values and special cases.
func (a LongProcessFormatter) formatUnit(unit timeUnit, hasContent bool) (string, bool) {
	// If the unit's value is zero:
	if unit.value == 0 {
		// If this is milliseconds (onlyIfSeconds=true), and its value is 0,
		// and we are showing ms on seconds (a.showmsonsec=true),
		// but other content (like "Ns") is already present, then don't show "0ms".
		if unit.IsOnlyIfSeconds() && a.showmsonsec && hasContent {
			return "", false
		}

		// If showZero is true (typically for seconds or milliseconds under specific conditions):
		if unit.IsShowZero() {
			// And if other content already exists (e.g., "1 day"), then don't add "0 seconds".
			// This is a bit redundant with the String() method's hasContent check for showZero units,
			// but provides an additional safeguard.
			if hasContent {
				return "", false
			}
			// Otherwise, format the zero value (e.g., "0 seconds" or "0s").
			// This will also handle "0ms" if showmsonsec is true and it's the only unit.
		} else if !unit.IsOnlyIfSeconds() && !a.showmsonsec {
			// If showZero is false, and it's not the special milliseconds case (where value might be 0 but still shown), then skip.
			return "", false
		}
		// If it is the special milliseconds case (unit.onlyIfSeconds && a.showmsonsec) and value is 0,
		// (and hasContent is false, checked above), it means we want to show "0ms" or "0 milliseconds", so proceed to formatting.
	}

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

	if a.abbreviated {
		return fmt.Sprintf("%d%s%s", unit.value, unit.GetNameAbbrev(), sp), true
	}

	if unit.value == 1 {
		return fmt.Sprintf("%d%s%s%s", unit.value, sep, unit.GetNameSingular(), sp), true
	}

	return fmt.Sprintf("%d%s%s%s", unit.value, sep, unit.GetNamePlural(), sp), true
}
