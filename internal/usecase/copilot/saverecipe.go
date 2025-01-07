package copilot

import (
	"hackbar-copilot/internal/domain/recipe"
)

// SaveRecipe implements Copilot.
func (c *copilot) SaveRecipe(rg recipe.RecipeGroup) error {
	return c.recipe.Save(rg)
}
