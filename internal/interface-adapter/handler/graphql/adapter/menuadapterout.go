package adapter

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

type MenuAdapterOut interface {
	MenuGroups(menuGroups []menu.Group, recipeGroups []*model.RecipeGroup) []*model.MenuGroup
}

func NewMenuAdapterOut() MenuAdapterOut {
	return &menuAdapterOut{}
}

type menuAdapterOut struct{}
