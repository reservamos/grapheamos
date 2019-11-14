package query

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

// Me resolves sample query
func (r *QueryResolver) Me(ctx context.Context) *UserResolver {
	return &UserResolver{}
}

// UserResolver represents a sample type
type UserResolver struct{}

// ID user id
func (r *UserResolver) ID() graphql.ID {
	return graphql.ID("user:1")
}

// Name user name
func (r *UserResolver) Name() string {
	return "Luke Skywalker"
}
