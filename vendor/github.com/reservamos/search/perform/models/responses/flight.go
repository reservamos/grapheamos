package responses

import (
	"strconv"
	"time"
)

// Flight this is the base structure for a flight
type Flight struct {
	Key                 string           `json:"key"`
	OriginID            string           `json:"origin_id"`
	OriginTerminal      string           `json:"origin_terminal"`
	DestinationID       string           `json:"destination_id"`
	DestinationTerminal string           `json:"destination_terminal"`
	Departure           string           `json:"departure"`
	Arrival             string           `json:"arrival"`
	FlightDuration      int              `json:"flight_duration"`
	ConnectionDuration  int              `json:"connection_duration"`
	Duration            int              `json:"duration"`
	Availability        int              `json:"availability"`
	EstimatedCostString string           `json:"estimated_cost"`
	Pricing             FlightPricing    `json:"pricing"`
	Segments            []FlightSegment  `json:"segments"`
	Discounts           []FlightDiscount `json:"discounts"`
}

// DepartureTime returns departue time in CDMX time zone
func (f *Flight) DepartureTime() time.Time { return servicesTimeToTime(f.Departure) }

// ArrivalTime returns arrival time in CDMX time zone
func (f *Flight) ArrivalTime() time.Time { return servicesTimeToTime(f.Arrival) }

// EstimatedCost returns the estimated cost
func (f *Flight) EstimatedCost() float64 {
	res, _ := strconv.ParseFloat(f.EstimatedCostString, 64)
	return float64(res)
}
