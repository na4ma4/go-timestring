package timestring_test

import (
	"testing"
	"time"

	"github.com/na4ma4/go-timestring"
)

func TestLongProcessAbvSimpleTable(t *testing.T) {
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
		itd, err := time.ParseDuration(tc.td)
		if err != nil {
			t.Errorf("unexpected error parsing duration: %s", err)
			continue
		}
		o := timestring.LongProcessAbv.String(itd)
		if o != tc.ex {
			t.Errorf("LongProcessAbv.String() returned invalid duration(%s): %s", itd.String(), o)
		}
	}

}
