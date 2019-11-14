package resolver

import (
	"github.com/reservamos/grapheamos/internal/graphql/resolver/mutation"
	"github.com/reservamos/grapheamos/internal/graphql/resolver/query"
)

// Resolver main resolver struct
type Resolver struct {
	*query.QueryResolver
	*mutation.MutationResolver
}
