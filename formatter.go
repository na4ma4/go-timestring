package timestring

import (
	"time"
)

type Formatter interface {
	String(time.Duration) string
}
