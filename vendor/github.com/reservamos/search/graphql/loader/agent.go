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
	agentLoaderKey key = "agent"
	agentIDField       = "id"
)

func init() {
	loaders[agentLoaderKey] = agentLoader{}.loadBatch
}

// FilmLoader contains the client required to load film resources.
type agentLoader struct{}

func (l agentLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var agents []sql.Agent
	query := DataLoaderQueryInts(keys, sql.Agent{}, agentIDField)
	config.SQL.Raw(query).Scan(&agents)

	var results = make([]*dataloader.Result, len(keys))
	for i, a := range agents {
		results[i] = &dataloader.Result{Data: a}
	}

	return results
}

// LoadAgent loads a agent using dataloader
func LoadAgent(ctx context.Context, key int) (*sql.Agent, error) {
	var a sql.Agent

	data, err := Load(ctx, agentLoaderKey, strconv.Itoa(key))

	if err != nil {
		return nil, err
	}

	a, ok := data.(sql.Agent)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", a, data)
	}

	return &a, nil
}
