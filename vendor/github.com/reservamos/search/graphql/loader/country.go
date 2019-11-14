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
	countryLoaderKey key = "country"
	countryIDField       = "id"
)

func init() {
	loaders[countryLoaderKey] = countryLoader{}.loadBatch
}

// FilmLoader contains the client required to load film resources.
type countryLoader struct{}

func (l countryLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var countries []sql.Country

	query := DataLoaderQueryInts(keys, sql.Country{}, countryIDField)
	config.SQL.Raw(query).Scan(&countries)

	var results = make([]*dataloader.Result, len(keys))
	for i, c := range countries {
		results[i] = &dataloader.Result{Data: c}
	}

	return results
}

// LoadCountry loads a country using dataloader
func LoadCountry(ctx context.Context, key int) (*sql.Country, error) {
	var c sql.Country

	data, err := Load(ctx, countryLoaderKey, strconv.Itoa(key))

	if err != nil {
		return nil, err
	}

	c, ok := data.(sql.Country)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", c, data)
	}

	return &c, nil
}
