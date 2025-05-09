package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rafa-garcia/go-playtomic-api/models"
)

// GetLessons retrieves lessons from the Playtomic API
func (c *Client) GetLessons(ctx context.Context, params *models.SearchLessonsParams) ([]models.Lesson, error) {
	var lessons []models.Lesson
	err := c.sendRequest(ctx, http.MethodGet, "/lessons", params.ToURLValues().Encode(), nil, &lessons)
	if err != nil {
		return nil, fmt.Errorf("fetching lessons: %w", err)
	}
	return lessons, nil
}
