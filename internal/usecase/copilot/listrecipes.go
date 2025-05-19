package copilot

import (
	"fmt"
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/usecase/sort"
	"slices"
)

// ListRecipes implements Copilot.
func (c *copilot) ListRecipes(sortFunc sort.Yield[recipe.RecipeGroup]) ([]recipe.RecipeGroup, error) {
	var root *sort.Node[recipe.RecipeGroup]
	for rg, err := range c.datasource.Recipe().All() {
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve recipe groups: %w", err)
		}
		slices.SortFunc(rg.Recipes, func(a, b recipe.Recipe) int {
			if a.Name == b.Name {
				return 0
			}
			if a.Name < b.Name {
				return -1
			}
			return 1
		})
		root = sort.Insert(root, rg, sortFunc)
	}
	recipeGroups := []recipe.RecipeGroup{}
	sort.InOrderTraversal(root, &recipeGroups)
	return recipeGroups, nil
}
