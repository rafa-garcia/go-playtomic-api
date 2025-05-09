package models

import (
	"fmt"
	"net/url"
	"strings"
)

// Lesson represents a lesson from the Playtomic API
type Lesson struct {
	TournamentID            string         `json:"tournament_id"`
	TournamentName          string         `json:"tournament_name"`
	TournamentImage         *string        `json:"tournament_image"`
	StartDate               string         `json:"start_date"`
	EndDate                 string         `json:"end_date"`
	Type                    string         `json:"type"`
	MinPlayers              int            `json:"min_players"`
	MaxPlayers              int            `json:"max_players"`
	RegisteredPlayers       []LessonPlayer `json:"registered_players"`
	ReservationIDs          interface{}    `json:"reservation_ids"` // Can be null or array
	LevelDescription        string         `json:"level_description"`
	Tags                    []string       `json:"tags"`
	Description             string         `json:"description"`
	PaymentMethodsAllowed   []string       `json:"payment_methods_allowed"`
	Price                   string         `json:"price"`
	SportID                 string         `json:"sport_id"`
	Gender                  string         `json:"gender"`
	RegistrationClosingTime string         `json:"registration_closing_time"`
	IsCancelled             bool           `json:"is_cancelled"`
	TournamentVisibility    string         `json:"tournament_visibility"`
	TournamentStatus        string         `json:"tournament_status"`
	AvailablePlaces         int            `json:"available_places"`
	Tenant                  LessonTenant   `json:"tenant"`
}

// LessonPlayer represents a player registered for a lesson
type LessonPlayer struct {
	UserID                string  `json:"user_id"`
	PaymentID             string  `json:"payment_id"`
	RegistrationPrice     string  `json:"registration_price"`
	PaymentMethodType     string  `json:"payment_method_type"`
	FullName              string  `json:"full_name"`
	LevelValue            float64 `json:"level_value"`
	Picture               string  `json:"picture"`
	PaidAtMerchant        bool    `json:"paid_at_merchant"`
	PaymentB2bBillingType string  `json:"payment_b2b_billing_type"`
}

// LessonTenant represents a club in the lesson context
type LessonTenant struct {
	TenantID      string                 `json:"tenant_id"`
	TenantName    string                 `json:"tenant_name"`
	TenantAddress Address                `json:"tenant_address"`
	TenantImages  []string               `json:"tenant_images"`
	Properties    map[string]interface{} `json:"properties"`
}

// SearchLessonsParams defines parameters for searching lessons
type SearchLessonsParams struct {
	Sort                 string
	TenantID             string // Only accepts a single tenant ID, not a list
	TournamentVisibility string
	Status               string
	Size                 int
	Page                 int
	FromStartDate        string
}

// ToURLValues converts SearchLessonsParams to url.Values
func (p *SearchLessonsParams) ToURLValues() url.Values {
	values := url.Values{}

	if s := strings.TrimSpace(p.Sort); s != "" {
		values.Set("sort", s)
	}

	if id := strings.TrimSpace(p.TenantID); id != "" {
		values.Set("tenant_id", id)
	}

	if v := strings.TrimSpace(p.TournamentVisibility); v != "" {
		values.Set("tournament_visibility", v)
	}

	if s := strings.TrimSpace(p.Status); s != "" {
		values.Set("status", s)
	}

	if p.Size > 0 {
		values.Set("size", fmt.Sprintf("%d", p.Size))
	}

	values.Set("page", fmt.Sprintf("%d", p.Page))

	if p.FromStartDate != "" {
		values.Set("from_start_date", p.FromStartDate)
	}

	return values
}
