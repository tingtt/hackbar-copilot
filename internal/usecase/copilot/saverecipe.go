package copilot

import (
	"fmt"
	"hackbar-copilot/internal/domain/recipe"
)

// SaveRecipe implements Copilot.
func (c *copilot) SaveRecipe(rg recipe.RecipeGroup) error {
	err := rg.Validate()
	if err != nil {
		return fmt.Errorf("invalid recipe group: %w", err)
	}

	err = c.datasource.Recipe().Save(rg)
	if err != nil {
		return fmt.Errorf("failed to save recipe group: %w", err)
	}

	for mi, err := range c.datasource.Menu().All() {
		if err != nil {
			return fmt.Errorf("failed to retrieve menu items: %w", err)
		}
		if mi.Name == rg.Name {
			arg := SaveAsMenuItemArg{Flavor: mi.Flavor, Options: map[string]MenuFromRecipeGroupArg{}}
			for _, option := range mi.Options {
				arg.Options[option.Name] = MenuFromRecipeGroupArg{
					ImageURL: option.ImageURL,
					Price:    option.Price,
				}
			}
			_, err := c.SaveAsMenuItem(rg.Name, arg)
			if err != nil {
				return fmt.Errorf("failed to save menu item: %w", err)
			}
			break
		}
	}
	return nil
}
