package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rafa-garcia/go-playtomic-api/models"
)

// GetMatches retrieves matches from the Playtomic API
func (c *Client) GetMatches(ctx context.Context, params *models.SearchMatchesParams) ([]models.Match, error) {
	var matches []models.Match
	err := c.sendRequest(ctx, http.MethodGet, "/matches", params.ToURLValues().Encode(), nil, &matches)
	if err != nil {
		return nil, fmt.Errorf("fetching matches: %w", err)
	}
	return matches, nil
}
