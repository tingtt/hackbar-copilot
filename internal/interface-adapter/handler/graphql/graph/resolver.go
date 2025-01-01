package graph

import (
	"hackbar-copilot/internal/usecase/orders"
	"hackbar-copilot/internal/usecase/recipes"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver(deps Dependencies) ResolverRoot {
	deps.convertToModel = &converter{}
	return &Resolver{deps}
}

type Resolver struct {
	deps Dependencies
}

type Dependencies struct {
	Orders         orders.Service
	Recipes        recipes.Service
	convertToModel converterI
}
