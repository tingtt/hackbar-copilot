package menuadapter

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

type OutputAdapter interface {
	MenuGroups(menuGroups []menu.Group, recipeGroups []*model.RecipeGroup) []*model.MenuGroup
}

func NewOutputAdapter() OutputAdapter {
	return &outputAdapter{}
}

type outputAdapter struct{}
