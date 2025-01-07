package adapter

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

type RecipeAdapterOut interface {
	RecipeGroups(
		recipeGroups []recipe.RecipeGroup,
		recipeTypes map[string]recipe.RecipeType,
		glassTypes map[string]recipe.GlassType,
	) []*model.RecipeGroup
	RecipeGroup(
		recipeTypes map[string]recipe.RecipeType,
		glassTypes map[string]recipe.GlassType,
	) func(recipe.RecipeGroup) *model.RecipeGroup
}

func NewRecipeAdapterOut() RecipeAdapterOut {
	return &adapterOut{}
}

type adapterOut struct{}
