package core

import (
	"context"
	"net/http"
	"time"
)

type Client struct {
	client  *http.Client
	baseURL string
}

func NewClient(baseURL string, timeout time.Duration) *Client {
	transport := &http.Transport{
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		DisableCompression:    true,
		DisableKeepAlives:     false,
		MaxConnsPerHost:       100,
		ResponseHeaderTimeout: timeout,
	}
	return &Client{
		client: &http.Client{
			Timeout:   timeout,
			Transport: transport,
		},
		baseURL: baseURL,
	}
}

func (c *Client) ForwardRequest(ctx context.Context, opts RequestOptions) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, opts.Method, c.baseURL+opts.Path, opts.Body)
	if err != nil {

		return nil, err
	}

	req.Header = opts.Headers

	if opts.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.Timeout)
		defer cancel()
		req = req.WithContext(ctx)
	}

	return c.client.Do(req)
}
