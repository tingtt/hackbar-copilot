package copilot

import (
	"fmt"
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/usecase/sort"
)

// ListRecipes implements Copilot.
func (c *copilot) ListRecipes(sortFunc sort.Yield[recipe.RecipeGroup]) ([]recipe.RecipeGroup, error) {
	var root *sort.Node[recipe.RecipeGroup]
	for rg, err := range c.datasource.Recipe().All() {
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve recipe groups: %w", err)
		}
		root = sort.Insert(root, rg, sortFunc)
	}
	recipeGroups := []recipe.RecipeGroup{}
	sort.InOrderTraversal(root, &recipeGroups)
	return recipeGroups, nil
}
