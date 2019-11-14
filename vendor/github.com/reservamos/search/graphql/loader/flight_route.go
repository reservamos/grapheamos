package loader

import (
	"fmt"
	"strconv"

	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/reservamos/search/internal/config"
	"github.com/reservamos/search/perform/models/sql"
)

const (
	flightRouteLoaderKey key = "flightRoute"
	flightRouteIDField       = "id"
)

func init() {
	loaders[flightRouteLoaderKey] = flightRouteLoader{}.loadBatch
}

// FilmLoader contains the client required to load film resources.
type flightRouteLoader struct{}

func (l flightRouteLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var flightRoutes []sql.FlightRoute
	query := DataLoaderQueryInts(keys, sql.FlightRoute{}, flightRouteIDField)
	config.SQL.Raw(query).Scan(&flightRoutes)

	var results = make([]*dataloader.Result, len(keys))
	for i, fr := range flightRoutes {
		results[i] = &dataloader.Result{Data: fr}
	}

	return results
}

// LoadFlightRoute loads a flightRoute using dataloader
func LoadFlightRoute(ctx context.Context, key int) (*sql.FlightRoute, error) {
	var fr sql.FlightRoute

	data, err := Load(ctx, flightRouteLoaderKey, strconv.Itoa(key))

	if err != nil {
		return nil, err
	}

	fr, ok := data.(sql.FlightRoute)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", fr, data)
	}

	return &fr, nil
}
