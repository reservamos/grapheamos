package responses

import (
	"time"

	"github.com/reservamos/search/internal/config"
	"github.com/reservamos/search/perform/utils"
)

type legFlight struct {
	Number  string `json:"number"`
	Carrier string `json:"carrier"`
}

// Leg is one part of a Segment
type Leg struct {
	OriginID            string    `json:"origin_id"`
	OriginTerminal      string    `json:"origin_terminal"`
	DestinationID       string    `json:"destination_id"`
	DestinationTerminal string    `json:"destination_terminal"`
	Departure           string    `json:"departure"`
	Arrival             string    `json:"arrival"`
	Duration            int       `json:"duration"`
	ConnectionDuration  int       `json:"connection_duration"`
	LegFlight           legFlight `json:"flight"`
}

// DepartureTime Returns the departure time in CDMX timezone
func (l *Leg) DepartureTime() time.Time { return servicesTimeToTime(l.Departure) }

// ArrivalTime Returns the Arrival time in CDMX timezone
func (l *Leg) ArrivalTime() time.Time { return servicesTimeToTime(l.Arrival) }

// servicesTimeToTime receives a date string to parse it to time.Times
func servicesTimeToTime(s string) time.Time {
	parsed, _ := time.Parse("2006-01-02T15:04:05.999 MST", s+" UTC")
	return utils.GetTimeTz(parsed, config.App.Timezone)
}
