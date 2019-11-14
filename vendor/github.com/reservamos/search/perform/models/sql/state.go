package sql

import "time"

// State is the largest territory division inside of a country
type State struct {
	ID        int       `gorm:"primary_key"  dl:"id"`
	CountryID int       `dl:"country_id"`
	Name      string    `dl:"name"`
	Slug      string    `dl:"slug"`
	CreatedAt time.Time `dl:"created_at"`
	UpdatedAt time.Time `dl:"updated_at"`
	Country   Country
}
