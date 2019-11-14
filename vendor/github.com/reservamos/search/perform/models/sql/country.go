package sql

import "time"

// Country is the largest terrytory division in Reservamos
type Country struct {
	ID        int       `gorm:"primary_key" dl:"id"`
	Name      string    `dl:"name"`
	Slug      string    `dl:"slug"`
	CreatedAt time.Time `dl:"created_at"`
	UpdatedAt time.Time `dl:"name"`
}
