package responses

// FlightDiscount discount to be used when requesting flight
type FlightDiscount struct {
	Type     string `json:"type"`
	Category string `json:"category"`
	Code     string `json:"code"`
}
