package telegram

import (
	"net/http"
	"time"
)

// Option is a function that configures a Client
type Option func(*Client)

// WithTimeout sets the timeout for HTTP requests
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
	}
}

// WithRetries sets the number of retries for failed HTTP requests
func WithRetries(retries int) Option {
	return func(c *Client) {
		c.retries = retries
	}
}

// WithRetryDelay sets the delay between retries
func WithRetryDelay(delay time.Duration) Option {
	return func(c *Client) {
		c.retryDelay = delay
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
