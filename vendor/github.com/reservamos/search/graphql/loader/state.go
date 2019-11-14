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
	stateLoaderKey key = "state"
	stateIDField       = "id"
)

func init() {
	loaders[stateLoaderKey] = stateLoader{}.loadBatch
}

// FilmLoader contains the client required to load film resources.
type stateLoader struct{}

func (l stateLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var states []sql.State
	query := DataLoaderQueryInts(keys, sql.State{}, stateIDField)
	config.SQL.Raw(query).Scan(&states)

	var results = make([]*dataloader.Result, len(keys))
	for i, r := range states {
		results[i] = &dataloader.Result{Data: r}
	}

	return results
}

// LoadState loads a state using dataloader
func LoadState(ctx context.Context, key int) (*sql.State, error) {
	var r sql.State

	data, err := Load(ctx, stateLoaderKey, strconv.Itoa(key))

	if err != nil {
		return nil, err
	}

	r, ok := data.(sql.State)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", r, data)
	}

	return &r, nil
}
