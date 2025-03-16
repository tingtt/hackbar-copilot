package http

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewServer(addr string, deps graph.Dependencies, option graphql.Option) *http.Server {
	mux := http.NewServeMux()

	// `/`
	mux.Handle("/", graphql.NewHandler(deps, option))
	// `/playground`
	mux.Handle("/playground", playground.Handler("GraphQL playground", "/"))

	return &http.Server{Addr: addr, Handler: h2c.NewHandler(mux, &http2.Server{})}
}
