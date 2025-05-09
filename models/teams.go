package models

// Team represents a team in a match
type Team struct {
	TeamID     string   `json:"team_id"`
	Players    []Player `json:"players"`
	MinPlayers int      `json:"min_players"`
	MaxPlayers int      `json:"max_players"`
	TeamResult *string  `json:"team_result"`
}
