package timestring

type globalTimeUnit struct {
	nameSingular  string
	namePlural    string
	nameAbbrev    string
	showZero      bool // Flag to show this unit even if its value is zero (e.g., "0 seconds")
	onlyIfSeconds bool // Flag to show this unit only if total duration is less than 60 seconds (for milliseconds)
}

// toTimeUnit converts a globalTimeUnit to a timeUnit with the specified value.
// This is used to create timeUnit instances with specific values in the formatters.
// It allows for easy conversion from the global time unit definitions to the specific time unit instances
// used in the ShortProcessFormatter and LongProcessFormatter.
func (gtu globalTimeUnit) toTimeUnit(value int64) timeUnit {
	return timeUnit{
		value: value,
		unit:  gtu,
	}
}

// IsOnlyIfSeconds returns true if this global time unit should only be shown if the total duration is less than 60 seconds.
func (gtu globalTimeUnit) IsOnlyIfSeconds() bool {
	return gtu.onlyIfSeconds
}

// IsShowZero returns true if this global time unit should be shown even if its value is zero.
func (gtu globalTimeUnit) IsShowZero() bool {
	return gtu.showZero
}

// GetNameSingular returns the singular name of the time unit.
func (gtu globalTimeUnit) GetNameSingular() string {
	return gtu.nameSingular
}

// GetNamePlural returns the plural name of the time unit.
func (gtu globalTimeUnit) GetNamePlural() string {
	return gtu.namePlural
}

// GetNameAbbrev returns the abbreviated name of the time unit.
func (gtu globalTimeUnit) GetNameAbbrev() string {
	return gtu.nameAbbrev
}

// Predefined global time units for easy access and consistency across the package.
// These are used in both ShortProcessFormatter and LongProcessFormatter.
// They are defined as global variables to avoid duplication and ensure consistent naming.
// They are not meant to be modified and should be used as constants.
// They are used to create timeUnit instances with specific values in the formatters.
//
//nolint:gochecknoglobals // These are constants for time units, not global state.
var (
	day = globalTimeUnit{
		nameSingular: "day", namePlural: "days", nameAbbrev: "d",
	}
	hour = globalTimeUnit{
		nameSingular: "hour", namePlural: "hours", nameAbbrev: "h",
	}
	minute = globalTimeUnit{
		nameSingular: "minute", namePlural: "minutes", nameAbbrev: "m",
	}
	second = globalTimeUnit{
		nameSingular: "second", namePlural: "seconds", nameAbbrev: "s", showZero: true,
	}
	millisecond = globalTimeUnit{
		nameSingular: "millisecond", namePlural: "milliseconds", nameAbbrev: "ms", onlyIfSeconds: true, showZero: true,
	}
)

// timeUnit stores information about a single unit of time.
type timeUnit struct {
	value int64
	unit  globalTimeUnit // Reference to the global time unit definition
}

// IsOnlyIfSeconds returns true if this time unit should only be shown if the total duration is less than 60 seconds.
func (tu timeUnit) IsOnlyIfSeconds() bool {
	return tu.unit.IsOnlyIfSeconds()
}

// IsShowZero returns true if this time unit should be shown even if its value is zero.
func (tu timeUnit) IsShowZero() bool {
	return tu.unit.IsShowZero()
}

// GetNameSingular returns the singular name of the time unit.
func (tu timeUnit) GetNameSingular() string {
	return tu.unit.GetNameSingular()
}

// GetNamePlural returns the plural name of the time unit.
func (tu timeUnit) GetNamePlural() string {
	return tu.unit.GetNamePlural()
}

// GetNameAbbrev returns the abbreviated name of the time unit.
func (tu timeUnit) GetNameAbbrev() string {
	return tu.unit.GetNameAbbrev()
}
