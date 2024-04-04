package httptls

import (
	"time"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

func WithTimeout(timeout time.Duration) func(*cycletls.Options) {
	return func(o *cycletls.Options) {
		o.Timeout = int(timeout)
	}
}
