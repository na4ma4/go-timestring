package timestring_test

import (
	"testing"
	"time"

	"github.com/na4ma4/go-timestring"
)

func TestLongProcessSimpleTable(t *testing.T) {
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
		itd, err := time.ParseDuration(tc.td)
		if err != nil {
			t.Errorf("unexpected error parsing duration: %s", err)
			continue
		}
		o := timestring.LongProcess.String(itd)
		if o != tc.ex {
			t.Errorf("LongProcess.String() returned invalid duration(%s): %s", itd.String(), o)
		}
	}

}
