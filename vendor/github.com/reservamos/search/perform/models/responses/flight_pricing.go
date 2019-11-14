package responses

import "strconv"

// FlightPricing model of web services flight pricing response
type FlightPricing struct {
	AmountString   string `json:"amount"`
	TaxesString    string `json:"taxes"`
	DiscountString string `json:"discount"`
	TotalString    string `json:"total"`
}

// Amount float64 version of response amount
func (fp *FlightPricing) Amount() float64 {
	res, _ := strconv.ParseFloat(fp.AmountString, 64)
	return float64(res)
}

// Taxes float64 version of response taxes
func (fp *FlightPricing) Taxes() float64 {
	res, _ := strconv.ParseFloat(fp.TaxesString, 64)
	return float64(res)
}

// Total float64 version of response total
func (fp *FlightPricing) Total() float64 {
	res, _ := strconv.ParseFloat(fp.TotalString, 64)
	return float64(res)
}

// Discount float64 version of response discount
func (fp *FlightPricing) Discount() float64 {
	res, _ := strconv.ParseFloat(fp.DiscountString, 64)
	return float64(res)
}
