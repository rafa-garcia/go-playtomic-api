package client

import "fmt"

// APIError represents an error returned by the Playtomic API
type APIError struct {
	StatusCode int
	Message    string
	Details    map[string]interface{}
}

// Error implements the error interface
func (e *APIError) Error() string {
	return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Message)
}
