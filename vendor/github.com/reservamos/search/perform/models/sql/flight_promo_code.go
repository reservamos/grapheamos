package sql

import (
	"regexp"
	"strconv"
	"time"

	"github.com/lib/pq"
	"github.com/reservamos/search/internal/config"
	"github.com/reservamos/search/perform/utils"
)

// FlightPromoCode discount code used to send to airlines when requesting WS
type FlightPromoCode struct {
	ID                 int `gorm:"primary_key"`
	Code               string
	StartsAt           string
	ExpiresAt          string
	TripDepartureStart string
	TripDepartureEnd   string
	Blackouts          pq.StringArray `gorm:"type:varchar(64)[]"`
	Days               pq.StringArray `gorm:"type:varchar(64)[]"`
}

// FlightPromoCodesFor get flight promo codes for a specific flight route
func FlightPromoCodesFor(r *FlightRoute) []FlightPromoCode {
	var codes []FlightPromoCode
	config.SQL.Raw(`
		SELECT * FROM flight_promo_codes_routes
		INNER JOIN flight_promo_codes
			ON flight_promo_codes.id = flight_promo_codes_routes.flight_promo_code_id
			AND flight_promo_codes.starts_at <= ?
			AND flight_promo_codes.expires_at >= ?
		INNER JOIN flight_routes
			ON flight_routes.id = flight_promo_codes_routes.flight_route_id
			AND flight_promo_codes_routes.flight_route_id = ?
		`, time.Now(), time.Now(), r.ID).
		Scan(&codes)

	return codes
}

// CodeString is built to satisfy interface
func (fpc FlightPromoCode) CodeString() string {
	return fpc.Code
}

// Percent is the discount percentage of the promo code
func (fpc FlightPromoCode) Percent() int {
	r, _ := regexp.Compile("\\w*(\\d{2})")
	percentString := r.FindStringSubmatch(fpc.Code)[1]
	percent, _ := strconv.Atoi(percentString)

	return percent
}

// TripDepartureStartTime returns parsed time of start of discount
func (fpc *FlightPromoCode) TripDepartureStartTime() time.Time {
	t, _ := time.Parse("2006-01-02T00:00:00Z", fpc.TripDepartureStart)
	return t
}

// TripDepartureEndTime returns parsed time of end of discount
func (fpc *FlightPromoCode) TripDepartureEndTime() time.Time {
	t, _ := time.Parse("2006-01-02T00:00:00Z", fpc.TripDepartureEnd)
	return t
}

// Validate returns whether or not the promo code is valid for a specific date
func (fpc FlightPromoCode) Validate(date time.Time) bool {
	return (len(fpc.Days) == 0 || utils.Include(fpc.Days, date.Format("Mon"))) &&
		!utils.Include(fpc.Blackouts, date.Format("Mon")) &&
		(date.Add(1 * time.Second).After(fpc.TripDepartureStartTime())) &&
		(date.Before(fpc.TripDepartureEndTime().Add(1 * time.Second)))
}
