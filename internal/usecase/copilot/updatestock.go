package copilot

// InStock implements Copilot.
func (c *copilot) UpdateStock(inStockMaterials, outOfStockMaterials []string) error {
	err := c.stock.Save(inStockMaterials, outOfStockMaterials)
	if err != nil {
		return err
	}

	materials, err := c.Materials(SortMaterialByName())
	if err != nil {
		return err
	}
	inStockMaterialNames := map[string]bool{}
	for _, material := range materials {
		if material.InStock {
			inStockMaterialNames[material.Name] = true
		}
	}

	for menuGroup, err := range c.menu.All() {
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
			err := c.menu.Save(menuGroup)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
