package models

// Coach represents a coach
type Coach struct {
	UserID                 string  `json:"user_id"`
	LockID                 string  `json:"lock_id"`
	Name                   string  `json:"name"`
	Picture                string  `json:"picture"`
	Email                  *string `json:"email"`
	Gender                 *string `json:"gender"`
	LevelValue             float64 `json:"level_value"`
	LevelValueConfidence   float64 `json:"level_value_confidence"`
	Phone                  *string `json:"phone"`
	CommunicationsLanguage string  `json:"communications_language"`
	IsPremium              bool    `json:"is_premium"`
}
