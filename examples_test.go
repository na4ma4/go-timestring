package timestring_test

import (
	"fmt"
	"time"

	"github.com/na4ma4/go-timestring"
)

// ExampleAbsoluteFormatter_String demonstrates the usage of AbsoluteFormatter's String method.
func ExampleAbsoluteFormatter_String() {
	d, _ := time.ParseDuration("49h15m30s")
	fmt.Println(timestring.Absolute.String(d))
	// Output: 2d 1h 15m 30s
}

// ExampleLongProcess_String demonstrates the usage of LongProcessFormatter's String method.
func ExampleLongProcess_String() {
	d, _ := time.ParseDuration("49h15m30s")
	fmt.Println(timestring.LongProcess.String(d))
	// Output: 2 days 1 hour 15 minutes 30 seconds
}

// ExampleShortProcess_String demonstrates the usage of ShortProcessFormatter's String method.
func ExampleShortProcess_String() {
	d, _ := time.ParseDuration("49h15m30s")
	fmt.Println(timestring.ShortProcess.String(d))
	// Output: 2d 1h 15m 30s
}
