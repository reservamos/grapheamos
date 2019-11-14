package responses

import (
	"time"

	"github.com/reservamos/search/perform/utils"
)

// BusesResponse is the structure of the response from the integrations service
type BusesResponse struct {
	Status int   `json:"status"`
	Buses  []Bus `json:"body"`
}

// Bus is the structure of each bus trip
type Bus struct {
	Key              string                 `json:"key"`
	Stopovers        int                    `json:"stopovers"`
	Capacity         int                    `json:"capacity"`
	AvailableAll     int                    `json:"available_all"`
	Line             string                 `json:"line"`
	Departure        int64                  `json:"departure"`
	Arrival          int64                  `json:"arrival"`
	Service          string                 `json:"service"`
	InitialDeparture int64                  `json:"initial_departure"`
	SupplierCode     string                 `json:"supplier_code"`
	Fare             string                 `json:"fare"`
	DiscountFare     string                 `json:"discount_fare"`
	Meta             map[string]interface{} `json:"meta"`
	Categories       []Category             `json:"categories_attributes"`
}

// DepartureTime is the time.Time object for the departure of the bus trip
func (b *Bus) DepartureTime() time.Time { return utils.UnixToTime(b.Departure) }

// ArrivalTime is the time.Time object for the arrival of the bus trip
func (b *Bus) ArrivalTime() time.Time { return utils.UnixToTime(b.Arrival) }

// InitialDepartureTime is the time.Time object for the initial departure time of the bus
func (b *Bus) InitialDepartureTime() time.Time { return utils.UnixToTime(b.InitialDeparture) }

// GeneralCategoryTotal returns the general category total of the categories array
func (b *Bus) GeneralCategoryTotal() float64 {
	for _, category := range b.Categories {
		if category.Category == "general" {
			return category.Total
		}
	}
	return -1
}

// Category is part of the Categories array of fares of a bus
type Category struct {
	Category     string                 `json:"category"`
	Availability int                    `json:"availability"`
	Amount       float64                `json:"amount"`
	ServiceFee   float64                `json:"service_fee"`
	Taxes        float64                `json:"taxes"`
	Total        float64                `json:"total"`
	Meta         map[string]interface{} `json:"meta"`
	Discounts    []Discount             `json:"discount_prices_attributes"`
}

// Discount is populated if the trip has a discount
type Discount struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Taxes  float64 `json:"taxes"`
	Total  float64 `json:"total"`
}
