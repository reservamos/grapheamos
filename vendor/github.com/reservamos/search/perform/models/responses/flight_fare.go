package responses

// FlightFare representation of a segment fare
type FlightFare struct {
	Key             string `json:"key"`
	Service         string `json:"service"`
	ClassOfService  string `json:"class_of_service"`
	ApplicationType string `json:"application_type"`
}
