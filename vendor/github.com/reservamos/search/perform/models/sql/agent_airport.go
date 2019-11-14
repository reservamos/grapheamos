package sql

// AgentAirport relationship between route (origin and destination) and airport
type AgentAirport struct {
	ID        int `gorm:"primary_key" dl:"id"`
	AirportID int `dl:"airport_id"`
	AgentID   int `dl:"agent_id"`
	Airport   Airport
	Agent     Agent
}
