package sql

import (
	"fmt"
	"strings"

	"github.com/reservamos/search/internal/config"
)

// Terminal is a physical bus terminal inside of a city
type Terminal struct {
	ID        int     `gorm:"primary_key" dl:"id"`
	Address   string  `dl:"address"`
	Name      string  `dl:"name"`
	ASCIIName string  `dl:"ascii_name"`
	Lat       float64 `dl:"lat"`
	Long      float64 `dl:"long"`
	CityID    int     `dl:"city_id"`
	Slug      string  `dl:"slug"`
	City      City
	Place     Place `gorm:"polymorphic:Findable;polymorphic_value:Terminal"`
	// Only for LineFindAll
	CitySlug string
	CityName string
}

func TerminalsByCityID(cityID int) []Terminal {
	var terminals []Terminal
	config.SQL.Raw(fmt.Sprintf(`
	SELECT id, lat, long, city_id
	FROM terminals
	WHERE city_id=%d 
	`, cityID)).Scan(&terminals)
	return terminals
}

// TerminalFindAllForRoutes returns an array of terminals given an array of route ids
func TerminalFindAllForRoutes(ids []int) []Terminal {
	idsString := idsToString(ids)

	terminals := []Terminal{}

	if len(ids) > 0 {
		config.SQL.Raw(fmt.Sprintf(findAllQuery, idsString)).Scan(&terminals)
	}

	return terminals
}

func idsToString(ids []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")
}

const findAllQuery = `
	SELECT DISTINCT terminals.id,
		terminals.name, terminals.ascii_name, terminals.lat, terminals.long, tp.slug,
		cp.slug city_slug, cities.name city_name
	FROM transporter_terminals
		INNER JOIN terminals ON transporter_terminals.terminal_id = terminals.id
		INNER JOIN cities ON terminals.city_id = cities.id
		INNER JOIN places tp ON tp.findable_type = 'Terminal' AND tp.findable_id = terminals.id
		INNER JOIN places cp ON cp.findable_type = 'City' AND cp.findable_id = cities.id
		INNER JOIN routes ON (
			routes.origin_id = transporter_terminals.id
			OR routes.destination_id = transporter_terminals.id
		)
	WHERE routes.id IN (%s)
`
