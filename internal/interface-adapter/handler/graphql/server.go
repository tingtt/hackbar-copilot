package graphql

import (
	"context"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func NewHandler(deps graph.Dependencies) http.Handler {
	handler := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: graph.NewResolver(deps),
	}))
	handler.AddTransport(transport.POST{})
	return httpHeaderMiddleware(graph.ContextKeyHeader)(handler)
}

func httpHeaderMiddleware(ctxKey graph.ContextKey) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// put it in context
			ctx := context.WithValue(r.Context(), ctxKey, r.Header)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
