package sql

import (
	"fmt"
	"strconv"

	"github.com/reservamos/search/internal/config"
	"github.com/reservamos/search/perform/models/sql/utils"
)

// Airline self described model
type Airline struct {
	ID          int     `gorm:"primary_key" dl:"id"`
	Name        string  `dl:"name"`
	Abbr        string  `dl:"abbr"`
	HumanAbbr   string  `dl:"human_abbr"`
	CarrierID   string  `dl:"carrier_id"`
	Logo        string  `dl:"logo"`
	Commission  float64 `dl:"commission"`
	BillingPage string  `dl:"billing_page"`
}

// LogoURL returns the absolute url for the line logo
func (a *Airline) LogoURL() string {
	return airlineLogoPrefix() + fmt.Sprintf("/%d/%s", a.ID, a.Logo)
}

// AirlinesWithCarrierIDs returns a slice of airlines given a slice of carrier ids
func AirlinesWithCarrierIDs(ids []string) []Airline {
	airlines := []Airline{}
	if len(ids) == 0 {
		return airlines
	}
	idsForQuery := utils.StringsForWhere(ids)
	query := `SELECT * FROM airlines WHERE carrier_id IN (%s)`
	config.SQL.Raw(fmt.Sprintf(query, idsForQuery)).Scan(&airlines)
	return airlines
}

func airlineLogoPrefix() string {
	return config.App.ImageRepo + "/uploads/airline/logo"
}

type airlineIdentityAttributes struct {
	ID        int
	CarrierID string
}

//CarrierToAirlineIDMap returns a map[string]string that maps carrier_ids to airline_ids
func CarrierToAirlineIDMap() map[string]string {
	airlineMap := make(map[string]string)

	var mappingSlice []airlineIdentityAttributes

	queryString := fmt.Sprintf(`
	SELECT id , carrier_id 
	FROM airlines`)

	config.SQL.Raw(queryString).Scan(&mappingSlice)

	for _, e := range mappingSlice {
		if _, ok := airlineMap[e.CarrierID]; !ok {
			airlineMap[e.CarrierID] = strconv.Itoa(e.ID)
		}
	}

	return airlineMap
}
