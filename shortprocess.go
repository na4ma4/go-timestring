package timestring

import (
	"fmt"
	"strings"
	"time"
)

// ShortProcess is the ready-to-use Short Process Formatter.
//
//nolint:gochecknoglobals // pre initialised formatter.
var ShortProcess Formatter = ShortProcessFormatter{abbreviated: true}

// ShortProcessFormatter is a Short Process Formatter.
// It provides a concise representation of time.Duration,
// always using abbreviated units and omitting zero-value units.
type ShortProcessFormatter struct {
	nospaces     bool
	nounitspaces bool
	abbreviated  bool // Always true for ShortProcessFormatter, but kept for interface compatibility
}

// Option returns a Short Process Formatter with the applied options.
// For ShortProcessFormatter, Abbreviated is always true.
// ShowMSOnSeconds is not applicable.
func (s ShortProcessFormatter) Option(opts ...FormatterOption) Formatter {
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
			// Not applicable for ShortProcessFormatter
		}
	}
	return s
}

// String returns a human readable string using the Short Process Formatter.
// It always uses abbreviated units and omits zero-value units.
//
// Example: "1d 2h 3m 4s", "2h 3m", "4s", "0s".
func (s ShortProcessFormatter) String(td time.Duration) string {
	if td == 0 {
		if s.nospaces && s.nounitspaces {
			return "0s"
		}
		if s.nospaces { // implies not nounitspaces
			return "0s"
		}
		// default
		return "0s" // Short process usually returns 0s even with spaces.
	}

	d := TimeDurationToDuration(td)
	parts := []string{}

	units := []struct {
		value int64
		name  string
	}{
		{value: d.Days, name: "d"},
		{value: d.Hours, name: "h"},
		{value: d.Minutes, name: "m"},
		{value: d.Seconds, name: "s"},
		{value: d.Milliseconds, name: "ms"},
	}

	for _, unit := range units {
		if unit.value > 0 {
			parts = append(parts, fmt.Sprintf("%d%s", unit.value, unit.name))
		}
	}

	if len(parts) == 0 {
		// This case handles durations like 0.5ms, which would have all unit values as 0.
		// Or if only milliseconds was > 0 but td < 1ms (so d.Milliseconds is 0).
		// For very small non-zero durations that don't make up 1ms.
		if s.nounitspaces {
			return "0s" // Fallback to 0s, could be 0ms if preferred, but 0s is common.
		}
		return "0s"
	}

	// If NoUnitSpaces is true, the space is already handled (or not handled) by the part formatting.
	// If NoSpaces is true, all spaces are removed.
	// The fmt.Sprintf in unit loop already adds the unit like "1d", "2h".
	// So we just need to join them.
	// If nounitspaces is true, e.g. "1d2h". This is controlled by separator.
	// If nospaces is true, e.g. "1d2h3m". This is also controlled by separator.
	// The existing logic in `particle` from longprocess for nospaces removes the final space.
	// Here, strings.Join adds spaces, then we might remove them.

	// Let's refine the part formatting and joining logic.
	// The `particle` method in longprocess.go has specific logic for spaces.
	// `fmt.Sprintf("%d%s%s", v, unitAbr, sp)` where sp is space if !nospaces
	// ShortProcessFormatter:
	// - abbreviated: always true
	// - nospaces: no spaces between "1d 2h" -> "1d2h"
	// - nounitspaces: no space between value and unit "1 d" -> "1d" (already default for abbreviated)

	// Corrected part generation and joining:
	processedParts := []string{}
	for _, unit := range units {
		if unit.value > 0 {
			unitStr := unit.name
			// nounitspaces is implicitly handled by "%d%s" for abbreviated units.
			// No need for 'sep' like in longprocess.
			processedParts = append(processedParts, fmt.Sprintf("%d%s", unit.value, unitStr))
		}
	}

	if len(processedParts) == 0 {
		return "0s" // Default for zero or very small durations
	}

	if s.nospaces { // This implies no spaces anywhere, including between number-unit pairs
		return strings.Join(processedParts, "")
	}

	// Default: join with spaces
	return strings.Join(processedParts, " ")
}
