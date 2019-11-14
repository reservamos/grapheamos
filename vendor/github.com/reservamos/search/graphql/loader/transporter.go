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
	transporterLoaderKey key = "transporter"
	transporterIDField       = "id"
)

func init() {
	loaders[transporterLoaderKey] = transporterLoader{}.loadBatch
}

// FilmLoader contains the client required to load film resources.
type transporterLoader struct{}

func (l transporterLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var transporters []sql.Transporter
	query := DataLoaderQueryInts(keys, sql.Transporter{}, transporterIDField)
	config.SQL.Raw(query).Scan(&transporters)

	var results = make([]*dataloader.Result, len(keys))
	for i, t := range transporters {
		results[i] = &dataloader.Result{Data: t}
	}

	return results
}

// LoadTransporter loads a transporter using dataloader
func LoadTransporter(ctx context.Context, key int) (*sql.Transporter, error) {
	var t sql.Transporter

	data, err := Load(ctx, transporterLoaderKey, strconv.Itoa(key))

	if err != nil {
		return nil, err
	}

	t, ok := data.(sql.Transporter)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", t, data)
	}

	return &t, nil
}
