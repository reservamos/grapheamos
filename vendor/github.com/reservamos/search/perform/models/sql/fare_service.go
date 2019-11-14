package sql

import (
	"fmt"
	"strings"

	"github.com/lib/pq"
	"github.com/reservamos/search/internal/config"
	"github.com/reservamos/search/perform/utils"
)

// FareService self described model
type FareService struct {
	ID        int            `gorm:"primary_key" dl:"id"`
	AirlineID int            `dl:"airline_id"`
	Service   string         `dl:"service"`
	Services  pq.StringArray `gorm:"type:varchar(64)[]" dl:"services"`
	// Only for mass populate
	AirlineCarrierID string
}

// FareServicesWithServiceCodes returns a slice of FareService given a slice of service codes
func FareServicesWithServiceCodes(codes []string) []FareService {
	fareServices := []FareService{}

	if len(codes) == 0 {
		return fareServices
	}

	codesForQuery := airportServicePair(codes)
	query := `
		SELECT fare_services.*, airlines.carrier_id airline_carrier_id
		FROM fare_services
      INNER JOIN airlines ON fare_services.airline_id = airlines.id
		WHERE (carrier_id, service) IN (VALUES %s)
	`
	config.SQL.Raw(fmt.Sprintf(query, codesForQuery)).Scan(&fareServices)
	return fareServices
}

func airportServicePair(codes []string) string {
	return strings.Join(utils.Map(codes, func(code string) string {
		serviceCode := strings.Split(code, "-")
		return fmt.Sprintf("('%s', '%s')", serviceCode[0], serviceCode[1])
	}), ",")
}
