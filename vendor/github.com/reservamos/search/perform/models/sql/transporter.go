package sql

import "time"

// Transporter holds the structure for the busline group
type Transporter struct {
	ID              int       `gorm:"primary_key" dl:"id"`
	Abbr            string    `dl:"abbr"`
	Name            string    `dl:"name"`
	CreatedAt       time.Time `dl:"created_at"`
	UpdatedAt       time.Time `dl:"updated_at"`
	Ally            bool      `dl:"ally"`
	PurchaseEnabled bool      `dl:"purchase_enabled"`
	Lines           []Line
}

// IsAggregator says whether or not the transporter comes from pricetravel
func (t Transporter) IsAggregator() bool {
	return t.Abbr == "pricetravel"
}
