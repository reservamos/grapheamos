package utils

import (
	"github.com/reservamos/search/internal/config"
	"strings"
)

// TODO: I consider this function should be moved to another mor convinient place

// PlaceType returns city, terminal or airport depending on the slug
func PlaceType(placeSlug string) string {
	placeTypes := []string{}
	config.SQL.Table("places").Where("slug = ?", placeSlug).Pluck("findable_type", &placeTypes)
	if len(placeTypes) > 0 {
		return strings.ToLower(placeTypes[0])
	}
	return "city"
}
