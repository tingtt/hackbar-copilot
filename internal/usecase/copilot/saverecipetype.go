package copilot

import (
	"fmt"
	"hackbar-copilot/internal/domain/recipe"
)

// SaveRecipeType implements Copilot.
func (c *copilot) SaveRecipeType(rt recipe.RecipeType) error {
	err := rt.Validate()
	if err != nil {
		return fmt.Errorf("failed to create recipe type: invalid recipe type: %w", err)
	}

	err = c.datasource.Recipe().SaveRecipeType(rt)
	if err != nil {
		return fmt.Errorf("failed to save recipe type: %w", err)
	}
	return nil
}
