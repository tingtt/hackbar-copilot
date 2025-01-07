package copilot

import (
	"hackbar-copilot/internal/domain/recipe"
)

// SaveRecipeType implements Copilot.
func (c *copilot) SaveRecipeType(rt recipe.RecipeType) error {
	return c.recipe.SaveRecipeType(rt)
}
