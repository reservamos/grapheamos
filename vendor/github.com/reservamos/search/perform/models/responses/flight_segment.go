package responses

import "time"

// FlightSegment model of web services flight segment response
type FlightSegment struct {
	Key                 string     `json:"key"`
	OriginID            string     `json:"origin_id"`
	OriginTerminal      string     `json:"origin_terminal"`
	DestinationID       string     `json:"destination_id"`
	DestinationTerminal string     `json:"destination_terminal"`
	Departure           string     `json:"departure"`
	Arrival             string     `json:"arrival"`
	Fare                FlightFare `json:"fare"`
	Legs                []Leg      `json:"legs"`
}

// DepartureTime Returns the departure time in CDMX timezone
func (fs *FlightSegment) DepartureTime() time.Time { return servicesTimeToTime(fs.Departure) }

// ArrivalTime Returns the Arrival time in CDMX timezone
func (fs *FlightSegment) ArrivalTime() time.Time { return servicesTimeToTime(fs.Arrival) }
