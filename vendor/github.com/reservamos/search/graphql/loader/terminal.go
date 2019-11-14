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
	terminalLoaderKey key = "terminal"
	terminalIDField       = "id"
)

func init() {
	loaders[terminalLoaderKey] = terminalLoader{}.loadBatch
}

// FilmLoader contains the client required to load film resources.
type terminalLoader struct{}

func (l terminalLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var terminals []sql.Terminal

	config.SQL.Raw(DataLoaderOpenQuery(
		keys,
		"places_terminals",
		dlModelFields(sql.Terminal{}),
		terminalIDField,
		"int",
	)).Scan(&terminals)

	var results = make([]*dataloader.Result, len(keys))
	for i, t := range terminals {
		results[i] = &dataloader.Result{Data: t}
	}

	return results
}

// LoadTerminal loads a terminal using dataloader
func LoadTerminal(ctx context.Context, key int) (*sql.Terminal, error) {
	var t sql.Terminal

	data, err := Load(ctx, terminalLoaderKey, strconv.Itoa(key))

	if err != nil {
		return nil, err
	}

	t, ok := data.(sql.Terminal)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", t, data)
	}

	return &t, nil
}
