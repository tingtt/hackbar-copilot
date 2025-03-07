package menuadapter

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

type OutputAdapter interface {
	MenuItems(items []menu.Item, recipeGroups []*model.RecipeGroup) []*model.MenuItem
}

func NewOutputAdapter() OutputAdapter {
	return &outputAdapter{}
}

type outputAdapter struct{}
