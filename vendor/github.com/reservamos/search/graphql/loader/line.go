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
	lineLoaderKey key = "line"
	lineIDField       = "id"
)

func init() {
	loaders[lineLoaderKey] = lineLoader{}.loadBatch
}

// FilmLoader contains the client required to load film resources.
type lineLoader struct{}

func (l lineLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var lines []sql.Line
	query := DataLoaderQueryInts(keys, sql.Line{}, lineIDField)
	config.SQL.Raw(query).Scan(&lines)

	var results = make([]*dataloader.Result, len(keys))
	for i, l := range lines {
		results[i] = &dataloader.Result{Data: l}
	}

	return results
}

// LoadLine loads a line using dataloader
func LoadLine(ctx context.Context, key int) (*sql.Line, error) {
	var l sql.Line

	data, err := Load(ctx, lineLoaderKey, strconv.Itoa(key))

	if err != nil {
		return nil, err
	}

	l, ok := data.(sql.Line)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", l, data)
	}

	return &l, nil
}
