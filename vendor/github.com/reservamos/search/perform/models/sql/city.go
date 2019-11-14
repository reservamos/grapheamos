package sql

import (
	"fmt"
	"strings"
	"time"

	"github.com/reservamos/search/internal/config"
	"github.com/reservamos/search/perform/models/sql/utils"
)

// City is part of a state, a city holds Terminals
type City struct {
	ID         int       `gorm:"primary_key" dl:"id"`
	Name       string    `dl:"name"`
	ASCIIName  string    `dl:"ascii_name"`
	TimeZoneID string    `dl:"time_zone_id"`
	StateID    int       `dl:"state_id"`
	CreatedAt  time.Time `dl:"created_at"`
	UpdatedAt  time.Time `dl:"updated_at"`
	Slug       string    `dl:"slug"`
	Place      Place     `gorm:"polymorphic:Findable;polymorphic_value:City"`
	State      State
	// Only for prepopulate
	StateName   string
	StateAbbr   string
	CountryName string
	CountryAbbr string
}

// CityFindAllWithSlugs finds all cities with given slugs
func CityFindAllWithSlugs(slugs []string) []City {

	citySlugs := []string{}
	citiesFromCitySlugs := []City{}
	terminalSlugs := []string{}
	citiesFromTerminalSlugs := []City{}
	airportSlugs := []string{}
	citiesFromAirportSlugs := []City{}

	for _, slug := range slugs {
		if strings.HasPrefix(slug, "t-") {
			terminalSlugs = append(terminalSlugs, slug)
		} else {
			if strings.HasPrefix(slug, "a-") {
				airportSlugs = append(airportSlugs, slug)
			} else {
				citySlugs = append(citySlugs, slug)
			}
		}
	}

	if len(citySlugs) > 0 {
		slugsSQL := utils.StringsForWhere(citySlugs)
		whereQuery := fmt.Sprintf(`WHERE places.slug IN (%s)`, slugsSQL)
		config.SQL.Raw(fmt.Sprintf(populatedCityFromCityQuery, whereQuery)).Scan(&citiesFromCitySlugs)
	}

	if len(terminalSlugs) > 0 {
		slugsSQL := utils.StringsForWhere(terminalSlugs)
		whereQuery := fmt.Sprintf(`WHERE places.slug IN (%s)`, slugsSQL)
		config.SQL.Raw(fmt.Sprintf(populatedCityFromTerminalQuery, whereQuery)).Scan(&citiesFromTerminalSlugs)
	}

	if len(airportSlugs) > 0 {
		slugsSQL := utils.StringsForWhere(airportSlugs)
		whereQuery := fmt.Sprintf(`WHERE places.slug IN (%s)`, slugsSQL)
		config.SQL.Raw(fmt.Sprintf(populatedCityFromAirportQuery, whereQuery)).Scan(&citiesFromAirportSlugs)
	}

	cities := []City{}
	if len(citySlugs) > 0 {
		cities = append(cities, citiesFromCitySlugs...)
	}
	if len(terminalSlugs) > 0 {
		cities = append(cities, citiesFromTerminalSlugs...)
	}
	if len(airportSlugs) > 0 {
		cities = append(cities, citiesFromAirportSlugs...)
	}

	return cities
}

const populatedCityFromCityQuery = `
	SELECT cities.*, places.slug,
		states.name state_name, states.abbr state_abbr,
		countries.name country_name, countries.abbr country_abbr
	FROM cities
	INNER JOIN places ON findable_type='City' and findable_id = cities.id
	INNER JOIN states ON cities.state_id = states.id
	INNER JOIN countries ON states.country_id = countries.id
	%s
`

const populatedCityFromTerminalQuery = `
	SELECT cities.*, city_places.slug slug,
		states.name state_name, states.abbr state_abbr,
		countries.name country_name, countries.abbr country_abbr
	FROM cities
	INNER JOIN terminals ON terminals.city_id = cities.id
	INNER JOIN places ON places.findable_type='Terminal' and places.findable_id = terminals.id
	INNER JOIN states ON cities.state_id = states.id
	INNER JOIN countries ON states.country_id = countries.id
	inner join places city_places on city_places.findable_type='City' and city_places.findable_id = cities.id
	%s
`

const populatedCityFromAirportQuery = `
	SELECT cities.*, city_places.slug slug,
		states.name state_name, states.abbr state_abbr,
		countries.name country_name, countries.abbr country_abbr
	FROM cities
	INNER JOIN airports ON airports.city_id = cities.id
	INNER JOIN places ON places.findable_type='Airport' and places.findable_id = airports.id
	INNER JOIN states ON cities.state_id = states.id
	INNER JOIN countries ON states.country_id = countries.id
	inner join places city_places on city_places.findable_type='City' and city_places.findable_id = cities.id
	%s
`
