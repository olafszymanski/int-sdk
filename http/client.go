package http

import (
	"time"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

type Doer interface {
	Do(request *Request) (*Response, error)
}

type Request struct {
	Method  string
	URL     string
	Body    []byte
	Headers map[string]string
	Timeout time.Duration
}

type Response struct {
	Status    int
	Body      []byte
	Headers   map[string]string
	TimeTaken time.Duration
}

type Client struct {
	client *cycletls.CycleTLS
}

func NewClient() *Client {
	c := cycletls.Init()
	return &Client{
		client: &c,
	}
}

func (c *Client) Do(request *Request) (*Response, error) {
	t := time.Now()
	r, err := c.client.Do(
		request.URL,
		getOptions(request),
		request.Method,
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

func getOptions(request *Request) cycletls.Options {
	opts := cycletls.Options{}
	if request.Body != nil {
		opts.Body = string(request.Body)
	}
	if request.Headers != nil {
		opts.Headers = request.Headers
	}
	if request.Timeout != 0 {
		opts.Timeout = int(request.Timeout)
	}
	return opts
}
