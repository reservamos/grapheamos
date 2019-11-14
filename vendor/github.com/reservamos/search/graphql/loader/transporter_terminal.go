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
	ttLoaderKey key = "transporterTerminal"
	ttIDField       = "id"
)

func init() {
	loaders[ttLoaderKey] = transporterTerminalLoader{}.loadBatch
}

// FilmLoader contains the client required to load film resources.
type transporterTerminalLoader struct{}

func (l transporterTerminalLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var transporterTerminals []sql.TransporterTerminal
	query := DataLoaderQueryInts(keys, sql.TransporterTerminal{}, ttIDField)
	config.SQL.Raw(query).Scan(&transporterTerminals)

	var results = make([]*dataloader.Result, len(keys))
	for i, tt := range transporterTerminals {
		results[i] = &dataloader.Result{Data: tt}
	}

	return results
}

// LoadTransporterTerminal loads a transporter terminal using dataloader
func LoadTransporterTerminal(ctx context.Context, key int) (*sql.TransporterTerminal, error) {
	var tt sql.TransporterTerminal

	data, err := Load(ctx, ttLoaderKey, strconv.Itoa(key))

	if err != nil {
		return nil, err
	}

	tt, ok := data.(sql.TransporterTerminal)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", tt, data)
	}

	return &tt, nil
}
