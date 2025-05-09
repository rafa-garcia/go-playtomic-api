package models

import (
	"fmt"
	"net/url"
	"strings"
)

// Class represents a class from the Playtomic API
type Class struct {
	Type             string           `json:"type"`
	AcademyClassID   string           `json:"academy_class_id"`
	SportID          string           `json:"sport_id"`
	Tenant           Tenant           `json:"tenant"`
	Resource         Resource         `json:"resource"`
	StartDate        string           `json:"start_date"`
	EndDate          string           `json:"end_date"`
	Coaches          []Coach          `json:"coaches"`
	RegistrationInfo RegistrationInfo `json:"registration_info"`
	CourseSummary    *CourseSummary   `json:"course_summary,omitempty"`
	AccessCode       *string          `json:"access_code"`
	Origin           string           `json:"origin"`
	IsCanceled       bool             `json:"is_canceled"`
	PrivateNotes     *string          `json:"private_notes"`
	PublicNotes      string           `json:"public_notes"`
	Status           string           `json:"status"`
	PaymentStatus    string           `json:"payment_status"`
}

// CourseSummary represents summary information about a course
type CourseSummary struct {
	CourseID   string `json:"course_id"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Visibility string `json:"visibility"`
	MinPlayers int    `json:"min_players"`
	MaxPlayers int    `json:"max_players"`
}

// SearchClassesParams defines parameters for searching classes
type SearchClassesParams struct {
	Sort             string
	Status           string
	Type             string
	TenantIDs        []string
	IncludeSummary   bool
	Size             int
	Page             int
	CourseVisibility string
	FromStartDate    string
	Coordinate       *Coordinate
	Radius           int
}

// ToURLValues converts SearchClassesParams to url.Values
func (p *SearchClassesParams) ToURLValues() url.Values {
	values := url.Values{}

	if s := strings.TrimSpace(p.Sort); s != "" {
		values.Set("sort", s)
	}

	if s := strings.TrimSpace(p.Status); s != "" {
		values.Set("status", s)
	}

	if t := strings.TrimSpace(p.Type); t != "" {
		values.Set("type", t)
	}

	if len(p.TenantIDs) > 0 {
		values.Set("tenant_id", strings.Join(p.TenantIDs, ","))
	}

	if p.IncludeSummary {
		values.Set("include_summary", "true")
	}

	if p.Size > 0 {
		values.Set("size", fmt.Sprintf("%d", p.Size))
	}

	values.Set("page", fmt.Sprintf("%d", p.Page))

	if cv := strings.TrimSpace(p.CourseVisibility); cv != "" {
		values.Set("course_visibility", cv)
	}

	if p.FromStartDate != "" {
		values.Set("from_start_date", p.FromStartDate)
	}

	if p.Coordinate != nil && len(p.TenantIDs) == 0 {
		values.Set("coordinate", fmt.Sprintf("%f,%f", p.Coordinate.Lat, p.Coordinate.Lon))

		if p.Radius > 0 {
			values.Set("radius", fmt.Sprintf("%d", p.Radius))
		}
	}

	return values
}
