package copilot

import (
	"hackbar-copilot/internal/domain/menu"
	"slices"
)

// SaveAsMenuGroup implements Copilot.
func (c *copilot) SaveAsMenuGroup(recipeGroupName string, arg SaveAsMenuGroupArg) (menu.Group, error) {
	rg, err := c.FindRecipeGroup(recipeGroupName)
	if err != nil {
		return menu.Group{}, err
	}

	mg := menu.Group{
		Name:     rg.Name,
		ImageURL: rg.ImageURL,
		Flavor:   arg.Flavor,
	}
	for _, r := range rg.Recipes {
		menuItem := menu.Item{
			Name:       r.Name,
			ImageURL:   nil,
			Materials:  []string{},
			OutOfStock: false,
			Price:      0,
		}
		for _, step := range r.Steps {
			if step.Material != nil {
				if !slices.Contains(menuItem.Materials, *step.Material) {
					menuItem.Materials = append(menuItem.Materials, *step.Material)
				}
			}
		}
		arg, ok := arg.Items[r.Name]
		if ok {
			menuItem.ImageURL = arg.ImageURL
			menuItem.Price = arg.Price
		}
		mg.Items = append(mg.Items, menuItem)
	}
	err = c.menu.Save(mg)
	return mg, err
}
