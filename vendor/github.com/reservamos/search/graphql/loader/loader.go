package loader

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
)

type key string

var loaders = map[key]dataloader.BatchFunc{}

// Collection holds an internal lookup of initialized batch data load functions.
type Collection struct {
	dataloaderFuncMap map[key]dataloader.BatchFunc
}

// NewCollection Initialize a lookup map of context keys to batch functions.
// 	When Attach is called on the Collection, the batch functions are used to create new dataloader
// 	instances. The instances are attached to the request context at the provided keys.
// 	The keys are then used to Extract the dataloader instances from the request context.
func NewCollection() Collection {
	return Collection{
		dataloaderFuncMap: loaders,
	}
}

// Attach attaches a new dataloader batch function to the collection
func (c Collection) Attach(ctx context.Context) context.Context {
	for k, batchFn := range c.dataloaderFuncMap {
		ctx = context.WithValue(ctx, k, dataloader.NewBatchedLoader(batchFn))
	}

	return ctx
}

// Extract is a helper function to make common get-value, assert-type, return-error-or-value
// operations easier.
func Extract(ctx context.Context, k key) (*dataloader.Loader, error) {
	ldr, ok := ctx.Value(k).(*dataloader.Loader)
	if !ok {
		return nil, fmt.Errorf("unable to find %s loader on the request context", k)
	}

	return ldr, nil
}

// Load given a loader key and a key retrieves an interface from dataloader
func Load(ctx context.Context, loaderKey key, key string) (interface{}, error) {
	ldr, err := Extract(ctx, loaderKey)
	if err != nil {
		return nil, err
	}
	data, err := ldr.Load(ctx, dataloader.StringKey(key))()
	if err != nil {
		return nil, err
	}
	return data, nil
}
