package graph

import (
	menuadapter "hackbar-copilot/internal/interface-adapter/handler/graphql/adapter/menu"
	recipeadapter "hackbar-copilot/internal/interface-adapter/handler/graphql/adapter/recipe"
	"hackbar-copilot/internal/usecase/copilot"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver(deps Dependencies) ResolverRoot {
	deps.recipeAdapter = recipeadapter.New()
	deps.menuAdapter = menuadapter.NewOutputAdapter()
	return &Resolver{deps}
}

type Resolver struct {
	Dependencies
}

type Dependencies struct {
	Copilot       copilot.Copilot
	recipeAdapter recipeadapter.Adapter
	menuAdapter   menuadapter.MenuAdapter
}
