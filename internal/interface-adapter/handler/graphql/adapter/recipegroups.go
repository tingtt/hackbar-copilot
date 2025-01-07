package adapter

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/utils/sliceutil"
)

// RecipeGroups implements RecipeAdapterOut.
func (adapter *adapterOut) RecipeGroups(recipeGroups []recipe.RecipeGroup, recipeTypes map[string]recipe.RecipeType, glassTypes map[string]recipe.GlassType) []*model.RecipeGroup {
	return sliceutil.Map(recipeGroups, adapter.RecipeGroup(recipeTypes, glassTypes))
}
