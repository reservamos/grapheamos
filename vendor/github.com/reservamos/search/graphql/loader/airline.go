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
	airlineLoaderKey          key = "airline"
	airlineIDField                = "id"
	airlineLoaderCarrierIDKey key = "airline_carrier_id"
	airlineCarrierIDField         = "carrier_id"
)

func init() {
	loaders[airlineLoaderCarrierIDKey] = airlineLoader{}.loadBatchCarrierID
	loaders[airlineLoaderKey] = airlineLoader{}.loadBatch

}

// FilmLoader contains the client required to load film resources.
type airlineLoader struct{}

func (l airlineLoader) loadBatchCarrierID(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var airlines []sql.Airline
	query := DataLoaderQueryStrings(keys, sql.Airline{}, airlineCarrierIDField)
	config.SQL.Raw(query).Scan(&airlines)

	var results = make([]*dataloader.Result, len(keys))
	for i, a := range airlines {
		results[i] = &dataloader.Result{Data: a}
	}

	return results
}

// LoadAirlineWithCarrierID loads a airline using dataloader using carrier id
func LoadAirlineWithCarrierID(ctx context.Context, key string) (*sql.Airline, error) {
	var a sql.Airline

	data, err := Load(ctx, airlineLoaderCarrierIDKey, key)

	if err != nil {
		return nil, err
	}

	a, ok := data.(sql.Airline)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", a, data)
	}

	return &a, nil
}

func (l airlineLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var airlines []sql.Airline
	query := DataLoaderQueryInts(keys, sql.Airline{}, airlineIDField)
	config.SQL.Raw(query).Scan(&airlines)

	var results = make([]*dataloader.Result, len(keys))
	for i, a := range airlines {
		results[i] = &dataloader.Result{Data: a}
	}

	return results
}

// LoadAirline loads a airline using dataloader
func LoadAirline(ctx context.Context, key int) (*sql.Airline, error) {
	var a sql.Airline

	data, err := Load(ctx, airlineLoaderKey, strconv.Itoa(key))

	if err != nil {
		return nil, err
	}

	a, ok := data.(sql.Airline)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", a, data)
	}

	return &a, nil
}
