package sql

import (
	"time"

	"github.com/reservamos/search/internal/config"
)

// Place integrates cities, terminals and airports into one slug
type Place struct {
	ID           int `gorm:"primary_key"`
	Slug         string
	FindableID   int
	FindableType string
	Hits         int
	Priority     int
	Popularity   float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// PlaceBySlug finds a place using the db slug
func PlaceBySlug(slug string) (Place, error) {
	var place Place
	result := config.SQL.First(&place, "slug = ?", slug)
	return place, result.Error
}
