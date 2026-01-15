package timestring_test

import (
	"fmt"
	"testing"
	"time"

	ts "github.com/na4ma4/go-timestring"
)

// BenchmarkLongProcessFormatter benchmarks the LongProcessFormatter's String method.
func BenchmarkLongProcessFormatter(b *testing.B) {
	d := 49*time.Hour + 15*time.Minute + 30*time.Second
	for range b.N {
		_ = ts.LongProcess.String(d)
	}
}

func TestLongProcessOptionsCombined(t *testing.T) {
	t.Parallel()

	itd, err := time.ParseDuration("1000h32m29s")
	if err != nil {
		t.Errorf("unexpected error parsing duration: %s", err)

		return
	}

	if o := ts.LongProcess.Option(ts.NoSpaces, ts.NoUnitSpaces).String(itd); o != "41days16hours32minutes29seconds" {
		t.Errorf("LongProcess.Option(NoSpaces, NoUnitSpaces).String() returned invalid duration (%s): %s",
			itd.String(), o)
	}

	if o := ts.LongProcess.Option(ts.Abbreviated, ts.NoSpaces, ts.NoUnitSpaces).String(itd); o != "41d16h32m29s" {
		t.Errorf("LongProcess.Option(Abbreviated, NoSpaces, NoUnitSpaces).String() returned invalid duration (%s): %s",
			itd.String(),
			o,
		)
	}
}

func TestLongProcessNoSpacesContained(t *testing.T) {
	t.Parallel()

	itd, err := time.ParseDuration("1000h32m29s")
	if err != nil {
		t.Errorf("unexpected error parsing duration: %s", err)

		return
	}

	if o := ts.LongProcess.String(itd); o != "41 days 16 hours 32 minutes 29 seconds" {
		t.Errorf("LongProcess.String() returned invalid duration before Option(NoSpaces) (%s): %s", itd.String(), o)
	}

	if o := ts.LongProcess.Option(ts.NoSpaces).String(itd); o != "41 days16 hours32 minutes29 seconds" {
		t.Errorf("LongProcess.Option(NoSpaces).String() returned invalid duration with Option(NoSpaces) (%s): %s",
			itd.String(),
			o,
		)
	}

	if o := ts.LongProcess.Option(ts.NoSpaces, ts.NoUnitSpaces).String(itd); o != "41days16hours32minutes29seconds" {
		t.Errorf("LongProcess.Option(NoSpaces, NoUnitSpaces).String() returned invalid duration (%s): %s",
			itd.String(),
			o,
		)
	}

	if o := ts.LongProcess.String(itd); o != "41 days 16 hours 32 minutes 29 seconds" {
		t.Errorf("LongProcess.String() returned invalid duration after Option(NoSpaces) (%s): %s", itd.String(), o)
	}
}

func TestLongProcessSimpleTable(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		td string
		ex string
	}{
		{"999.9ms", "0 seconds"},
		{"0h0m1s", "1 second"},
		{"0h0m10s", "10 seconds"},
		{"0h0m59s", "59 seconds"},
		{"0h0m60s", "1 minute"},
		{"0h1m2s", "1 minute 2 seconds"},
		{"0h59m59s", "59 minutes 59 seconds"},
		{"0h60m59s", "1 hour 59 seconds"},
		{"0h60m59.999999s", "1 hour 59 seconds"},
		{"1h0m0s", "1 hour"},
		{"1h0m0.001s", "1 hour"},
		{"12h0m0s", "12 hours"},
		{"12h0m0.001s", "12 hours"},
		{"23h0m0s", "23 hours"},
		{"23h0m0.001s", "23 hours"},
		{"24h0m0s", "1 day"},
		{"24h0m0.001s", "1 day"},
		{"25h0m0s", "1 day 1 hour"},
		{"25h0m0.001s", "1 day 1 hour"},
		{"3000h32m29s", "125 days 32 minutes 29 seconds"},
	}
	for _, tc := range tcs {
		t.Run(tc.td, func(t *testing.T) {
			t.Parallel()

			itd, err := time.ParseDuration(tc.td)
			if err != nil {
				t.Errorf("unexpected error parsing duration: %s", err)

				return
			}

			if o := ts.LongProcess.String(itd); o != tc.ex {
				t.Errorf("LongProcess.String() returned invalid duration(%s): %s", itd.String(), o)
			}
		})
	}
}

func TestLongProcessAbbreviatedSimpleTable(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		td string
		ex string
	}{
		{"999.9ms", "0s"},
		{"0h0m1s", "1s"},
		{"0h0m10s", "10s"},
		{"0h0m59s", "59s"},
		{"0h0m60s", "1m"},
		{"0h1m2s", "1m 2s"},
		{"0h59m59s", "59m 59s"},
		{"0h60m59s", "1h 59s"},
		{"0h60m59.999999s", "1h 59s"},
		{"1h0m0s", "1h"},
		{"1h0m0.001s", "1h"},
		{"12h0m0s", "12h"},
		{"12h0m0.001s", "12h"},
		{"23h0m0s", "23h"},
		{"23h0m0.001s", "23h"},
		{"24h0m0s", "1d"},
		{"24h0m0.001s", "1d"},
		{"25h0m0s", "1d 1h"},
		{"25h0m0.001s", "1d 1h"},
		{"3000h32m29s", "125d 32m 29s"},
	}
	for _, tc := range tcs {
		t.Run(tc.td, func(t *testing.T) {
			t.Parallel()

			itd, err := time.ParseDuration(tc.td)
			if err != nil {
				t.Errorf("unexpected error parsing duration: %s", err)

				return
			}

			if o := ts.LongProcess.Option(ts.Abbreviated).String(itd); o != tc.ex {
				t.Errorf("LongProcess.Option(Abbreviated).String() returned invalid duration(%s): %s", itd.String(), o)
			}
		})
	}
}

func TestLongProcessAbbreviatedSimpleNoSpacesTable(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		td string
		ex string
	}{
		{"999.9ms", "0s"},
		{"0h0m1s", "1s"},
		{"0h0m10s", "10s"},
		{"0h0m59s", "59s"},
		{"0h0m60s", "1m"},
		{"0h1m2s", "1m2s"},
		{"0h59m59s", "59m59s"},
		{"0h60m59s", "1h59s"},
		{"0h60m59.999999s", "1h59s"},
		{"1h0m0s", "1h"},
		{"1h0m0.001s", "1h"},
		{"12h0m0s", "12h"},
		{"12h0m0.001s", "12h"},
		{"23h0m0s", "23h"},
		{"23h0m0.001s", "23h"},
		{"24h0m0s", "1d"},
		{"24h0m0.001s", "1d"},
		{"25h0m0s", "1d1h"},
		{"25h0m0.001s", "1d1h"},
		{"3000h32m29s", "125d32m29s"},
	}
	for _, tc := range tcs {
		t.Run(tc.td, func(t *testing.T) {
			t.Parallel()

			itd, err := time.ParseDuration(tc.td)
			if err != nil {
				t.Errorf("unexpected error parsing duration: %s", err)

				return
			}

			if o := ts.LongProcess.Option(ts.Abbreviated).Option(ts.NoSpaces).String(itd); o != tc.ex {
				t.Errorf("LongProcess.Option(Abbreviated).Option(NoSpaces).String() returned invalid duration(%s): %s",
					itd.String(),
					o,
				)
			}

			if o := ts.LongProcess.Option(ts.Abbreviated, ts.NoSpaces).String(itd); o != tc.ex {
				t.Errorf("LongProcess.Option(Abbreviated, NoSpaces).String() returned invalid duration(%s): %s",
					itd.String(), o)
			}
		})
	}
}

func TestLongProcessOptionTable(t *testing.T) {
	t.Parallel()

	oa := ts.Abbreviated
	ons := ts.NoSpaces
	osm := ts.ShowMSOnSeconds
	onu := ts.NoUnitSpaces

	tcs := []struct {
		td   string
		ex   string
		opts []ts.FormatterOption
	}{
		{"999.9ms", "999ms", []ts.FormatterOption{oa, ons, osm}},
		{"10s999.9ms", "10s999ms", []ts.FormatterOption{oa, ons, osm}},
		{"59s999.9ms", "59s999ms", []ts.FormatterOption{oa, ons, osm}},
		{"60s999.9ms", "1m", []ts.FormatterOption{oa, ons, osm}},
		{"61s999.9ms", "1m1s", []ts.FormatterOption{oa, ons, osm}},
		{"999.9ms", "999ms", []ts.FormatterOption{oa, ons, osm}},
		{"0h0m1s", "1s", []ts.FormatterOption{oa, ons, osm}},
		{"0h0m10s", "10s", []ts.FormatterOption{oa, ons, osm}},
		{"0h0m59s", "59s", []ts.FormatterOption{oa, ons, osm}},
		{"0h0m60s", "1m", []ts.FormatterOption{oa, ons, osm}},
		{"0h1m2s", "1m2s", []ts.FormatterOption{oa, ons, osm}},
		{"0h59m59s", "59m59s", []ts.FormatterOption{oa, ons, osm}},
		{"0h60m59s", "1h59s", []ts.FormatterOption{oa, ons, osm}},
		{"0h60m59.999999s", "1h59s", []ts.FormatterOption{oa, ons, osm}},
		{"1h0m0s", "1h", []ts.FormatterOption{oa, ons, osm}},
		{"1h0m0.001s", "1h", []ts.FormatterOption{oa, ons, osm}},
		{"12h0m0s", "12h", []ts.FormatterOption{oa, ons, osm}},
		{"12h0m0.001s", "12h", []ts.FormatterOption{oa, ons, osm}},
		{"23h0m0s", "23h", []ts.FormatterOption{oa, ons, osm}},
		{"23h0m0.001s", "23h", []ts.FormatterOption{oa, ons, osm}},
		{"24h0m0s", "1d", []ts.FormatterOption{oa, ons, osm}},
		{"24h0m0.001s", "1d", []ts.FormatterOption{oa, ons, osm}},
		{"25h0m0s", "1d1h", []ts.FormatterOption{oa, ons, osm}},
		{"25h0m0.001s", "1d1h", []ts.FormatterOption{oa, ons, osm}},
		{"3000h32m29s", "125d32m29s", []ts.FormatterOption{oa, ons, osm}},
		{"0s", "0s", []ts.FormatterOption{oa, ons, osm}},
		{"0s", "0 seconds", []ts.FormatterOption{ons, osm}},
		{"0s", "0 seconds", []ts.FormatterOption{osm}},
		{"0s", "0s", []ts.FormatterOption{oa, osm}},
		{"0s", "0s", []ts.FormatterOption{oa, ons}},
		{"0s", "0s", []ts.FormatterOption{oa}},
		{"0s", "0 seconds", []ts.FormatterOption{ons}},
		{"0s", "0 seconds", []ts.FormatterOption{}},
		{"0s", "0seconds", []ts.FormatterOption{onu}},
		{"3000h32m29s", "125 days 32 minutes 29 seconds", []ts.FormatterOption{}},
		{"3000h32m29s", "125days32minutes29seconds", []ts.FormatterOption{ons, onu}},
	}
	for _, tc := range tcs {
		t.Run(tc.td, func(t *testing.T) {
			t.Parallel()

			itd, err := time.ParseDuration(tc.td)
			if err != nil {
				t.Errorf("unexpected error parsing duration: %s", err)

				return
			}

			if o := ts.LongProcess.Option(tc.opts...).String(itd); o != tc.ex {
				t.Errorf("LongProcess.Option(%s).String() returned invalid duration(%s): expected(%s) got(%s)",
					fmt.Sprintf("%q", tc.opts),
					itd.String(),
					tc.ex,
					o,
				)
			}
		})
	}
}
