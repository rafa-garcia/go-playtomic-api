package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafa-garcia/go-playtomic-api/models"
)

func TestGetLessons(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/lessons" {
			t.Errorf("Expected path /lessons, got %s", r.URL.Path)
		}

		query := r.URL.Query()
		if query.Get("tenant_id") != "test-tenant-id" {
			t.Errorf("Expected tenant_id query param to be 'test-tenant-id', got '%s'", query.Get("tenant_id"))
		}
		if query.Get("tournament_visibility") != "PUBLIC" {
			t.Errorf("Expected tournament_visibility query param to be 'PUBLIC', got '%s'", query.Get("tournament_visibility"))
		}

		mockResponse := []models.Lesson{
			{
				TournamentID:   "lesson-123",
				TournamentName: "Test Lesson",
				StartDate:      "2023-01-01T10:00:00",
				EndDate:        "2023-01-01T12:00:00",
				Type:           "CLASS",
				MinPlayers:     2,
				MaxPlayers:     4,
				RegisteredPlayers: []models.LessonPlayer{
					{
						UserID:                "user-123",
						PaymentID:             "payment-456",
						RegistrationPrice:     "30.00",
						PaymentMethodType:     "CREDIT_CARD",
						FullName:              "John Doe",
						LevelValue:            3.5,
						Picture:               "profile.jpg",
						PaidAtMerchant:        true,
						PaymentB2bBillingType: "INVOICE",
					},
				},
				LevelDescription:     "2.0 - 4.0",
				TournamentVisibility: "PUBLIC",
				TournamentStatus:     "REGISTRATION_OPEN",
				AvailablePlaces:      2,
				Tenant: models.LessonTenant{
					TenantID:   "test-tenant-id",
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

	params := &models.SearchLessonsParams{
		TenantID:             "test-tenant-id",
		TournamentVisibility: "PUBLIC",
	}

	lessons, err := client.GetLessons(context.Background(), params)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(lessons) != 1 {
		t.Fatalf("Expected 1 lesson, got %d", len(lessons))
	}

	lesson := lessons[0]
	if lesson.TournamentID != "lesson-123" {
		t.Errorf("Expected TournamentID 'lesson-123', got %s", lesson.TournamentID)
	}
	if lesson.TournamentName != "Test Lesson" {
		t.Errorf("Expected TournamentName 'Test Lesson', got %s", lesson.TournamentName)
	}
	if lesson.Type != "CLASS" {
		t.Errorf("Expected Type 'CLASS', got %s", lesson.Type)
	}

	if len(lesson.RegisteredPlayers) != 1 {
		t.Fatalf("Expected 1 registered player, got %d", len(lesson.RegisteredPlayers))
	}

	player := lesson.RegisteredPlayers[0]
	if player.UserID != "user-123" {
		t.Errorf("Expected player UserID 'user-123', got %s", player.UserID)
	}
	if player.FullName != "John Doe" {
		t.Errorf("Expected player FullName 'John Doe', got %s", player.FullName)
	}

	if lesson.Tenant.TenantID != "test-tenant-id" {
		t.Errorf("Expected tenant ID 'test-tenant-id', got %s", lesson.Tenant.TenantID)
	}
	if lesson.Tenant.TenantName != "Test Club" {
		t.Errorf("Expected tenant name 'Test Club', got %s", lesson.Tenant.TenantName)
	}
}
