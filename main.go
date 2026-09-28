package gopayhero

import (
	"net/http"
	"time"
)

type Option func(*Client)

func WithHttpClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

func NewClient(apiUsername string, apiPassword string, opts ...Option) *Client {
	client := &Client{
		apiUsername: apiUsername,
		apiPassword: apiPassword,
		baseURL:     "https://backend.payhero.co.ke/api/v2",
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(client)
	}
	return client
}
