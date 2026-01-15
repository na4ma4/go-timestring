package timestring_test

import (
	"testing"
	"time"

	ts "github.com/na4ma4/go-timestring"
)

// func parseDurationForTest(t *testing.T, s string) time.Duration {
// 	t.Helper()
// 	d, err := time.ParseDuration(s)
// 	if err != nil {
// 		t.Fatalf("Failed to parse duration '%s': %v", s, err)
// 	}
// 	return d
// }

// BenchmarkAbsoluteFormatter benchmarks the AbsoluteFormatter's String method.
func BenchmarkAbsoluteFormatter(b *testing.B) {
	d := 49*time.Hour + 15*time.Minute + 30*time.Second
	for range b.N {
		_ = ts.Absolute.String(d)
	}
}

func TestAbsoluteFormatter_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		duration time.Duration
		options  []ts.FormatterOption
		expected string
	}{
		{
			name:     "Zero duration",
			duration: 0,
			expected: "0s",
		},
		{
			name:     "Only milliseconds",
			duration: 500 * time.Millisecond,
			expected: "500ms",
		},
		{
			name:     "Only seconds",
			duration: 5 * time.Second,
			expected: "5s",
		},
		{
			name:     "Seconds and milliseconds",
			duration: 5*time.Second + 500*time.Millisecond,
			expected: "5s 500ms",
		},
		{
			name:     "Only minutes",
			duration: 10 * time.Minute,
			expected: "10m",
		},
		{
			name:     "Minutes and seconds",
			duration: 10*time.Minute + 5*time.Second,
			expected: "10m 5s",
		},
		{
			name:     "Only hours",
			duration: 2 * time.Hour,
			expected: "2h",
		},
		{
			name:     "Hours and minutes",
			duration: 2*time.Hour + 10*time.Minute,
			expected: "2h 10m",
		},
		{
			name:     "Only days",
			duration: 3 * 24 * time.Hour, // 3 days
			expected: "3d",
		},
		{
			name:     "Days plus ns",
			duration: 3*24*time.Hour + 100*time.Nanosecond, // 3 days plus 100ns
			expected: "3d 100ns",
		},
		{
			name:     "Days and hours",
			duration: (3*24*time.Hour + 2*time.Hour), // 3 days 2 hours
			expected: "3d 2h",
		},
		{
			name:     "All units",
			duration: (1*24*time.Hour + 2*time.Hour + 3*time.Minute + 4*time.Second + 5*time.Millisecond), // 1d 2h 3m 4s 5ms
			expected: "1d 2h 3m 4s 5ms",
		},
		{
			name:     "Complex duration with all parts",
			duration: (1*24*time.Hour + 2*time.Hour + 3*time.Minute + 4*time.Second + 5*time.Millisecond),
			expected: "1d 2h 3m 4s 5ms",
		},
		{
			name:     "Less than 1µs but non-zero",
			duration: 100 * time.Nanosecond,
			expected: "100ns",
		},
		{
			name:     "Less than 1ms but non-zero",
			duration: 100 * time.Microsecond,
			expected: "100µs",
		},

		// --- Options ---
		{
			name:     "NoSpaces option",
			duration: (1*24*time.Hour + 2*time.Hour + 3*time.Minute), // 1d 2h 3m
			options:  []ts.FormatterOption{ts.NoSpaces},
			expected: "1d2h3m",
		},
		{
			name:     "NoUnitSpaces option (should have no effect as abbreviated units don't have internal spaces)",
			duration: (1*24*time.Hour + 2*time.Hour + 3*time.Minute), // 1d 2h 3m
			options:  []ts.FormatterOption{ts.NoUnitSpaces},
			expected: "1d 2h 3m",
		},
		{
			name:     "NoSpaces and NoUnitSpaces",
			duration: (1*24*time.Hour + 2*time.Hour + 3*time.Minute), // 1d 2h 3m
			options:  []ts.FormatterOption{ts.NoSpaces, ts.NoUnitSpaces},
			expected: "1d2h3m",
		},
		{
			name:     "Zero duration with NoSpaces",
			duration: 0,
			options:  []ts.FormatterOption{ts.NoSpaces},
			expected: "0s",
		},
		{
			name:     "Seconds and ms with NoSpaces",
			duration: 5*time.Second + 500*time.Millisecond,
			options:  []ts.FormatterOption{ts.NoSpaces},
			expected: "5s500ms",
		},
		{
			name:     "ShowMSOnSeconds (should have no effect on AbsoluteFormatter)",
			duration: 5 * time.Second,
			options:  []ts.FormatterOption{ts.ShowMSOnSeconds},
			expected: "5s",
		},
		{
			name:     "Abbreviated (already default, should have no effect)",
			duration: 5 * time.Second,
			options:  []ts.FormatterOption{ts.Abbreviated},
			expected: "5s",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			formatter := ts.Absolute
			if len(tc.options) > 0 {
				formatter = formatter.Option(tc.options...)
			}
			result := formatter.String(tc.duration)
			if result != tc.expected {
				t.Errorf("Expected '%s', but got '%s' for duration %v with options %v",
					tc.expected, result, tc.duration, tc.options,
				)
			}
		})
	}
}
