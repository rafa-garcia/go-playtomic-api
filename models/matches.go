package models

import (
	"fmt"
	"net/url"
	"strings"
)

// Match represents a match from the Playtomic API
type Match struct {
	MatchID                       string             `json:"match_id"`
	ReservationID                 *string            `json:"reservation_id"`
	RecurringMatchConfigurationID *string            `json:"recurring_match_configuration_id"`
	Location                      string             `json:"location"`
	SportID                       string             `json:"sport_id"`
	Teams                         []Team             `json:"teams"`
	MinPlayersPerTeam             int                `json:"min_players_per_team"`
	MaxPlayersPerTeam             int                `json:"max_players_per_team"`
	OwnerID                       *string            `json:"owner_id"`
	Status                        string             `json:"status"`
	GameStatus                    string             `json:"game_status"`
	StartDate                     string             `json:"start_date"`
	EndDate                       string             `json:"end_date"`
	Tenant                        Tenant             `json:"tenant"`
	LocationInfo                  LocationInfo       `json:"location_info"`
	MatchType                     string             `json:"match_type"`
	MatchOrganization             string             `json:"match_organization"`
	CompetitionMode               string             `json:"competition_mode"`
	Gender                        string             `json:"gender"`
	MaxLevel                      float64            `json:"max_level"`
	MinLevel                      float64            `json:"min_level"`
	Price                         string             `json:"price"`
	PaymentRequired               bool               `json:"payment_required"`
	ResourceProperties            ResourceProperties `json:"resource_properties"`
	RegistrationInfo              RegistrationInfo   `json:"registration_info"`
	MatchOrigin                   string             `json:"match_origin"`
	RegistrationType              string             `json:"registration_type"`
	RegistrationStatus            string             `json:"registration_status"`
	IsPremium                     bool               `json:"is_premium"`
	IsBooked                      bool               `json:"is_booked"`
	CreatedAt                     string             `json:"created_at"`
	Visibility                    string             `json:"visibility"`
}

// LocationInfo represents information about the location of a match
type LocationInfo struct {
	ID      string   `json:"id"`
	Type    string   `json:"type"`
	Name    string   `json:"name"`
	Address Address  `json:"address"`
	Images  []string `json:"images"`
}

// ResourceProperties contains properties of a resource used for a match
type ResourceProperties struct {
	ResourceType    string `json:"resource_type"`
	ResourceSize    string `json:"resource_size"`
	ResourceFeature string `json:"resource_feature"`
}

// SearchMatchesParams defines parameters for searching matches
type SearchMatchesParams struct {
	Sort          string
	HasPlayers    bool
	SportID       string
	TenantIDs     []string
	Visibility    string
	FromStartDate string
	Size          int
	Page          int
}

// ToURLValues converts SearchMatchesParams to url.Values
func (p *SearchMatchesParams) ToURLValues() url.Values {
	values := url.Values{}

	if s := strings.TrimSpace(p.Sort); s != "" {
		values.Set("sort", s)
	}

	if p.HasPlayers {
		values.Set("has_players", "true")
	}

	if s := strings.TrimSpace(p.SportID); s != "" {
		values.Set("sport_id", s)
	}

	if len(p.TenantIDs) > 0 {
		values.Set("tenant_id", strings.Join(p.TenantIDs, ","))
	}

	if v := strings.TrimSpace(p.Visibility); v != "" {
		values.Set("visibility", v)
	}

	if p.FromStartDate != "" {
		values.Set("from_start_date", p.FromStartDate)
	}

	if p.Size > 0 {
		values.Set("size", fmt.Sprintf("%d", p.Size))
	}

	if p.Page > 0 {
		values.Set("page", fmt.Sprintf("%d", p.Page))
	}

	return values
}
