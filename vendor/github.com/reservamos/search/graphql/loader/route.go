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
	routeLoaderKey key = "route"
	routeIDField       = "id"
)

func init() {
	loaders[routeLoaderKey] = routeLoader{}.loadBatch
}

// FilmLoader contains the client required to load film resources.
type routeLoader struct{}

func (l routeLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var routes []sql.Route
	query := DataLoaderQueryInts(keys, sql.Route{}, routeIDField)
	config.SQL.Raw(query).Scan(&routes)

	var results = make([]*dataloader.Result, len(keys))
	for i, r := range routes {
		results[i] = &dataloader.Result{Data: r}
	}

	return results
}

// LoadRoute loads a route using dataloader
func LoadRoute(ctx context.Context, key int) (*sql.Route, error) {
	var r sql.Route

	data, err := Load(ctx, routeLoaderKey, strconv.Itoa(key))

	if err != nil {
		return nil, err
	}

	r, ok := data.(sql.Route)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", r, data)
	}

	return &r, nil
}
