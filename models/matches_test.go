package models

import (
	"net/url"
	"testing"
)

func TestSearchMatchesParamsToURLValues(t *testing.T) {
	tests := []struct {
		name     string
		params   SearchMatchesParams
		expected url.Values
	}{
		{
			name:     "Empty params",
			params:   SearchMatchesParams{},
			expected: url.Values{},
		},
		{
			name: "Complete params",
			params: SearchMatchesParams{
				Sort:          "start_date,DESC",
				HasPlayers:    true,
				SportID:       "PADEL",
				TenantIDs:     []string{"tenant-123", "tenant-456"},
				Visibility:    "VISIBLE",
				FromStartDate: "2023-01-01T00:00:00",
				Size:          50,
				Page:          2,
			},
			expected: url.Values{
				"sort":            []string{"start_date,DESC"},
				"has_players":     []string{"true"},
				"sport_id":        []string{"PADEL"},
				"tenant_id":       []string{"tenant-123,tenant-456"},
				"visibility":      []string{"VISIBLE"},
				"from_start_date": []string{"2023-01-01T00:00:00"},
				"size":            []string{"50"},
				"page":            []string{"2"},
			},
		},
		{
			name: "Has players only",
			params: SearchMatchesParams{
				HasPlayers: true,
			},
			expected: url.Values{
				"has_players": []string{"true"},
			},
		},
		{
			name: "Whitespace handling",
			params: SearchMatchesParams{
				Sort:       "  start_date,DESC  ",
				SportID:    "  PADEL  ",
				Visibility: "  VISIBLE  ",
			},
			expected: url.Values{
				"sort":       []string{"start_date,DESC"},
				"sport_id":   []string{"PADEL"},
				"visibility": []string{"VISIBLE"},
			},
		},
		{
			name: "Single tenant ID",
			params: SearchMatchesParams{
				TenantIDs: []string{"tenant-123"},
			},
			expected: url.Values{
				"tenant_id": []string{"tenant-123"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.params.ToURLValues()

			for k, expectedVals := range tt.expected {
				resultVals, ok := result[k]
				if !ok {
					t.Errorf("Expected key %q not found in result", k)
					continue
				}

				if len(resultVals) != len(expectedVals) {
					t.Errorf("Key %q: expected %d values, got %d", k, len(expectedVals), len(resultVals))
					continue
				}

				for i, expectedVal := range expectedVals {
					if resultVals[i] != expectedVal {
						t.Errorf("Key %q, index %d: expected %q, got %q", k, i, expectedVal, resultVals[i])
					}
				}
			}

			for k := range result {
				if _, ok := tt.expected[k]; !ok {
					t.Errorf("Unexpected key %q in result", k)
				}
			}
		})
	}
}
