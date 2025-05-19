package copilot

import (
	"fmt"
	"hackbar-copilot/internal/domain/menu"
	"slices"
)

// SaveAsMenuItem implements Copilot.
func (c *copilot) SaveAsMenuItem(recipeGroupName string, arg SaveAsMenuItemArg) (menu.Item, error) {
	if arg.Remove {
		err := c.datasource.Menu().Remove(recipeGroupName)
		if err != nil {
			return menu.Item{}, fmt.Errorf("failed to remove menu item: %w", err)
		}
		return menu.Item{}, nil
	}

	rg, err := c.FindRecipeGroup(recipeGroupName)
	if err != nil {
		return menu.Item{}, fmt.Errorf("failed to find recipe group: %w", err)
	}

	materialNames := map[string]bool{}
	materials, err := c.Materials(SortMaterialByName())
	if err != nil {
		return menu.Item{}, fmt.Errorf("failed to retrieve materials: %w", err)
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
		arg, ok := arg.Options[r.Name]
		if !ok {
			continue
		}
		menuItem := menu.ItemOption{
			Name:       r.Name,
			Category:   r.Category,
			ImageURL:   arg.ImageURL,
			Materials:  []string{},
			OutOfStock: false,
			Price:      arg.Price,
		}
		for _, step := range r.Steps {
			if step.Material != nil {
				if !slices.Contains(menuItem.Materials, *step.Material) {
					menuItem.Materials = append(menuItem.Materials, *step.Material)
				}
			}
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

	err = mi.Validate()
	if err != nil {
		return menu.Item{}, fmt.Errorf("failed to create menu item: %w", err)
	}

	err = c.datasource.Menu().Save(mi)
	if err != nil {
		return menu.Item{}, fmt.Errorf("failed to save menu item: %w", err)
	}
	if len(newMaterials) > 0 {
		err = c.datasource.Stock().Save(newMaterials, nil)
		if err != nil {
			return menu.Item{}, fmt.Errorf("failed to save new materials: %w", err)
		}
	}
	return mi, nil
}
