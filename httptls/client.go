package httptls

import (
	"net/http"
	"time"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

type Response struct {
	Status    int
	Body      []byte
	Headers   map[string]string
	TimeTaken time.Duration
}

type HTTPClient struct {
	client *cycletls.CycleTLS
}

func NewHTTPClient() *HTTPClient {
	c := cycletls.Init()
	return &HTTPClient{
		client: &c,
	}
}

func (c *HTTPClient) Get(url string, options ...func(*cycletls.Options)) (*Response, error) {
	return c.do(url, http.MethodGet, options...)
}

func (c *HTTPClient) Post(url string, options ...func(*cycletls.Options)) (*Response, error) {
	return c.do(url, http.MethodPost, options...)
}

func (c *HTTPClient) do(url, method string, options ...func(*cycletls.Options)) (*Response, error) {
	opts := cycletls.Options{}
	for _, o := range options {
		o(&opts)
	}

	t := time.Now()
	r, err := c.client.Do(
		url,
		opts,
		method,
	)
	tt := time.Since(t)

	res := &Response{
		Status:    r.Status,
		Body:      []byte(r.Body),
		Headers:   r.Headers,
		TimeTaken: tt,
	}

	if err != nil {
		return res, err
	}
	return res, nil
}
