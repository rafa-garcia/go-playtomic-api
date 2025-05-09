package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rafa-garcia/go-playtomic-api/models"
)

// GetClasses retrieves classes from the Playtomic API
func (c *Client) GetClasses(ctx context.Context, params *models.SearchClassesParams) ([]models.Class, error) {
	var classes []models.Class
	err := c.sendRequest(ctx, http.MethodGet, "/classes", params.ToURLValues().Encode(), nil, &classes)
	if err != nil {
		return nil, fmt.Errorf("fetching classes: %w", err)
	}
	return classes, nil
}
