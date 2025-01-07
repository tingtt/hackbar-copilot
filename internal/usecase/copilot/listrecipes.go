package copilot

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/usecase/sort"
)

// ListRecipes implements Copilot.
func (c *copilot) ListRecipes(sortFunc sort.Yield[recipe.RecipeGroup]) ([]recipe.RecipeGroup, error) {
	var root *sort.Node[recipe.RecipeGroup]
	for rg, err := range c.recipe.All() {
		if err != nil {
			return nil, err
		}
		root = sort.Insert(root, rg, sortFunc)
	}
	recipeGroups := []recipe.RecipeGroup{}
	sort.InOrderTraversal(root, &recipeGroups)
	return recipeGroups, nil
}
