package loader

import (
	"fmt"

	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/reservamos/search/internal/config"
	"github.com/reservamos/search/perform/models/sql"
)

const (
	fareServiceLoaderKey key = "fareService"
	fareServiceField         = "(airline_id||'-'||service)"
)

func init() {
	loaders[fareServiceLoaderKey] = fareServiceLoader{}.loadBatch
}

// FilmLoader contains the client required to load film resources.
type fareServiceLoader struct{}

func (l fareServiceLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var fareServices []sql.FareService
	query := DataLoaderQueryStrings(keys, sql.FareService{}, fareServiceField)
	config.SQL.Raw(query).Scan(&fareServices)

	var results = make([]*dataloader.Result, len(keys))
	for i, fs := range fareServices {
		results[i] = &dataloader.Result{Data: fs}
	}

	return results
}

// LoadFareService loads a fareService using dataloader
func LoadFareService(ctx context.Context, airlineID int, service string) (*sql.FareService, error) {
	var fs sql.FareService

	data, err := Load(ctx, fareServiceLoaderKey, fmt.Sprintf("%d-%s", airlineID, service))

	if err != nil {
		return nil, err
	}

	fs, ok := data.(sql.FareService)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", fs, data)
	}

	return &fs, nil
}
