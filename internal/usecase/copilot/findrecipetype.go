package copilot

import (
	"hackbar-copilot/internal/domain/recipe"
)

// FindRecipeType implements Copilot.
func (c *copilot) FindRecipeType() (map[string]recipe.RecipeType, error) {
	recipeTypes := map[string]recipe.RecipeType{}
	for gt, err := range c.recipe.AllRecipeTypes() {
		if err != nil {
			return nil, err
		}
		recipeTypes[gt.Name] = gt
	}
	return recipeTypes, nil
}
