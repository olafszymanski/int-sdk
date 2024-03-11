package client

import "github.com/Danny-Dasilva/CycleTLS/cycletls"

func WithTimeout(timeout int) func(*cycletls.Options) {
	return func(o *cycletls.Options) {
		o.Timeout = timeout
	}
}
