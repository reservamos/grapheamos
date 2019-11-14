package sql

import (
	"fmt"
	"time"

	"github.com/reservamos/search/internal/config"
	"github.com/reservamos/search/internal/wsclient"
	"github.com/reservamos/search/perform/models/responses"
	"github.com/reservamos/search/perform/models/sql/utils"
)

// Route holds information from point A(TransporterTerminal) to point B (TransporterTerminal)
type Route struct {
	ID            int                 `gorm:"primary_key" dl:"id"`
	OriginID      int                 `dl:"origin_id"`
	DestinationID int                 `dl:"destination_id"`
	TransporterID int                 `dl:"transporter_id"`
	Active        bool                `dl:"active"`
	LastSuccessAt time.Time           `dl:"last_success_at"`
	Attempts      int                 `dl:"attempts"`
	ActivatedAt   time.Time           `dl:"activated_at"`
	DeactivatedAt time.Time           `dl:"deactivated_at"`
	CreatedAt     time.Time           `dl:"created_at"`
	UpdatedAt     time.Time           `dl:"updated_at"`
	Origin        TransporterTerminal `gorm:"ForeignKey:OriginID"`
	Destination   TransporterTerminal `gorm:"ForeignKey:DestinationID"`
	Transporter   Transporter
}

// RouteFindPopulated Finds route with prepopulated transporter, origin and destination (transporter terminals)
// TODO: Efficient but verbose, is there a better way to do this ?
func RouteFindPopulated(ID int) (*Route, error) {
	r, err := utils.BuildBusRouteQuery(ID)
	if err != nil {
		return nil, err
	}
	return routeFillPopulated(r), nil
}

// RoutesFindPopulated Finds route with prepopulated transporter, origin and destination (transporter terminals)
func RoutesFindPopulated(cityOriginSlug string, cityDestinationSlug string) ([]*Route, error) {
	routes, err := utils.BuildBusRouteQueries(cityOriginSlug, cityDestinationSlug)
	if err != nil {
		return nil, err
	}
	var res []*Route
	for _, r := range routes {
		res = append(res, routeFillPopulated(r))
	}
	return res, nil
}

func routeFillPopulated(r *utils.PopulatedBusRoute) *Route {
	route := Route{
		ID: r.ID,
		Origin: TransporterTerminal{
			ID:         r.OttID,
			TerminalID: r.OttTerminalID,
			Terminal: Terminal{
				ID:        r.OtID,
				Name:      r.OtName,
				ASCIIName: r.OtASCIIName,
				Lat:       r.OtLat,
				Long:      r.OtLong,
				CityID:    r.OtCityID,
				Place: Place{
					ID:           r.OtpID,
					Slug:         r.OtpSlug,
					FindableID:   r.OtpFindableID,
					FindableType: r.OtpFindableType,
					Hits:         r.OtpHits,
					Priority:     r.OtpPriority,
					Popularity:   r.OtpPopularity,
					CreatedAt:    r.OtpCreatedAt,
					UpdatedAt:    r.OtpUpdatedAt,
				},
				City: City{
					ID:         r.OcID,
					Name:       r.OcName,
					ASCIIName:  r.OcASCIIName,
					TimeZoneID: r.OcTimeZoneID,
					StateID:    r.OcStateID,
					CreatedAt:  r.OcCreatedAt,
					UpdatedAt:  r.OcUpdatedAt,
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
			TransporterID: r.OttTransporterID,
			Code:          r.OttCode,
			CreatedAt:     r.OttCreatedAt,
			UpdatedAt:     r.OttUpdatedAt,
		},
		OriginID: r.OriginID,
		Destination: TransporterTerminal{
			ID:         r.DttID,
			TerminalID: r.DttTerminalID,
			Terminal: Terminal{
				ID:        r.DtID,
				Name:      r.DtName,
				ASCIIName: r.DtASCIIName,
				Lat:       r.DtLat,
				Long:      r.DtLong,
				CityID:    r.DtCityID,
				Place: Place{
					ID:           r.DtpID,
					Slug:         r.DtpSlug,
					FindableID:   r.DtpFindableID,
					FindableType: r.DtpFindableType,
					Hits:         r.DtpHits,
					Priority:     r.DtpPriority,
					Popularity:   r.DtpPopularity,
					CreatedAt:    r.DtpCreatedAt,
					UpdatedAt:    r.DtpUpdatedAt,
				},
				City: City{
					ID:         r.DcID,
					Name:       r.DcName,
					ASCIIName:  r.DcASCIIName,
					TimeZoneID: r.DcTimeZoneID,
					StateID:    r.DcStateID,
					CreatedAt:  r.DcCreatedAt,
					UpdatedAt:  r.DcUpdatedAt,
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
			TransporterID: r.DttTransporterID,
			Code:          r.DttCode,
			CreatedAt:     r.DttCreatedAt,
			UpdatedAt:     r.DttUpdatedAt,
		},
		DestinationID: r.DestinationID,
		Transporter: Transporter{
			ID:        r.TID,
			Abbr:      r.TAbbr,
			Name:      r.TName,
			CreatedAt: r.TCreatedAt,
			UpdatedAt: r.TUpdatedAt,
		},
		TransporterID: r.TransporterID,
		Active:        r.Active,
		LastSuccessAt: r.LastSuccessAt,
		Attempts:      r.Attempts,
		DeactivatedAt: r.DeactivatedAt,
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
	}
	config.SQL.Model(&route.Transporter).Association("Lines").Find(&route.Transporter.Lines)
	if len(route.Transporter.Lines) > 0 {
		return &route
	}
	route.Transporter.Lines = AggregatedLines(route.Transporter.ID)
	return &route
}

// RouteIDsByPlaceSlugs finds routes given place slugs
// TODO: Cache results to prevent long queries
func RouteIDsByPlaceSlugs(originSlug string, destinationSlug string) []int {
	var routeIds []int
	var ids []struct {
		ID int
	}

	queryPartial := placeSlugsQuery[utils.PlaceType(originSlug)][utils.PlaceType(destinationSlug)]
	query := fmt.Sprintf(queryPartial+`
		WHERE routes.active AND transporters.purchase_enabled AND op.slug = '%s' AND dp.slug = '%s'
	`, originSlug, destinationSlug)
	config.SQL.Raw(query).Scan(&ids)

	for _, result := range ids {
		routeIds = append(routeIds, result.ID)
	}
	return routeIds
}

// Crawl method performs a crawl in the integrations web service and returns an array of responses.Bus
func (r *Route) Crawl(date time.Time) (*[]responses.Bus, error) {
	var oCoordinates wsclient.RouteCoordinate
	var dCoordinates wsclient.RouteCoordinate
	r.routeGetCoordinates(&oCoordinates, &dCoordinates)

	args := wsclient.BusesRequestArgs{
		Transporter:           r.Transporter.Abbr,
		Departure:             date,
		Origin:                r.Origin.Code,
		Destination:           r.Destination.Code,
		OriginCoordinate:      oCoordinates,
		DestinationCoordinate: dCoordinates,
	}

	response, err := wsclient.GetBuses(args)
	if err != nil {
		return nil, err
	}

	return &response.Buses, err
}

// ProductID is used for marketing purposes in reporting
func (r *Route) ProductID() string {
	return r.Origin.Terminal.City.Place.Slug + "-" + r.Destination.Terminal.City.Place.Slug
}

// routeGetCoordinates populates origin wsclient.RouteCoordinate and destination wsclient.RouteCoordinate
// with the appropriate terminal coordinates. This method helps crawl parameters
func (r *Route) routeGetCoordinates(originCoord *wsclient.RouteCoordinate, destinationCoord *wsclient.RouteCoordinate) {
	var routeCoords struct {
		OriginLat      float64
		OriginLon      float64
		DestinationLat float64
		DestinationLon float64
	}

	query := fmt.Sprintf(`
		SELECT ot.lat origin_lat, ot.long origin_lon, dt.lat destination_lat, dt.long destination_lon
		FROM routes
			INNER JOIN transporter_terminals ott ON routes.origin_id = ott.id
			INNER JOIN terminals ot ON ott.terminal_id = ot.id
			INNER JOIN transporter_terminals dtt ON routes.destination_id = dtt.id
			INNER JOIN terminals dt ON dtt.terminal_id = dt.id
			WHERE routes.origin_id = %d AND routes.destination_id = %d
	`, r.OriginID, r.DestinationID)
	config.SQL.Raw(query).Scan(&routeCoords)
	if routeCoords.OriginLat != 0 && routeCoords.DestinationLat != 0 {
		originCoord.Lat = routeCoords.OriginLat
		originCoord.Lon = routeCoords.OriginLon
		destinationCoord.Lat = routeCoords.DestinationLat
		destinationCoord.Lon = routeCoords.DestinationLon
	}
}

// RoutesByPlaceSlugs helper queries to prevent string building overhead
var placeSlugsQuery = map[string]map[string]string{
	"city": map[string]string{
		"city": `
			SELECT routes.id
			FROM routes
			INNER JOIN transporter_terminals ott ON routes.origin_id = ott.id
			INNER JOIN terminals ot ON ott.terminal_id = ot.id
			INNER JOIN cities oc ON ot.city_id = oc.id
			INNER JOIN places op ON op.findable_type = 'City' AND op.findable_id = oc.id
			INNER JOIN transporter_terminals dtt ON routes.destination_id = dtt.id
			INNER JOIN terminals dt ON dtt.terminal_id = dt.id
			INNER JOIN cities dc ON dt.city_id = dc.id
			INNER JOIN places dp ON dp.findable_type = 'City' AND dp.findable_id = dc.id
			INNER JOIN transporters ON routes.transporter_id = transporters.id
		`,
		"terminal": `
			SELECT routes.id
			FROM routes
			INNER JOIN transporter_terminals ott ON routes.origin_id = ott.id
			INNER JOIN terminals ot ON ott.terminal_id = ot.id
			INNER JOIN cities oc ON ot.city_id = oc.id
			INNER JOIN places op ON op.findable_type = 'City' AND op.findable_id = oc.id
			INNER JOIN transporter_terminals dtt ON routes.destination_id = dtt.id
			INNER JOIN terminals dt ON dtt.terminal_id = dt.id
			INNER JOIN places dp ON dp.findable_type = 'Terminal' AND dp.findable_id = dt.id
			INNER JOIN transporters ON routes.transporter_id = transporters.id
		`,
	},
	"terminal": map[string]string{
		"city": `
			SELECT routes.id
			FROM routes
			INNER JOIN transporter_terminals ott ON routes.origin_id = ott.id
			INNER JOIN terminals ot ON ott.terminal_id = ot.id
			INNER JOIN places op ON op.findable_type = 'Terminal' AND op.findable_id = ot.id
			INNER JOIN transporter_terminals dtt ON routes.destination_id = dtt.id
			INNER JOIN terminals dt ON dtt.terminal_id = dt.id
			INNER JOIN cities dc ON dt.city_id = dc.id
			INNER JOIN places dp ON dp.findable_type = 'City' AND dp.findable_id = dc.id
			INNER JOIN transporters ON routes.transporter_id = transporters.id
		`,
		"terminal": `
			SELECT routes.id
			FROM routes
			INNER JOIN transporter_terminals ott ON routes.origin_id = ott.id
			INNER JOIN terminals ot ON ott.terminal_id = ot.id
			INNER JOIN places op ON op.findable_type = 'Terminal' AND op.findable_id = ot.id
			INNER JOIN transporter_terminals dtt ON routes.destination_id = dtt.id
			INNER JOIN terminals dt ON dtt.terminal_id = dt.id
			INNER JOIN places dp ON dp.findable_type = 'Terminal' AND dp.findable_id = dt.id
			INNER JOIN transporters ON routes.transporter_id = transporters.id
		`,
	},
}
