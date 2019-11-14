package graphql

import (
	"context"
	"fmt"
	"net/http"
	"os"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/reservamos/graphql-start/assets"
	"github.com/reservamos/graphql-start/internal/graphql/handler"
	"github.com/reservamos/graphql-start/internal/graphql/resolver"
	"github.com/reservamos/graphql-start/internal/graphql/schema"
)

//Start starts the server which will be manage all clients requests
func Start() {
	ctx := context.Background()
	http.Handle("/graphql", graphQLHandler(ctx))
	// GraphiQL development tool
	http.Handle("/", graphiQLHandler())
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}

func graphQLHandler(ctx context.Context) http.Handler {
	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})
	return handler.AddContext(ctx, &handler.GraphQL{
		Schema: graphqlSchema,
	})
}

func graphiQLHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := assets.Asset("assets/web/graphiql.html")
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
}
