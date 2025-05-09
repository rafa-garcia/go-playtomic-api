package models

// BasePlayer contains the common fields across all player representations
type BasePlayer struct {
	UserID     string  `json:"user_id"`
	LevelValue float64 `json:"level_value"`
	Picture    string  `json:"picture"`
}

// Player represents a player in class/match contexts
type Player struct {
	BasePlayer
	Name                   string  `json:"name"`
	GuestID                *string `json:"guest_id"`
	Email                  *string `json:"email"`
	Gender                 *string `json:"gender"`
	LevelValueConfidence   float64 `json:"level_value_confidence"`
	Phone                  *string `json:"phone"`
	CommunicationsLanguage string  `json:"communications_language"`
	IsPremium              bool    `json:"is_premium"`
	FamilyMemberID         *string `json:"family_member_id"`
}
