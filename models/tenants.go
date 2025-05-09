package models

// Tenant represents a club/venue in the Playtomic API
type Tenant struct {
	TenantID        string                 `json:"tenant_id"`
	TenantName      string                 `json:"tenant_name"`
	Address         Address                `json:"address"`
	Images          []string               `json:"images"`
	Properties      map[string]interface{} `json:"properties"`
	PlaytomicStatus string                 `json:"playtomic_status"`
}

// Address represents a physical address
type Address struct {
	Street                string     `json:"street"`
	PostalCode            string     `json:"postal_code"`
	City                  string     `json:"city"`
	SubAdministrativeArea string     `json:"sub_administrative_area"`
	AdministrativeArea    string     `json:"administrative_area"`
	Country               string     `json:"country"`
	CountryCode           string     `json:"country_code"`
	Coordinate            Coordinate `json:"coordinate"`
	Timezone              string     `json:"timezone"`
}

// Coordinate represents a geographical coordinate
type Coordinate struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
