package models

// Resource represents a court or other resource
type Resource struct {
	ID         string                 `json:"id"`
	LockID     string                 `json:"lock_id"`
	Name       string                 `json:"name"`
	Properties map[string]interface{} `json:"properties"`
}
