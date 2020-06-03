package gqlgen

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/cleyonpiv/gqlgen-sqlc-example/dataloaders"
	"github.com/cleyonpiv/gqlgen-sqlc-example/pg"
)

// NewHandler returns a new graphql endpoint handler.
func NewHandler(repo pg.Repository, dl dataloaders.Retriever) http.Handler {
	return handler.GraphQL(NewExecutableSchema(Config{
		Resolvers: &Resolver{
			Repository:  repo,
			DataLoaders: dl,
		},
	}))
}

// NewPlaygroundHandler returns a new GraphQL Playground handler.
func NewPlaygroundHandler(endpoint string) http.Handler {
	return handler.Playground("GraphQL Playground", endpoint)
}
