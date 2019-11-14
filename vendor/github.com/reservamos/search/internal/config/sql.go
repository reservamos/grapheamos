package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // needed to gorm.Open with "postgres" argument
)

// SQL gorm.DB postgres sql connection
var SQL *gorm.DB

type sqlConfig struct {
	URL  string `required:"true"`
	Pool int    `required:"true"`
	Log  bool   `default:"true"`
}

// RefreshMaterializedViews refreshes views on a schedule to prevent stale data
func RefreshMaterializedViews() {
	SQL.Exec(`
	REFRESH MATERIALIZED VIEW places_cities;
	REFRESH MATERIALIZED VIEW places_terminals;
	REFRESH MATERIALIZED VIEW active_bus_origin_cities;
	REFRESH MATERIALIZED VIEW active_bus_destination_cities;
	REFRESH MATERIALIZED VIEW active_flight_origin_cities;
	REFRESH MATERIALIZED VIEW active_flight_destination_cities;
	REFRESH MATERIALIZED VIEW active_origin_terminals;
	REFRESH MATERIALIZED VIEW active_destination_terminals;
	REFRESH MATERIALIZED VIEW active_origin_airports;
	REFRESH MATERIALIZED VIEW active_destination_airports; 
	REFRESH MATERIALIZED VIEW active_line_origin_cities;
	REFRESH MATERIALIZED VIEW active_line_destination_cities;
	REFRESH MATERIALIZED VIEW popular_searches_daily;
	REFRESH MATERIALIZED VIEW popular_searches_weekly;
	REFRESH MATERIALIZED VIEW bus_city_routes;
	REFRESH MATERIALIZED VIEW bus_transporter_city_routes;
	REFRESH MATERIALIZED VIEW bus_line_city_routes;
	REFRESH MATERIALIZED VIEW terminal_routes;
	REFRESH MATERIALIZED VIEW transporter_terminal_routes;
	REFRESH MATERIALIZED VIEW line_terminal_routes;
	REFRESH MATERIALIZED VIEW places_routes;
	REFRESH MATERIALIZED VIEW transporter_places_routes;
	REFRESH MATERIALIZED VIEW line_places_routes;
	REFRESH MATERIALIZED VIEW active_origin_airports;
	REFRESH MATERIALIZED VIEW active_destination_airports;
	REFRESH MATERIALIZED VIEW route_plans_cities;
	`)
}

func newSQLConfig() *sqlConfig {
	var c sqlConfig
	err := ReadConfig("SQL", &c)
	if err != nil {
		panic(err)
	}
	return &c
}

func initDB() {
	c := newSQLConfig()

	db, err := gorm.Open("postgres", c.URL)
	if err != nil {
		fmt.Println(err)
	}

	db.DB().SetMaxIdleConns(c.Pool)
	db.DB().SetMaxOpenConns(c.Pool)

	db.LogMode(c.Log)

	SQL = db
}
