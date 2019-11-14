package responses

// FlightsResponse  this is the base structure that we will receive from the service
type FlightsResponse struct {
	Status  int      `json:"status"`
	Flights []Flight `json:"body"`
}
