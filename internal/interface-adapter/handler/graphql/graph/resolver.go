package graph

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/adapter"
	"hackbar-copilot/internal/usecase/copilot"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver(deps Dependencies) ResolverRoot {
	deps.recipeAdapter = adapter.NewRecipeAdapter()
	return &Resolver{deps}
}

type Resolver struct {
	Dependencies
}

type Dependencies struct {
	Copilot       copilot.Copilot
	recipeAdapter adapter.RecipeAdapter
}
