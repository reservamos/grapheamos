package sql

import (
	"fmt"
	"net/url"
	"time"

	"github.com/lib/pq"
	"github.com/reservamos/search/internal/config"
	"github.com/reservamos/search/perform/models/sql/utils"
)

// Line is a Busline, tipically part of a Transporter
type Line struct {
	ID            int            `gorm:"primary_key" dl:"id"`
	Name          string         `dl:"name"`
	Abbr          string         `dl:"abbr"`
	HumanAbbr     string         `dl:"human_abbr"`
	Services      pq.StringArray `gorm:"type:varchar(64)[]" dl:"services"`
	ServiceType   string         `dl:"service_type"`
	Logo          string         `dl:"logo"`
	TransporterID int            `dl:"transporter_id"`
	CreatedAt     time.Time      `dl:"created_at"`
	UpdatedAt     time.Time      `dl:"updated_at"`
	Transporter   Transporter
	// Optional Fields only when populated
	TransporterName       string
	TransporterAbbr       string
	AllowsSeatSelection   bool
	VolatilePricing       bool
	Ally                  bool
	TicketCounterExchange bool
	Commission            float64
	RatingsTotal          int64
	RatingsAverage        float64
}

// LineFindAllPopulatedWithSlugs finds busline with transporter and ratings
func LineFindAllPopulatedWithSlugs(lineSlugs []string) []Line {
	var lines []Line
	if len(lineSlugs) == 0 {
		return lines
	}
	slugsSQL := utils.StringsForWhere(lineSlugs)
	whereQuery := fmt.Sprintf(`WHERE lines.abbr IN (%s)`, slugsSQL)
	config.SQL.Raw(fmt.Sprintf(populatedLineQuery, whereQuery)).Scan(&lines)
	return lines
}

// AggregatedLines returns a list of lines based on an aggregator ID
func AggregatedLines(aggID int) []Line {
	var lines []Line
	config.SQL.Raw(fmt.Sprintf(`
		SELECT lines.* 
		FROM lines
			INNER JOIN transporters ON lines.transporter_id = transporters.id
			INNER JOIN transporters aggregators ON transporters.aggregator_id = aggregators.id
		WHERE aggregators.id = %d
	`, aggID)).Scan(&lines)
	return lines
}

// LineFindAllPopulated finds all buslines with transporter and ratings
func LineFindAllPopulated() []Line {
	var lines []Line
	config.SQL.Raw(fmt.Sprintf(populatedLineQuery, "")).Scan(&lines)
	return lines
}

// LogoURL returns the absolute url for the line logo
func (l *Line) LogoURL() string {
	if l.Logo == "" {
		return ""
	}
	return config.App.ImageRepo + "/uploads/line/logo" + fmt.Sprintf("/%d/%s", l.ID, url.QueryEscape(l.Logo))
}

const populatedLineQuery = `
	SELECT
		lines.id, lines.name, lines.abbr, lines.human_abbr, lines.services, lines.service_type, lines.logo,
		transporters.name transporter_name, transporters.abbr transporter_abbr, transporters.allows_seat_selection,
		transporters.volatile_pricing, transporters.ally, transporters.ticket_counter_exchange,
		transporters.commission,
		COUNT(ratings.*) ratings_total, AVG(ratings.general) ratings_average
	FROM lines
		INNER JOIN transporters ON lines.transporter_id = transporters.id
		LEFT OUTER JOIN ratings ON ratings.line_id = lines.id
	%s
	GROUP BY
		lines.id, lines.name, lines.abbr, lines.human_abbr, lines.services, lines.service_type, lines.logo,
		transporter_name, transporter_abbr, transporters.allows_seat_selection, transporters.volatile_pricing,
		transporters.ally, transporters.ticket_counter_exchange, transporters.commission
`
