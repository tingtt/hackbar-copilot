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
	for _, mg := range menu {
		if mg.Name == rg.Name {
			arg := SaveAsMenuGroupArg{Flavor: mg.Flavor}
			for _, item := range mg.Items {
				arg.Items[item.Name] = MenuFromRecipeGroupArg{
					ImageURL: item.ImageURL,
					Price:    item.Price,
				}
			}
			_, err := c.SaveAsMenuGroup(rg.Name, arg)
			if err != nil {
				return err
			}
			break
		}
	}
	return nil
}
