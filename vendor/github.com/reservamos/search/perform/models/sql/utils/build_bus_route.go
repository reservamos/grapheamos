package utils

import (
	"fmt"
	"time"

	"github.com/reservamos/search/internal/config"
)

// PopulatedBusRoute mapped result of the flight route query
type PopulatedBusRoute struct {
	ID               int
	OriginID         int
	DestinationID    int
	TransporterID    int
	Active           bool
	Attempts         int
	LastSuccessAt    time.Time
	DeactivatedAt    time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	OttID            int
	OttTerminalID    int
	OttTransporterID int
	OttCode          string
	OttCreatedAt     time.Time
	OttUpdatedAt     time.Time
	OtID             int
	OtName           string
	OtASCIIName      string
	OtLat            float64
	OtLong           float64
	OtCityID         int
	OcID             int
	OcName           string
	OcASCIIName      string
	OcTimeZoneID     string
	OcStateID        int
	OcCreatedAt      time.Time
	OcUpdatedAt      time.Time
	OtpID            int
	OtpSlug          string
	OtpFindableID    int
	OtpFindableType  string
	OtpHits          int
	OtpPriority      int
	OtpCreatedAt     time.Time
	OtpUpdatedAt     time.Time
	OtpPopularity    float64
	OcpID            int
	OcpSlug          string
	OcpFindableID    int
	OcpFindableType  string
	OcpHits          int
	OcpPriority      int
	OcpCreatedAt     time.Time
	OcpUpdatedAt     time.Time
	OcpPopularity    float64
	DttID            int
	DttTerminalID    int
	DttTransporterID int
	DttCode          string
	DttCreatedAt     time.Time
	DttUpdatedAt     time.Time
	DtID             int
	DtName           string
	DtASCIIName      string
	DtLat            float64
	DtLong           float64
	DtCityID         int
	DcID             int
	DcName           string
	DcASCIIName      string
	DcTimeZoneID     string
	DcStateID        int
	DcCreatedAt      time.Time
	DcUpdatedAt      time.Time
	DtpID            int
	DtpSlug          string
	DtpFindableID    int
	DtpFindableType  string
	DtpHits          int
	DtpPriority      int
	DtpCreatedAt     time.Time
	DtpUpdatedAt     time.Time
	DtpPopularity    float64
	DcpID            int
	DcpSlug          string
	DcpFindableID    int
	DcpFindableType  string
	DcpHits          int
	DcpPriority      int
	DcpCreatedAt     time.Time
	DcpUpdatedAt     time.Time
	DcpPopularity    float64
	TID              int
	TAbbr            string
	TName            string
	TCreatedAt       time.Time
	TUpdatedAt       time.Time
}

// BuildBusRouteQuery returns the big inner join query for a bus route
func BuildBusRouteQuery(id int) (*PopulatedBusRoute, error) {
	var result PopulatedBusRoute
	query := busRouteQuery
	query = query + routeCondition
	err := config.SQL.Raw(fmt.Sprintf(query, id)).Scan(&result).Error
	return &result, err
}

// BuildBusRouteQuery returns the big inner join query for a bus route
func BuildBusRouteQueries(cityOriginSlug, cityDestinationSlug string) ([]*PopulatedBusRoute, error) {
	var result []*PopulatedBusRoute
	query := busRouteQuery
	query = query + cityConnectionCondition
	err := config.SQL.Raw(fmt.Sprintf(query, cityOriginSlug, cityDestinationSlug)).Scan(&result).Error
	return result, err
}

const busRouteQuery = `
SELECT
  routes.id,
  routes.origin_id,
  routes.destination_id,
  routes.transporter_id,
  routes.active,
  routes.last_success_at,
  routes.attempts,
  routes.deactivated_at,
	routes.created_at,
	routes.updated_at,
  ott.id "ott_id",
  ott.terminal_id "ott_terminal_id",
  ott.transporter_id "ott_transporter_id",
  ott.code "ott_code",
  ott.created_at "ott_created_at",
  ott.updated_at "ott_updated_at",
  ot.id "ot_id",
  ot.name "ot_name",
  ot.ascii_name "ot_ascii_name",
  ot.lat "ot_lat",
  ot.long "ot_long",
  ot.city_id "ot_city_id",
  oc.id "oc_id",
  oc.name "oc_name",
  oc.ascii_name "oc_ascii_name",
  oc.created_at "oc_created_at",
  oc.updated_at "oc_updated_at",
  oc.time_zone_id "oc_time_zone_id",
  oc.state_id "oc_state_id",
	otp.id "otp_id",
	otp.slug "otp_slug",
	otp.findable_id "otp_findable_id",
	otp.findable_type "otp_findable_type",
	otp.hits "otp_hits",
	otp.priority "otp_priority",
	otp.popularity "otp_popularity",
	otp.created_at "otp_created_at",
	otp.updated_at "otp_updated_at",
	ocp.id "ocp_id",
	ocp.slug "ocp_slug",
	ocp.findable_id "ocp_findable_id",
	ocp.findable_type "ocp_findable_type",
	ocp.hits "ocp_hits",
	ocp.priority "ocp_priority",
	ocp.popularity "ocp_popularity",
	ocp.created_at "ocp_created_at",
	ocp.updated_at "ocp_updated_at",
  dtt.id "dtt_id",
  dtt.terminal_id "dtt_terminal_id",
  dtt.transporter_id "dtt_transporter_id",
  dtt.code "dtt_code",
  dtt.created_at "dtt_created_at",
  dtt.updated_at "dtt_updated_at",
  dt.id "dt_id",
  dt.name "dt_name",
  dt.ascii_name "dt_ascii_name",
  dt.lat "dt_lat",
  dt.long "dt_long",
  dt.city_id "dt_city_id",
  dc.id "dc_id",
  dc.name "dc_name",
  dc.ascii_name "dc_ascii_name",
  dc.created_at "dc_created_at",
  dc.updated_at "dc_updated_at",
  dc.time_zone_id "dc_time_zone_id",
  dc.state_id "dc_state_id",
	dtp.id "dtp_id",
	dtp.slug "dtp_slug",
	dtp.findable_id "dtp_findable_id",
	dtp.findable_type "dtp_findable_type",
	dtp.hits "dtp_hits",
	dtp.priority "dtp_priority",
	dtp.popularity "dtp_popularity",
	dtp.created_at "dtp_created_at",
	dtp.updated_at "dtp_updated_at",
	dcp.id "dcp_id",
	dcp.slug "dcp_slug",
	dcp.findable_id "dcp_findable_id",
	dcp.findable_type "dcp_findable_type",
	dcp.hits "dcp_hits",
	dcp.priority "dcp_priority",
	dcp.popularity "dcp_popularity",
	dcp.created_at "dcp_created_at",
	dcp.updated_at "dcp_updated_at",
  transporters.id "t_id",
  transporters.abbr "t_abbr",
  transporters.name "t_name",
  transporters.created_at "t_created_at",
  transporters.updated_at "t_updated_at"
FROM routes
  INNER JOIN transporter_terminals ott ON routes.origin_id = ott.id
  INNER JOIN terminals ot ON ott.terminal_id = ot.id
  INNER JOIN cities oc ON ot.city_id = oc.id
  INNER JOIN states ost ON oc.state_id = ost.id
  INNER JOIN countries oco ON ost.country_id = oco.id
	INNER JOIN places otp ON ot.id = otp.findable_id AND otp.findable_type = 'Terminal'
	INNER JOIN places ocp ON oc.id = ocp.findable_id AND ocp.findable_type = 'City'
  INNER JOIN transporter_terminals dtt ON routes.destination_id = dtt.id
  INNER JOIN terminals dt ON dtt.terminal_id = dt.id
  INNER JOIN cities dc ON dt.city_id = dc.id
  INNER JOIN states dst ON dc.state_id = dst.id
  INNER JOIN countries dco ON dst.country_id = dco.id
	INNER JOIN places dtp ON dt.id = dtp.findable_id AND dtp.findable_type = 'Terminal'
	INNER JOIN places dcp ON dc.id = dcp.findable_id AND dcp.findable_type = 'City'
  INNER JOIN transporters ON routes.transporter_id = transporters.id
	WHERE
`

const routeCondition = `
  routes.id = %d and routes.active = 't';
`

const cityConnectionCondition = `
  ocp.slug = '%s' and dcp.slug = '%s' and routes.active = 't';
`
