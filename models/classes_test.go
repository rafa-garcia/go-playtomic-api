package models

import (
	"net/url"
	"testing"
)

func TestSearchClassesParamsToURLValues(t *testing.T) {
	tests := []struct {
		name     string
		params   SearchClassesParams
		expected url.Values
	}{
		{
			name:   "Empty params",
			params: SearchClassesParams{},
			expected: url.Values{
				"page": []string{"0"},
			},
		},
		{
			name: "Complete params",
			params: SearchClassesParams{
				Sort:             "start_date,ASC",
				Status:           "PENDING,IN_PROGRESS",
				Type:             "COURSE,PUBLIC",
				TenantIDs:        []string{"tenant-123", "tenant-456"},
				IncludeSummary:   true,
				Size:             50,
				Page:             2,
				CourseVisibility: "PUBLIC",
				FromStartDate:    "2023-01-01T00:00:00",
				Coordinate: &Coordinate{
					Lat: 40.416775,
					Lon: -3.703790,
				},
				Radius: 10,
			},
			expected: url.Values{
				"sort":              []string{"start_date,ASC"},
				"status":            []string{"PENDING,IN_PROGRESS"},
				"type":              []string{"COURSE,PUBLIC"},
				"tenant_id":         []string{"tenant-123,tenant-456"},
				"include_summary":   []string{"true"},
				"size":              []string{"50"},
				"page":              []string{"2"},
				"course_visibility": []string{"PUBLIC"},
				"from_start_date":   []string{"2023-01-01T00:00:00"},
				"coordinate":        []string{"40.416775,-3.703790"},
				"radius":            []string{"10"},
			},
		},
		{
			name: "Coordinate without radius",
			params: SearchClassesParams{
				Coordinate: &Coordinate{
					Lat: 40.416775,
					Lon: -3.703790,
				},
			},
			expected: url.Values{
				"coordinate": []string{"40.416775,-3.703790"},
				"page":       []string{"0"},
			},
		},
		{
			name: "Whitespace handling",
			params: SearchClassesParams{
				Sort:             "  start_date,ASC  ",
				Status:           "  PENDING,IN_PROGRESS  ",
				Type:             "  COURSE,PUBLIC  ",
				CourseVisibility: "  PUBLIC  ",
			},
			expected: url.Values{
				"sort":              []string{"start_date,ASC"},
				"status":            []string{"PENDING,IN_PROGRESS"},
				"type":              []string{"COURSE,PUBLIC"},
				"course_visibility": []string{"PUBLIC"},
				"page":              []string{"0"},
			},
		},
		{
			name: "With tenants but no coordinates",
			params: SearchClassesParams{
				TenantIDs: []string{"tenant-123"},
				Coordinate: &Coordinate{
					Lat: 40.416775,
					Lon: -3.703790,
				},
			},
			expected: url.Values{
				"tenant_id": []string{"tenant-123"},
				"page":      []string{"0"},
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
