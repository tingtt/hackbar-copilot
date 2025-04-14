package copilot

import (
	"hackbar-copilot/internal/domain/menu"
	"slices"
)

// SaveAsMenuItem implements Copilot.
func (c *copilot) SaveAsMenuItem(recipeGroupName string, arg SaveAsMenuItemArg) (menu.Item, error) {
	if arg.Remove {
		err := c.menu.Remove(recipeGroupName)
		return menu.Item{}, err
	}

	rg, err := c.FindRecipeGroup(recipeGroupName)
	if err != nil {
		return menu.Item{}, err
	}

	materialNames := map[string]bool{}
	materials, err := c.Materials(SortMaterialByName())
	if err != nil {
		return menu.Item{}, err
	}
	for _, material := range materials {
		materialNames[material.Name] = true
	}
	newMaterials := []string{}

	mi := menu.Item{
		Name:     rg.Name,
		ImageURL: rg.ImageURL,
		Flavor:   arg.Flavor,
	}
	for _, r := range rg.Recipes {
		menuItem := menu.ItemOption{
			Name:       r.Name,
			Category:   r.Category,
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
		arg, ok := arg.Options[r.Name]
		if ok {
			menuItem.ImageURL = arg.ImageURL
			menuItem.Price = arg.Price
		}
		mi.Options = append(mi.Options, menuItem)
		for _, materialName := range menuItem.Materials {
			if _, exists := materialNames[materialName]; !exists {
				if !slices.Contains(newMaterials, materialName) {
					newMaterials = append(newMaterials, materialName)
				}
			}
		}
	}
	err = c.menu.Save(mi)
	if err != nil {
		return menu.Item{}, err
	}
	if len(newMaterials) > 0 {
		err = c.stock.Save(newMaterials, nil)
		if err != nil {
			return menu.Item{}, err
		}
	}
	return mi, nil
}
