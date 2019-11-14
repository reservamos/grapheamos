package sql

import (
	"fmt"
	"time"

	"github.com/reservamos/search/internal/config"
	"github.com/reservamos/search/internal/wsclient"
	"github.com/reservamos/search/perform/models/responses"
	"github.com/reservamos/search/perform/models/sql/utils"
)

// FlightRoute Flight routes obtained from postgress
type FlightRoute struct {
	ID            int          `gorm:"primary_key" dl:"id"`
	OriginID      int          `dl:"origin_id"`
	DestinationID int          `dl:"destination_id"`
	AgentID       int          `dl:"agent_id"`
	Active        bool         `dl:"active"`
	LastSuccessAt time.Time    `dl:"last_success_at"`
	Attempts      int          `dl:"attempts"`
	ActivatedAt   time.Time    `dl:"activated_at"`
	DeactivatedAt time.Time    `dl:"deactivated_at"`
	CreatedAt     time.Time    `dl:"created_at"`
	UpdatedAt     time.Time    `dl:"updated_at"`
	Destination   AgentAirport `gorm:"ForeignKey:DestinationId"`
	Origin        AgentAirport `gorm:"ForeignKey:OriginId"`
	PromoCodes    []FlightPromoCode
	Agent         Agent
}

// FlightRouteFindPopulated Finds route with prepopulated transporter, origin and destination (transporter terminals)
// TODO: Efficient but verbose, is there a better way to do this ?
func FlightRouteFindPopulated(id int) *FlightRoute {
	r := utils.BuildFlightRouteQuery(id)
	return flightRouteFillPopulated(r)
}

// FlightRouteFindPopulated Finds route with prepopulated transporter, origin and destination (transporter terminals)
// TODO: Efficient but verbose, is there a better way to do this ?
func FlightRoutesFindPopulated(cityOriginSlug string, cityDestinationSlug string) []*FlightRoute {
	routes := utils.BuildFlightRouteQueries(cityOriginSlug, cityDestinationSlug)
	var res []*FlightRoute
	for _, r := range routes {
		res = append(res, flightRouteFillPopulated(r))
	}

	return res
}

func flightRouteFillPopulated(r *utils.PopulatedFlightRoute) *FlightRoute {
	return &FlightRoute{
		ID:            r.ID,
		OriginID:      r.OriginID,
		DestinationID: r.DestinationID,
		AgentID:       r.AgentID,
		CreatedAt:     r.CreatedAt,
		Origin: AgentAirport{
			ID:        r.OaaID,
			AirportID: r.OaaAirportID,
			AgentID:   r.OaaAgentID,
			Airport: Airport{
				ID:       r.OaiID,
				Name:     r.OaiName,
				IataCode: r.OaiIataCode,
				Long:     r.OaiLong,
				Lat:      r.OaiLat,
				CityID:   r.OaiCityID,
				City: City{
					ID:         r.OciID,
					Name:       r.OciName,
					ASCIIName:  r.OciASCIIName,
					TimeZoneID: r.OciTimeZoneID,
					UpdatedAt:  r.OciUpdatedAt,
					CreatedAt:  r.OciCreatedAT,
					Place: Place{
						ID:           r.OcpID,
						Slug:         r.OcpSlug,
						FindableID:   r.OcpFindableID,
						FindableType: r.OcpFindableType,
						Hits:         r.OcpHits,
						Priority:     r.OcpPriority,
						Popularity:   r.OcpPopularity,
						CreatedAt:    r.OcpCreatedAt,
						UpdatedAt:    r.OcpUpdatedAt,
					},
				},
			},
		},
		Destination: AgentAirport{
			ID:        r.DaaID,
			AirportID: r.DaaAirportID,
			AgentID:   r.DaaAgentID,
			Airport: Airport{
				ID:       r.DaiID,
				Name:     r.DaiName,
				IataCode: r.DaiIataCode,
				Long:     r.DaiLong,
				Lat:      r.DaiLat,
				CityID:   r.DaiCityID,
				City: City{
					ID:         r.DciID,
					Name:       r.DciName,
					ASCIIName:  r.DciASCIIName,
					TimeZoneID: r.DciTimeZoneID,
					UpdatedAt:  r.DciUpdatedAt,
					CreatedAt:  r.DciCreatedAT,
					Place: Place{
						ID:           r.DcpID,
						Slug:         r.DcpSlug,
						FindableID:   r.DcpFindableID,
						FindableType: r.DcpFindableType,
						Hits:         r.DcpHits,
						Priority:     r.DcpPriority,
						Popularity:   r.DcpPopularity,
						CreatedAt:    r.DcpCreatedAt,
						UpdatedAt:    r.DcpUpdatedAt,
					},
				},
			},
		},
		Agent: Agent{
			ID:   r.AgentID,
			Name: r.AgentName,
			Abbr: r.AgentAbbr,
		},
	}
}

// FlightRouteIDsByPlaceSlugs finds flight routes given place slugs
func FlightRouteIDsByPlaceSlugs(originSlug string, destinationSlug string) []int {
	var flightRouteIds []int
	var ids []struct {
		ID int
	}

	queryPartial := flightPlaceSlugsQuery[utils.PlaceType(originSlug)][utils.PlaceType(destinationSlug)]
	query := fmt.Sprintf(queryPartial+`
		WHERE flight_routes.active AND op.slug = '%s' AND dp.slug = '%s'
	`, originSlug, destinationSlug)
	config.SQL.Raw(query).Scan(&ids)

	for _, result := range ids {
		flightRouteIds = append(flightRouteIds, result.ID)
	}
	return flightRouteIds
}

// FlightRouteByOriginAndDestination returns a flight route given origin and destination id
func FlightRouteByOriginAndDestination(oID int, dID int) (FlightRoute, error) {
	fr := FlightRoute{OriginID: oID, DestinationID: dID}

	db := config.SQL.First(&fr)
	if db.Error != nil {
		return fr, db.Error
	}

	return fr, nil
}

//Crawl method performs a crawl in the integrations web service and returns an array of responses.Bus
func (r *FlightRoute) Crawl(date time.Time, passengers []string) (*[]responses.Flight, error) {

	args := wsclient.FlightsRequestArgs{
		Transporter: r.Agent.Abbr,
		Db:          date,
		Origin:      r.Origin.Airport.IataCode,
		Destination: r.Destination.Airport.IataCode,
		Passengers:  passengers,
	}

	if code := r.BestPromoCodeForDate(date); code != "" {
		args.Discount = wsclient.FlightDiscount{
			Type:     "promotion",
			Category: "adult",
			Code:     code,
		}
	}

	response, err := wsclient.GetFlights(args)
	if err != nil {
		return nil, err
	}
	return &response.Flights, err
}

// BestPromoCodeForDate gets string for best promo code on a given date
func (r *FlightRoute) BestPromoCodeForDate(date time.Time) string {
	flightPromoCodes := FlightPromoCodesFor(r)
	promoCodes := make([]utils.PromoCode, len(flightPromoCodes))
	for i := range flightPromoCodes {
		promoCodes[i] = flightPromoCodes[i]
	}
	return utils.BestPromoCode(promoCodes, date)
}

var flightPlaceSlugsQuery = map[string]map[string]string{
	"city": map[string]string{
		"city": `
	    SELECT flight_routes.id
	      FROM flight_routes
	    INNER JOIN agent_airports oaa ON flight_routes.origin_id = oaa.id
	    INNER JOIN airports oa ON oaa.airport_id = oa.id
	    INNER JOIN cities oc ON oa.city_id = oc.id
	    INNER JOIN places op ON op.findable_type = 'City' AND op.findable_id = oc.id
	    INNER JOIN agent_airports daa ON flight_routes.destination_id = daa.id
	    INNER JOIN airports da ON daa.airport_id = da.id
	    INNER JOIN cities dc ON da.city_id = dc.id
	    INNER JOIN places dp ON dp.findable_type = 'City' AND dp.findable_id = dc.id
		`,
		"airport": `
	    SELECT flight_routes.id
	      FROM flight_routes
	    INNER JOIN agent_airports oaa ON flight_routes.origin_id = oaa.id
	    INNER JOIN airports oa ON oaa.airport_id = oa.id
	    INNER JOIN cities oc ON oa.city_id = oc.id
	    INNER JOIN places op ON op.findable_type = 'City' AND op.findable_id = oc.id
	    INNER JOIN agent_airports daa ON flight_routes.destination_id = daa.id
	    INNER JOIN airports da ON daa.airport_id = da.id
	    INNER JOIN places dp ON dp.findable_type = 'Airport' AND dp.findable_id = da.id
		`,
	},
	"airport": map[string]string{
		"city": `
	    SELECT flight_routes.id
	      FROM flight_routes
	    INNER JOIN agent_airports oaa ON flight_routes.origin_id = oaa.id
	    INNER JOIN airports oa ON oaa.airport_id = oa.id
	    INNER JOIN places op ON op.findable_type = 'Airport' AND op.findable_id = oa.id
	    INNER JOIN agent_airports daa ON flight_routes.destination_id = daa.id
	    INNER JOIN airports da ON daa.airport_id = da.id
	    INNER JOIN cities dc ON da.city_id = dc.id
	    INNER JOIN places dp ON dp.findable_type = 'City' AND dp.findable_id = dc.id
		`,
		"airport": `
	    SELECT flight_routes.id
	      FROM flight_routes
	    INNER JOIN agent_airports oaa ON flight_routes.origin_id = oaa.id
	    INNER JOIN airports oa ON oaa.airport_id = oa.id
	    INNER JOIN places op ON op.findable_type = 'Airport' AND op.findable_id = oa.id
	    INNER JOIN agent_airports daa ON flight_routes.destination_id = daa.id
	    INNER JOIN airports da ON daa.airport_id = da.id
	    INNER JOIN places dp ON dp.findable_type = 'Airport' AND dp.findable_id = da.id
		`,
	},
}
