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
	lineRatingLoaderKey        key = "lineRating"
	lineRatingIDField              = "line_id"
	transporterRatingLoaderKey key = "transporterRating"
	transporterRatingIDField       = "transporter_id"
)

func init() {
	loaders[lineRatingLoaderKey] = lineRatingLoader{}.loadBatch
	loaders[transporterRatingLoaderKey] = transporterRatingLoader{}.loadBatch
}

// FilmLoader contains the client required to load film resources.
type lineRatingLoader struct{}
type transporterRatingLoader struct{}

func (l lineRatingLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var lineRatings []sql.LineRating
	query := DataLoaderQueryInts(keys, sql.LineRating{}, lineRatingIDField)
	config.SQL.Raw(query).Scan(&lineRatings)

	var results = make([]*dataloader.Result, len(keys))
	for i, r := range lineRatings {
		results[i] = &dataloader.Result{Data: r}
	}

	return results
}

func (l transporterRatingLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var transporterRatings []sql.TransporterRating
	query := DataLoaderQueryInts(keys, sql.TransporterRating{}, transporterRatingIDField)
	config.SQL.Raw(query).Scan(&transporterRatings)

	var results = make([]*dataloader.Result, len(keys))
	for i, r := range transporterRatings {
		results[i] = &dataloader.Result{Data: r}
	}

	return results
}

// LoadLineRating loads a lineRating using dataloader
func LoadLineRating(ctx context.Context, key int) (*sql.LineRating, error) {
	var r sql.LineRating

	data, err := Load(ctx, lineRatingLoaderKey, strconv.Itoa(key))

	if err != nil {
		return nil, err
	}

	r, ok := data.(sql.LineRating)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", r, data)
	}

	return &r, nil
}

// LoadTransporterRating loads a transporter rating using dataloader
func LoadTransporterRating(ctx context.Context, key int) (*sql.TransporterRating, error) {
	var r sql.TransporterRating

	data, err := Load(ctx, transporterRatingLoaderKey, strconv.Itoa(key))

	if err != nil {
		return nil, err
	}

	r, ok := data.(sql.TransporterRating)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", r, data)
	}

	return &r, nil
}
