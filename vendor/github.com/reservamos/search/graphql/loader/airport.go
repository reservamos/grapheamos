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
	airportLoaderKey key = "airport_iata"
	airportCodeField     = "iata_code"
	airportIDField       = "id"
	airportIDKey     key = "airport"
)

func init() {
	loaders[airportLoaderKey] = airportLoader{}.loadBatchWithIATA
	loaders[airportIDKey] = airportLoader{}.loadBatch
}

// FilmLoader contains the client required to load film resources.
type airportLoader struct{}

func (l airportLoader) loadBatchWithIATA(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var airports []sql.Airport
	query := DataLoaderQueryStrings(keys, sql.Airport{}, airportCodeField)
	config.SQL.Raw(query).Scan(&airports)

	var results = make([]*dataloader.Result, len(keys))
	for i, a := range airports {
		results[i] = &dataloader.Result{Data: a}
	}

	return results
}

// LoadAirportWithIATA loads a airport using dataloader
func LoadAirportWithIATA(ctx context.Context, key string) (*sql.Airport, error) {
	var a sql.Airport

	data, err := Load(ctx, airportLoaderKey, key)

	if err != nil {
		return nil, err
	}

	a, ok := data.(sql.Airport)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", a, data)
	}

	return &a, nil
}

func (l airportLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var airports []sql.Airport
	query := DataLoaderQueryInts(keys, sql.Airport{}, airportIDField)
	config.SQL.Raw(query).Scan(&airports)

	var results = make([]*dataloader.Result, len(keys))
	for i, a := range airports {
		results[i] = &dataloader.Result{Data: a}
	}

	return results
}

// LoadAirport loads a airport using dataloader
func LoadAirport(ctx context.Context, key int) (*sql.Airport, error) {
	var a sql.Airport

	data, err := Load(ctx, airportIDKey, strconv.Itoa(key))

	if err != nil {
		return nil, err
	}

	a, ok := data.(sql.Airport)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", a, data)
	}

	return &a, nil
}
