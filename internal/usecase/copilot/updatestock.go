package copilot

import "fmt"

// InStock implements Copilot.
func (c *copilot) UpdateStock(inStockMaterials, outOfStockMaterials []string) error {
	err := c.datasource.Stock().Save(inStockMaterials, outOfStockMaterials)
	if err != nil {
		return fmt.Errorf("failed to save stock: %w", err)
	}

	materials, err := c.Materials(SortMaterialByName())
	if err != nil {
		return fmt.Errorf("failed to retrieve materials: %w", err)
	}
	inStockMaterialNames := map[string]bool{}
	for _, material := range materials {
		if material.InStock {
			inStockMaterialNames[material.Name] = true
		}
	}

	for menuGroup, err := range c.datasource.Menu().All() {
		if err != nil {
			return err
		}
		update := false
		for i, item := range menuGroup.Options {
			currentStockStatus := item.OutOfStock
			menuGroup.Options[i].OutOfStock = false
			for _, materialName := range item.Materials {
				if _, inStock := inStockMaterialNames[materialName]; !inStock {
					menuGroup.Options[i].OutOfStock = true
					break
				}
			}
			if menuGroup.Options[i].OutOfStock != currentStockStatus && !update {
				update = true
			}
		}
		if update {
			err := menuGroup.Validate()
			if err != nil {
				return fmt.Errorf("failed to update menu group: invalid menu group: %w", err)
			}

			err = c.datasource.Menu().Save(menuGroup)
			if err != nil {
				return fmt.Errorf("failed to save menu group: %w", err)
			}
		}
	}
	return nil
}
