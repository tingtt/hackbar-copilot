package copilot

import (
	"fmt"
	"hackbar-copilot/internal/domain/recipe"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
)

// FindRecipeGroup implements Copilot.
func (c *copilot) FindRecipeGroup(name string) (recipe.RecipeGroup, error) {
	for rg, err := range c.datasource.Recipe().All() {
		if err != nil {
			return recipe.RecipeGroup{}, fmt.Errorf("failed to retrieve recipe groups: %w", err)
		}
		if rg.Name == name {
			return rg, nil
		}
	}
	return recipe.RecipeGroup{}, fmt.Errorf("recipe group \"%s\" %w", name, usecaseutils.ErrNotFound)
}
