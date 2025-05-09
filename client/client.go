// Package client provides a Go client for accessing the Playtomic API.
package client

import (
	"net/http"
	"time"
)

const (
	// DefaultBaseURL is the default Playtomic API endpoint
	DefaultBaseURL = "https://api.playtomic.io/v1"

	// DefaultTimeout is the default client timeout
	DefaultTimeout = 30 * time.Second

	// DefaultMaxRetries is the default number of request retries
	DefaultMaxRetries = 3

	// DefaultUserAgent is the default User-Agent sent with requests
	DefaultUserAgent = "PlaytomicGoClient/1.0"
)

// Client provides access to the Playtomic API
type Client struct {
	httpClient *http.Client
	baseURL    string
	userAgent  string
	maxRetries int
	debug      bool
}

// NewClient creates a new Playtomic API client with the given options
func NewClient(opts ...Option) *Client {
	c := &Client{
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		baseURL:    DefaultBaseURL,
		userAgent:  DefaultUserAgent,
		maxRetries: DefaultMaxRetries,
	}

	// Apply options
	for _, opt := range opts {
		opt(c)
	}

	return c
}
