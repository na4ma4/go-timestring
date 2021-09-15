package timestring

import (
	"time"
)

// Formatter is the interface that any timestring formatter should match.
type Formatter interface {
	Option(...FormatterOption) Formatter
	String(time.Duration) string
}

// FormatterOption is a list of options that can be applied to the standard formatters.
type FormatterOption uint

const (
	// NoSpaces is a FormatterOption that tells the formatter to ignore spaces between values.
	NoSpaces FormatterOption = iota

	// NoUnitSpaces is a FormatterOption that tells the formatter to ignore spaces betwee values and
	// units.
	NoUnitSpaces

	// Abbreviated  is a FormatterOption that tells the formatter to abbreviate the units
	// (eg. "d" instead of "days").
	Abbreviated
)
