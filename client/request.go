package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// sendRequest sends a request to the Playtomic API and decodes the response
func (c *Client) sendRequest(ctx context.Context, method, endpoint string, queryParams string, body io.Reader, result interface{}) error {
	reqURL := fmt.Sprintf("%s%s?%s", c.baseURL, endpoint, queryParams)

	var resp *http.Response
	var err error
	var req *http.Request

	// Create the request
	req, err = http.NewRequestWithContext(ctx, method, reqURL, body)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)

	// Try the request with retries
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		resp, err = c.httpClient.Do(req)
		if err == nil {
			break
		}

		// If this was the last attempt, return the error
		if attempt == c.maxRetries {
			return fmt.Errorf("sending request after %d attempts: %w", c.maxRetries, err)
		}

		// Wait before retrying (could implement exponential backoff)
		select {
		case <-time.After(time.Duration(attempt) * 500 * time.Millisecond):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	// Handle non-200 responses
	if resp.StatusCode != http.StatusOK {
		var apiErr struct {
			Error   string                 `json:"error"`
			Details map[string]interface{} `json:"details"`
		}

		if err := json.Unmarshal(respBody, &apiErr); err == nil && apiErr.Error != "" {
			return &APIError{
				StatusCode: resp.StatusCode,
				Message:    apiErr.Error,
				Details:    apiErr.Details,
			}
		}

		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    "Unexpected response from API",
		}
	}

	// Decode into result
	if err := json.Unmarshal(respBody, result); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	return nil
}
