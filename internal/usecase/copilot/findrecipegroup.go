package copilot

import (
	"hackbar-copilot/internal/domain/recipe"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
)

// FindRecipeGroup implements Copilot.
func (c *copilot) FindRecipeGroup(name string) (recipe.RecipeGroup, error) {
	for rg, err := range c.recipe.All() {
		if err != nil {
			return recipe.RecipeGroup{}, err
		}
		if rg.Name == name {
			return rg, nil
		}
	}
	return recipe.RecipeGroup{}, usecaseutils.ErrNotFound
}
