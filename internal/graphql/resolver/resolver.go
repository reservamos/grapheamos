package resolver

import (
	"github.com/reservamos/graphql-start/internal/graphql/resolver/mutation"
	"github.com/reservamos/graphql-start/internal/graphql/resolver/query"
)

// Resolver main resolver struct
type Resolver struct {
	*query.ResolveQ
	*mutation.ResolveM
}
