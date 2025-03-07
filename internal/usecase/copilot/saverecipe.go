package copilot

import (
	"hackbar-copilot/internal/domain/recipe"
)

// SaveRecipe implements Copilot.
func (c *copilot) SaveRecipe(rg recipe.RecipeGroup) error {
	err := c.recipe.Save(rg)
	if err != nil {
		return err
	}

	menu, err := c.ListMenu(SortMenuGroupByName())
	if err != nil {
		return err
	}
	for _, mi := range menu {
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
				return err
			}
			break
		}
	}
	return nil
}
