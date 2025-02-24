package graphql

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph"
	"hackbar-copilot/internal/interface-adapter/handler/middleware"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

type Option struct {
	JWTSecret string
}

func NewHandler(deps graph.Dependencies, option Option) http.Handler {
	handler := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: graph.NewResolver(deps),
	}))
	handler.AddTransport(transport.POST{})

	customHandler := middleware.Wrap(handler)
	customHandler.Use(middleware.JWT([]byte(option.JWTSecret)))
	return customHandler.Handler
}
