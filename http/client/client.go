package client

import "github.com/Danny-Dasilva/CycleTLS/cycletls"

type HTTPClient struct {
	client *cycletls.CycleTLS
}

func NewHTTPClient() *HTTPClient {
	c := cycletls.Init()
	return &HTTPClient{
		client: &c,
	}
}

func (h *HTTPClient) Do(url, method string, options ...func(*cycletls.Options)) (cycletls.Response, error) {
	opts := cycletls.Options{}
	for _, o := range options {
		o(&opts)
	}
	return h.client.Do(
		url,
		opts,
		method,
	)
}
