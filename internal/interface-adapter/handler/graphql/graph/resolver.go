package graph

import (
	menuadapter "hackbar-copilot/internal/interface-adapter/handler/graphql/adapter/menu"
	orderadapter "hackbar-copilot/internal/interface-adapter/handler/graphql/adapter/order"
	recipeadapter "hackbar-copilot/internal/interface-adapter/handler/graphql/adapter/recipe"
	"hackbar-copilot/internal/usecase/copilot"
	"hackbar-copilot/internal/usecase/order"
)

//go:generate go tool gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver(deps Dependencies) ResolverRoot {
	deps.recipeAdapter = recipeadapter.New()
	deps.menuAdapter = menuadapter.NewOutputAdapter()
	deps.orderAdapter = orderadapter.New()
	return &Resolver{deps}
}

type Resolver struct {
	Dependencies
}

type Dependencies struct {
	Copilot       copilot.Copilot
	recipeAdapter recipeadapter.Adapter
	menuAdapter   menuadapter.MenuAdapter

	OrderService order.Order
	orderAdapter orderadapter.Adapter
}
