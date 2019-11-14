package sql

import (
	"fmt"

	"github.com/reservamos/search/internal/config"
	"github.com/reservamos/search/perform/models/sql/utils"
)

// Airport self described model
type Airport struct {
	ID       int     `gorm:"primary_key" dl:"id"`
	Name     string  `dl:"name"`
	IataCode string  `dl:"iata_code"`
	Lat      float64 `dl:"lat"`
	Long     float64 `dl:"long"`
	CityID   int     `dl:"city_id"`
	City     City
	// Only for populated query
	CityName string
	CitySlug string
}

func AirportsByCityID(cityID int) []Airport {
	var airports []Airport
	config.SQL.Raw(fmt.Sprintf(`
	SELECT id, lat, long, city_id
	FROM airports 
	WHERE city_id=%d
	`, cityID)).Scan(&airports)
	return airports
}

// AirportsWithIATACodes returns a slice of aiports given a slice of iata codes
func AirportsWithIATACodes(codes []string) []Airport {
	airports := []Airport{}

	if len(codes) == 0 {
		return airports
	}

	codesForQuery := utils.StringsForWhere(codes)
	query := `
		SELECT airports.*, cities.name city_name, places.slug city_slug
		FROM airports
			INNER JOIN cities ON airports.city_id = cities.id
			INNER JOIN places ON places.findable_type = 'City' AND places.findable_id = cities.id
			WHERE iata_code IN (%s)
	`
	config.SQL.Raw(fmt.Sprintf(query, codesForQuery)).Scan(&airports)
	return airports
}
