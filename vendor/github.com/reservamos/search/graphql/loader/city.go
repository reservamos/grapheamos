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
	cityLoaderKey key = "city"
	cityIDField       = "id"
)

func init() {
	loaders[cityLoaderKey] = cityLoader{}.loadBatch
}

// FilmLoader contains the client required to load film resources.
type cityLoader struct{}

func (l cityLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var cities []sql.City

	config.SQL.Raw(DataLoaderOpenQuery(
		keys,
		"places_cities",
		dlModelFields(sql.City{}),
		cityIDField,
		"int",
	)).Scan(&cities)

	var results = make([]*dataloader.Result, len(keys))
	for i, c := range cities {
		results[i] = &dataloader.Result{Data: c}
	}

	return results
}

// LoadCity loads a city using dataloader
func LoadCity(ctx context.Context, key int) (*sql.City, error) {
	var c sql.City

	data, err := Load(ctx, cityLoaderKey, strconv.Itoa(key))

	if err != nil {
		return nil, err
	}

	c, ok := data.(sql.City)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", c, data)
	}

	return &c, nil
}
