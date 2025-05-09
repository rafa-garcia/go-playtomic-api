package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafa-garcia/go-playtomic-api/models"
)

func TestGetClasses(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/classes" {
			t.Errorf("Expected path /classes, got %s", r.URL.Path)
		}

		query := r.URL.Query()
		if query.Get("tenant_id") != "test-tenant-id-1,test-tenant-id-2" {
			t.Errorf("Expected tenant_id query param to be 'test-tenant-id-1,test-tenant-id-2', got '%s'", query.Get("tenant_id"))
		}
		if query.Get("include_summary") != "true" {
			t.Errorf("Expected include_summary query param to be 'true', got '%s'", query.Get("include_summary"))
		}

		mockResponse := []models.Class{
			{
				AcademyClassID: "class-123",
				SportID:        "PADEL",
				StartDate:      "2023-01-01T10:00:00",
				EndDate:        "2023-01-01T12:00:00",
				Type:           "COURSE",
				CourseSummary: &models.CourseSummary{
					CourseID:   "course-456",
					Name:       "Test Class",
					Gender:     "MIXED",
					Visibility: "PUBLIC",
					MinPlayers: 2,
					MaxPlayers: 4,
				},
				Resource: models.Resource{
					ID:   "resource-789",
					Name: "Court 1",
				},
				Tenant: models.Tenant{
					TenantID:   "test-tenant-id-1",
					TenantName: "Test Club",
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	client := NewClient(
		WithBaseURL(server.URL),
	)

	params := &models.SearchClassesParams{
		TenantIDs:      []string{"test-tenant-id-1", "test-tenant-id-2"},
		IncludeSummary: true,
	}

	classes, err := client.GetClasses(context.Background(), params)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(classes) != 1 {
		t.Fatalf("Expected 1 class, got %d", len(classes))
	}

	class := classes[0]
	if class.AcademyClassID != "class-123" {
		t.Errorf("Expected AcademyClassID 'class-123', got %s", class.AcademyClassID)
	}
	if class.Type != "COURSE" {
		t.Errorf("Expected Type 'COURSE', got %s", class.Type)
	}

	if class.CourseSummary == nil {
		t.Fatalf("Expected course summary, got nil")
	}
	if class.CourseSummary.Name != "Test Class" {
		t.Errorf("Expected course name 'Test Class', got %s", class.CourseSummary.Name)
	}

	if class.Resource.Name != "Court 1" {
		t.Errorf("Expected resource name 'Court 1', got %s", class.Resource.Name)
	}

	if class.Tenant.TenantID != "test-tenant-id-1" {
		t.Errorf("Expected tenant ID 'test-tenant-id-1', got %s", class.Tenant.TenantID)
	}
	if class.Tenant.TenantName != "Test Club" {
		t.Errorf("Expected tenant name 'Test Club', got %s", class.Tenant.TenantName)
	}
}
