package utils

import (
	"fmt"
	"time"

	"github.com/reservamos/search/internal/config"
)

// PopulatedFlightRoute mapped result of the flight route query
type PopulatedFlightRoute struct {
	ID              int
	OriginID        int
	DestinationID   int
	CreatedAt       time.Time
	AgentID         int
	OaaID           int
	OaaAirportID    int
	OaaAgentID      int
	OaiID           int
	OaiName         string
	OaiIataCode     string
	OaiLong         float64
	OaiLat          float64
	OaiCityID       int
	OciID           int
	OciName         string
	OciASCIIName    string
	OciTimeZoneID   string
	OciUpdatedAt    time.Time
	OciCreatedAT    time.Time
	OcpID           int
	OcpSlug         string
	OcpFindableID   int
	OcpFindableType string
	OcpHits         int
	OcpPriority     int
	OcpCreatedAt    time.Time
	OcpUpdatedAt    time.Time
	OcpPopularity   float64
	DaaID           int
	DaaAirportID    int
	DaaAgentID      int
	DaiID           int
	DaiName         string
	DaiIataCode     string
	DaiLong         float64
	DaiLat          float64
	DaiCityID       int
	DciID           int
	DciName         string
	DciASCIIName    string
	DciTimeZoneID   string
	DciUpdatedAt    time.Time
	DciCreatedAT    time.Time
	DcpID           int
	DcpSlug         string
	DcpFindableID   int
	DcpFindableType string
	DcpHits         int
	DcpPriority     int
	DcpCreatedAt    time.Time
	DcpUpdatedAt    time.Time
	DcpPopularity   float64
	AgentName       string
	AgentAbbr       string
}

// BuildFlightRouteQuery returns the big inner join query for a flight route
func BuildFlightRouteQuery(id int) *PopulatedFlightRoute {
	var result PopulatedFlightRoute
	query := flightRouteQuery
	query = query + flightRouteCondition
	config.SQL.Raw(fmt.Sprintf(query, id)).Scan(&result)
	return &result
}

// BuildFlightRouteQuery returns the big inner join query for a flight route
func BuildFlightRouteQueries(cityOriginSlug string, cityDestinationSlug string) []*PopulatedFlightRoute {
	var result []*PopulatedFlightRoute
	query := flightRouteQuery
	query = query + flightCityConnectionCondition
	config.SQL.Raw(fmt.Sprintf(query, cityOriginSlug, cityDestinationSlug)).Scan(&result)
	return result
}

const flightRouteQuery = `
  SELECT flight_routes.id "id",
         flight_routes.origin_id "origin_id",
         flight_routes.destination_id "destination_id",
         flight_routes.agent_id "agent_id",
         flight_routes.created_at "created_at",
         oaa.id "oaa_id",
         oaa.airport_id "oaa_airport_id",
         oaa.agent_id "oaa_agent_id",
         oai.id "oai_id",
         oai.name "oai_name",
         oai.iata_code "oai_iata_code",
         oai.lat "oai_lat",
         oai.long "oai_long",
         oai.city_id "oai_city_id",
         oci.id "oci_id",
         oci.name "oci_name",
         oci.ascii_name "oci_ascii_name",
         oci.created_at "oci_created_at",
         oci.updated_at "oci_updated_at",
         oci.time_zone_id "oci_time_zone_id",
         oci.state_id "oci_state_id",
                  ocp.id "ocp_id",
                  ocp.slug "ocp_slug",
                  ocp.findable_id "ocp_findable_id",
                  ocp.findable_type "ocp_findable_type",
                  ocp.hits "ocp_hits",
                  ocp.priority "ocp_priority",
                  ocp.popularity "ocp_popularity",
                  ocp.created_at "ocp_created_at",
                  ocp.updated_at "ocp_updated_at",
         daa.id "daa_id",
         daa.airport_id "daa_airport_id",
         daa.agent_id "daa_agent_id",
         dai.id "dai_id",
         dai.name "dai_name",
         dai.iata_code "dai_iata_code",
         dai.lat "dai_lat",
         dai.long "dai_long",
         dai.city_id "dai_city_id",
         dci.id "dci_id",
         dci.name "dci_name",
         dci.ascii_name "dci_ascii_name",
         dci.created_at "dci_created_at",
         dci.updated_at "dci_updated_at",
         dci.time_zone_id "dci_time_zone_id",
         dci.state_id "dci_state_id",
                  dcp.id "dcp_id",
                  dcp.slug "dcp_slug",
                  dcp.findable_id "dcp_findable_id",
                  dcp.findable_type "dcp_findable_type",
                  dcp.hits "dcp_hits",
                  dcp.priority "dcp_priority",
                  dcp.popularity "dcp_popularity",
                  dcp.created_at "dcp_created_at",
                  dcp.updated_at "dcp_updated_at",
         agents.id "agent_id",
         agents.name "agent_name",
         agents.abbr "agent_abbr"
  FROM flight_routes
    INNER JOIN agent_airports oaa ON flight_routes.origin_id = oaa.id
    INNER JOIN airports oai ON oaa.airport_id = oai.id
    INNER JOIN cities oci ON oai.city_id = oci.id
    INNER JOIN states ost ON oci.state_id = ost.id
    INNER JOIN countries oco ON ost.country_id = oco.id
    INNER JOIN agent_airports daa ON flight_routes.destination_id = daa.id
          INNER JOIN places ocp ON oci.id = ocp.findable_id AND ocp.findable_type = 'City'
    INNER JOIN airports dai ON daa.airport_id = dai.id
    INNER JOIN cities dci ON dai.city_id = dci.id
    INNER JOIN states dst ON dci.state_id = dst.id
    INNER JOIN countries dco ON dst.country_id = dco.id
          INNER JOIN places dcp ON dci.id = dcp.findable_id AND dcp.findable_type = 'City'    
    INNER JOIN agents ON flight_routes.agent_id = agents.id
  WHERE
`

const flightRouteCondition = `
  flight_routes.id = %d;
`

const flightCityConnectionCondition = `
  ocp.slug = '%s' and dcp.slug = '%s' and flight_routes.active and flight_routes.should_crawl;
`
