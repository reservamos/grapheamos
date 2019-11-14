package loader

import (
	"fmt"

	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/reservamos/search/internal/config"
	"github.com/reservamos/search/perform/models/sql"
)

const (
	siteLoaderKey key = "site"
	siteIDField       = "(exposable_type||'-'||exposable_id)"
)

func init() {
	loaders[siteLoaderKey] = siteLoader{}.loadBatch
}

// FilmLoader contains the client required to load film resources.
type siteLoader struct{}

func (l siteLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var sites []sql.CmsSiteContent
	query := DataLoaderQueryStrings(keys, sql.CmsSiteContent{}, siteIDField)
	config.SQL.Raw(query).Scan(&sites)

	var results = make([]*dataloader.Result, len(keys))
	for i, s := range sites {
		results[i] = &dataloader.Result{Data: s}
	}

	return results
}

// LoadCmsSiteContent loads a site using dataloader
func LoadCmsSiteContent(ctx context.Context, key string) (*sql.CmsSiteContent, error) {
	var s sql.CmsSiteContent

	data, err := Load(ctx, siteLoaderKey, key)

	if err != nil {
		return nil, err
	}

	s, ok := data.(sql.CmsSiteContent)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", s, data)
	}

	return &s, nil
}
