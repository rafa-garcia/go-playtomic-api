package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafa-garcia/go-playtomic-api/models"
)

func TestGetMatches(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/matches" {
			t.Errorf("Expected path /matches, got %s", r.URL.Path)
		}

		query := r.URL.Query()
		if query.Get("tenant_id") != "test-tenant-id-1,test-tenant-id-2" {
			t.Errorf("Expected tenant_id query param to be 'test-tenant-id-1,test-tenant-id-2', got '%s'", query.Get("tenant_id"))
		}
		if query.Get("has_players") != "true" {
			t.Errorf("Expected has_players query param to be 'true', got '%s'", query.Get("has_players"))
		}
		if query.Get("sport_id") != "PADEL" {
			t.Errorf("Expected sport_id query param to be 'PADEL', got '%s'", query.Get("sport_id"))
		}

		mockResponse := []models.Match{
			{
				MatchID:           "match-123",
				SportID:           "PADEL",
				StartDate:         "2023-01-01T10:00:00",
				EndDate:           "2023-01-01T12:00:00",
				MatchType:         "COMPETITIVE",
				MinPlayersPerTeam: 2,
				MaxPlayersPerTeam: 2,
				MinLevel:          2.5,
				MaxLevel:          4.0,
				Gender:            "MIXED",
				Price:             "30 EUR",
				Teams: []models.Team{
					{
						TeamID:     "0",
						MinPlayers: 2,
						MaxPlayers: 2,
						Players: []models.Player{
							{
								BasePlayer: models.BasePlayer{
									UserID:     "user-123",
									LevelValue: 3.5,
									Picture:    "profile1.jpg",
								},
								Name: "Player One",
							},
						},
					},
					{
						TeamID:     "1",
						MinPlayers: 2,
						MaxPlayers: 2,
						Players: []models.Player{
							{
								BasePlayer: models.BasePlayer{
									UserID:     "user-456",
									LevelValue: 3.0,
									Picture:    "profile2.jpg",
								},
								Name: "Player Two",
							},
						},
					},
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

	params := &models.SearchMatchesParams{
		TenantIDs:  []string{"test-tenant-id-1", "test-tenant-id-2"},
		HasPlayers: true,
		SportID:    "PADEL",
	}

	matches, err := client.GetMatches(context.Background(), params)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(matches) != 1 {
		t.Fatalf("Expected 1 match, got %d", len(matches))
	}

	match := matches[0]
	if match.MatchID != "match-123" {
		t.Errorf("Expected MatchID 'match-123', got %s", match.MatchID)
	}
	if match.MatchType != "COMPETITIVE" {
		t.Errorf("Expected MatchType 'COMPETITIVE', got %s", match.MatchType)
	}
	if match.MinLevel != 2.5 {
		t.Errorf("Expected MinLevel 2.5, got %f", match.MinLevel)
	}
	if match.MaxLevel != 4.0 {
		t.Errorf("Expected MaxLevel 4.0, got %f", match.MaxLevel)
	}

	if len(match.Teams) != 2 {
		t.Fatalf("Expected 2 teams, got %d", len(match.Teams))
	}

	team1 := match.Teams[0]
	if len(team1.Players) != 1 {
		t.Fatalf("Expected 1 player in team 1, got %d", len(team1.Players))
	}

	player1 := team1.Players[0]
	if player1.UserID != "user-123" {
		t.Errorf("Expected player UserID 'user-123', got %s", player1.UserID)
	}
	if player1.Name != "Player One" {
		t.Errorf("Expected player Name 'Player One', got %s", player1.Name)
	}

	if match.Tenant.TenantID != "test-tenant-id-1" {
		t.Errorf("Expected tenant ID 'test-tenant-id-1', got %s", match.Tenant.TenantID)
	}
	if match.Tenant.TenantName != "Test Club" {
		t.Errorf("Expected tenant name 'Test Club', got %s", match.Tenant.TenantName)
	}
}
