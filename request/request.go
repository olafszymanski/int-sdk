package request

import (
	"fmt"
	"net/http"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

var (
	ErrRequest              = fmt.Errorf("request failed")
	ErrUnexpectedStatusCode = fmt.Errorf("unexpected status code")
)

func Do(client *cycletls.CycleTLS, url, method string, timeout int) (*cycletls.Response, error) {
	res, err := client.Do(
		url,
		cycletls.Options{
			Timeout: timeout,
		},
		method,
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrRequest, err)
	}
	if res.Status != http.StatusOK {
		return nil, fmt.Errorf("%w: %d", ErrUnexpectedStatusCode, res.Status)
	}
	return &res, nil
}
