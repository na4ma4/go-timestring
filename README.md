# go-timestring

[![Go](https://github.com/na4ma4/go-timestring/actions/workflows/go.yml/badge.svg)](https://github.com/na4ma4/go-timestring/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/na4ma4/go-timestring/?status.svg)](https://godoc.org/github.com/na4ma4/go-timestring)
[![GitHub issues](https://img.shields.io/github/issues/na4ma4/go-timestring)](https://github.com/na4ma4/go-timestring/issues)
[![GitHub forks](https://img.shields.io/github/forks/na4ma4/go-timestring)](https://github.com/na4ma4/go-timestring/network)
[![GitHub stars](https://img.shields.io/github/stars/na4ma4/go-timestring)](https://github.com/na4ma4/go-timestring/stargazers)
[![GitHub license](https://img.shields.io/github/license/na4ma4/go-timestring)](https://github.com/na4ma4/go-timestring/blob/main/LICENSE)

`go-timestring` is a Go package that provides human-readable formatters for `time.Duration`. It offers different ways to represent durations, suitable for various display needs, such as long, verbose formats or short, abbreviated ones.

## Installation

```bash
go get github.com/na4ma4/go-timestring
```

## Usage

The package provides several pre-configured formatters. You can use them directly or customize them with options.

### Formatters

#### `LongProcess`

The `LongProcess` formatter is designed for displaying durations in a full, human-readable format, similar to how one might describe uptime for a server or a long-running process.

**Default Behavior:**
- Displays time units in their full names (e.g., "days", "hours", "minutes", "seconds").
- Includes spaces between values, unit names, and different time unit parts.
- For durations less than a second, it will show "0 seconds" by default.

**Example:**

```go
package main

import (
	"fmt"
	"time"
	"github.com/na4ma4/go-timestring"
)

func main() {
	duration1, _ := time.ParseDuration("49h15m30s100ms")
	fmt.Println(timestring.LongProcess.String(duration1))
	// Output: 2 days 1 hour 15 minutes 30 seconds

	duration2 := 500 * time.Millisecond
	fmt.Println(timestring.LongProcess.String(duration2))
	// Output: 0 seconds
}
```

#### `ShortProcess` and `Absolute` (New)

The `ShortProcess` formatter provides a concise, abbreviated representation of durations. It's suitable for contexts where space is limited, but a clear representation of the duration is still needed.

**Default Behavior:**
- Uses abbreviated unit names (e.g., "d", "h", "m", "s", "ms").
- Omits time units that have a zero value.
- Includes spaces between different time unit parts (e.g., "1d 2h").
- For a zero duration or durations that round down to zero for all its units (e.g. <1ms), it displays "0s".

The `Absolute` formatter provides an absolute, precise and concise abbreviated representation of durations. It's suitable where sub-ms times are required.

**Example:**

```go
package main

import (
	"fmt"
	"time"
	"github.com/na4ma4/go-timestring"
)

func main() {
	duration1, _ := time.ParseDuration("49h15m30s100ms") // 2 days, 1 hour, 15 minutes, 30 seconds, 100 milliseconds
	fmt.Println(timestring.ShortProcess.String(duration1))
	// Output: 2d 1h 15m 30s 100ms

	duration2 := 2*time.Hour + 30*time.Minute
	fmt.Println(timestring.ShortProcess.String(duration2))
	// Output: 2h 30m

	duration3 := 5 * time.Second
	fmt.Println(timestring.ShortProcess.String(duration3))
	// Output: 5s

	duration4 := 500 * time.Millisecond
	fmt.Println(timestring.ShortProcess.String(duration4))
	// Output: 500ms

	duration5 := time.Duration(0)
	fmt.Println(timestring.ShortProcess.String(duration5))
	// Output: 0s

	duration6 := 100 * time.Nanosecond
	fmt.Println(timestring.ShortProcess.String(duration6))
	// Output: 0s
	fmt.Println(timestring.Absolute.String(duration6))
	// Output: 100ns
}
```

### Customization Options

Both formatters implement the `Formatter` interface, which includes an `Option()` method. This method allows for customization of the output string.

Available options:

- `timestring.NoSpaces`: Removes spaces between unit parts (e.g., "1d2h3m" instead of "1d 2h 3m").
- `timestring.NoUnitSpaces`: Removes spaces between the numeric value and its unit name (e.g., "1day" instead of "1 day"). Note: For abbreviated formats like `ShortProcess` or `LongProcess` with `Abbreviated` option, this has no visible effect as "1d" already has no space.
- `timestring.Abbreviated`: (Mainly for `LongProcess`) Uses abbreviated unit names (e.g., "d", "h", "m", "s"). `ShortProcess` is always abbreviated.
- `timestring.ShowMSOnSeconds`: (For `LongProcess`) Displays milliseconds when the duration is less than 60 seconds.

**Option Usage Example:**

```go
package main

import (
	"fmt"
	"time"
	"github.com/na4ma4/go-timestring"
)

func main() {
	duration, _ := time.ParseDuration("25h30m") // 1 day, 1 hour, 30 minutes

	// LongProcess with NoSpaces and Abbreviated
	formattedLong := timestring.LongProcess.Option(
		timestring.NoSpaces,
		timestring.Abbreviated,
	).String(duration)
	fmt.Println(formattedLong) // Output: 1d1h30m

	// ShortProcess with NoSpaces
	formattedShort := timestring.ShortProcess.Option(
		timestring.NoSpaces,
	).String(duration)
	fmt.Println(formattedShort) // Output: 1d1h30m
}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

(The internal structure of `LongProcessFormatter` was recently refactored for clarity and maintainability, without altering its public API.)
