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

// timeUnit stores information about a single unit of time.
type timeUnit struct {
	value         int64
	nameSingular  string
	namePlural    string
	nameAbbrev    string
	showZero      bool // Flag to show this unit even if its value is zero (e.g., "0 seconds")
	onlyIfSeconds bool // Flag to show this unit only if total duration is less than 60 seconds (for milliseconds)
}

// String returns a human readable string using the Long Process Formatter.
func (a LongProcessFormatter) String(td time.Duration) string {
	o := ""
	d := TimeDurationToDuration(td)

	units := []timeUnit{
		{value: d.Days, nameSingular: "day", namePlural: "days", nameAbbrev: "d"},
		{value: d.Hours, nameSingular: "hour", namePlural: "hours", nameAbbrev: "h"},
		{value: d.Minutes, nameSingular: "minute", namePlural: "minutes", nameAbbrev: "m"},
		{value: d.Seconds, nameSingular: "second", namePlural: "seconds", nameAbbrev: "s", showZero: true},
		// Milliseconds also needs showZero for cases like "0ms" when duration is < 1ms and showmsonsec=true
		{value: d.Milliseconds, nameSingular: "millisecond", namePlural: "milliseconds", nameAbbrev: "ms", onlyIfSeconds: true, showZero: true},
	}

	hasContent := false
	for _, unit := range units {
		if unit.onlyIfSeconds && (!a.showmsonsec || td.Seconds() >= 60) {
			continue
		}

		currentValue := unit.value

		// If the current unit is 'seconds' (nameAbbrev == "s")
		// AND its value is 0
		// AND we are showing milliseconds on seconds (a.showmsonsec)
		// AND the millisecond part of the duration is greater than 0
		// THEN, we skip adding "0s" to prevent "0s999ms"; "999ms" is preferred.
		if unit.nameAbbrev == "s" && currentValue == 0 && a.showmsonsec && d.Milliseconds > 0 {
			continue
		}

		// General handling for zero-value units that are marked with showZero:
		if currentValue == 0 && unit.showZero {
			if hasContent {
				// If other, larger units ("1 day", "2 hours") are already part of the string,
				// then don't add this zero-value unit (e.g., "0 seconds" or "0ms").
				continue
			}
			// If no content yet, this zero-value unit (e.g. "0 seconds" or "0ms") should be formatted.
			// formatUnit will be called with hasContent = false.
		} else if currentValue == 0 && !unit.showZero {
			// If value is 0 and showZero is false, always skip.
			continue
		}
		// If currentValue > 0, it will proceed to formatUnit.

		if valStr, ok := a.formatUnit(unit, hasContent); ok {
			o += valStr
			if len(valStr) > 0 { // Ensure valStr actually contributed content
				hasContent = true
			}
		}
	}

	if len(o) == 0 {
		// Fallback for empty output: default to "0 seconds" or "0s"
		// This typically happens for durations less than 1ms or if all units are filtered out.
		if a.abbreviated {
			return "0s"
		}
		return "0 seconds"
	}

	if a.nospaces {
		return o
	}

	// Remove trailing space if content was added and spaces are not disabled
	if hasContent && !a.nospaces && len(o) > 0 {
		return o[:len(o)-1]
	}
	return o // Should not happen if hasContent is true and nospaces is false
}

// formatUnit formats a single time unit based on the formatter's options.
// hasContent indicates if any prior time units have already been added to the output string.
func (a LongProcessFormatter) formatUnit(unit timeUnit, hasContent bool) (string, bool) {
	// If the unit's value is zero:
	if unit.value == 0 {
		// If this is milliseconds (onlyIfSeconds=true), and its value is 0,
		// and we are showing ms on seconds (a.showmsonsec=true),
		// but other content (like "Ns") is already present, then don't show "0ms".
		if unit.onlyIfSeconds && a.showmsonsec && hasContent {
			return "", false
		}

		// If showZero is true (typically for seconds or milliseconds under specific conditions):
		if unit.showZero {
			// And if other content already exists (e.g., "1 day"), then don't add "0 seconds".
			// This is a bit redundant with the String() method's hasContent check for showZero units,
			// but provides an additional safeguard.
			if hasContent {
				return "", false
			}
			// Otherwise, format the zero value (e.g., "0 seconds" or "0s").
			// This will also handle "0ms" if showmsonsec is true and it's the only unit.
		} else if !(unit.onlyIfSeconds && a.showmsonsec) {
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
		return fmt.Sprintf("%d%s%s", unit.value, unit.nameAbbrev, sp), true
	}

	if unit.value == 1 {
		return fmt.Sprintf("%d%s%s%s", unit.value, sep, unit.nameSingular, sp), true
	}

	return fmt.Sprintf("%d%s%s%s", unit.value, sep, unit.namePlural, sp), true
}

// particle takes a unit and the display units and returns either the formatted unit and true or a empty
// string and false.
// This method is kept for now as formatUnit is not yet fully equivalent and might be reverted/adjusted.
// The particle method has been replaced by formatUnit and the loop in String().
// The individual time unit methods (days, hours, etc.) have also been replaced.
