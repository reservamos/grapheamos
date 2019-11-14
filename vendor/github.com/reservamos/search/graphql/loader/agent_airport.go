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
	agentAirportLoaderKey key = "agentAirport"
	agentAirportIDField       = "id"
)

func init() {
	loaders[agentAirportLoaderKey] = agentAirportLoader{}.loadBatch
}

// FilmLoader contains the client required to load film resources.
type agentAirportLoader struct{}

func (l agentAirportLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var agentAirports []sql.AgentAirport
	query := DataLoaderQueryInts(keys, sql.AgentAirport{}, agentAirportIDField)
	config.SQL.Raw(query).Scan(&agentAirports)

	var results = make([]*dataloader.Result, len(keys))
	for i, aa := range agentAirports {
		results[i] = &dataloader.Result{Data: aa}
	}

	return results
}

// LoadAgentAirport loads a agentAirport using dataloader
func LoadAgentAirport(ctx context.Context, key int) (*sql.AgentAirport, error) {
	var aa sql.AgentAirport

	data, err := Load(ctx, agentAirportLoaderKey, strconv.Itoa(key))

	if err != nil {
		return nil, err
	}

	aa, ok := data.(sql.AgentAirport)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", aa, data)
	}

	return &aa, nil
}
