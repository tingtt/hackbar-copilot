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

	materialNames := map[string]bool{}
	materials, err := c.Materials(SortMaterialByName())
	if err != nil {
		return menu.Group{}, err
	}
	for _, material := range materials {
		materialNames[material.Name] = true
	}
	newMaterials := []string{}

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
		for _, materialName := range menuItem.Materials {
			if _, exists := materialNames[materialName]; !exists {
				if !slices.Contains(newMaterials, materialName) {
					newMaterials = append(newMaterials, materialName)
				}
			}
		}
	}
	err = c.menu.Save(mg)
	if err != nil {
		return menu.Group{}, err
	}
	if len(newMaterials) > 0 {
		err = c.stock.Save(newMaterials, nil)
		if err != nil {
			return menu.Group{}, err
		}
	}
	return mg, nil
}
