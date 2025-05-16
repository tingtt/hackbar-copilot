package copilot

import (
	"fmt"
	"hackbar-copilot/internal/domain/recipe"
)

// FindRecipeType implements Copilot.
func (c *copilot) FindRecipeType() (map[string]recipe.RecipeType, error) {
	recipeTypes := map[string]recipe.RecipeType{}
	for gt, err := range c.datasource.Recipe().AllRecipeTypes() {
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve recipe types: %w", err)
		}
		recipeTypes[gt.Name] = gt
	}
	return recipeTypes, nil
}
